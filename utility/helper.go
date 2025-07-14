package utility

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-module/carbon/v2"
)

func GetHeadUrl(ctx context.Context, url string) string {
	if url == "" {
		return url
	}
	match, err := regexp.MatchString(`(http|https)://.*`, url)
	if err != nil {
		return url
	}
	if match {
		return url
	}
	ossUrlVar, _ := g.Cfg().Get(ctx, "server.ossUrl")
	ossUrl := ossUrlVar.String()
	s := strings.Split(ossUrl, ",")
	index := grand.N(0, len(s)-1)

	url = strings.TrimLeft(url, "/")
	return fmt.Sprintf("%s%s%s", s[index], "/", url)
}

func DelUrlHead(srcinUrl string) string {
	if srcinUrl == "" {
		return ""
	}

	match, err := regexp.MatchString(`(http|https)://.*`, srcinUrl)
	if err != nil {
		return srcinUrl
	}
	if !match {
		return srcinUrl
	}

	parsed, err := url.Parse(srcinUrl)
	if err != nil {
		return srcinUrl
	}
	return strings.Trim(parsed.Path, "/")
}

func GetAge(birth int64) int32 {
	car := carbon.CreateFromTimestamp(birth)
	return int32(car.Age())
}

func InArray[T int | string | int64 | int32 | uint](needle T, haystack []T) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

// ArrayDiff 差集
func ArrayDiff[T int | string | int64](dst, src []T) []T {
	m := make(map[T]bool)
	for _, item := range src {
		m[item] = true
	}
	var diff []T
	for _, item := range dst {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// ArrayUnique 唯一
func ArrayUnique[T int | string | int64 | int32 | uint](src []T) []T {
	if len(src) == 0 {
		return src
	}
	m := make(map[T]bool)
	var uniqueArr []T

	for _, item := range src {
		if _, ok := m[item]; !ok {
			m[item] = true

			uniqueArr = append(uniqueArr, item)
		}
	}
	return uniqueArr
}

// ArrayIntersect 交集
func ArrayIntersect[T int | string | int64](slice1, slice2 []T) []T {
	elemMap := make(map[T]bool)
	for _, item := range slice2 {
		elemMap[item] = true
	}
	var intersect []T
	for _, item := range slice1 {
		if _, ok := elemMap[item]; ok {
			intersect = append(intersect, item)
		}
	}
	return intersect
}

// ArrayChunk 分块
func ArrayChunk[T int | string | int64 | interface{}](slice1 []T, size int) [][]T {
	var chunks [][]T
	length := len(slice1)
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = len(slice1)
		}
		chunks = append(chunks, slice1[i:end])
	}
	return chunks
}

// SearchMapKeyByValue 根据map的值值查询key
func SearchMapKeyByValue[T int | string | int64 | int32, T1 string](value T1, dataMap map[T]T1) (key T, found bool) {
	for key, val := range dataMap {
		if val == value {
			return key, true
		}
	}
	return
}

func CopyFields(src interface{}, to interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	json.Unmarshal(data, to)
	return nil
}

func RedisLock(ctx context.Context, cache *gredis.Redis, key string, value interface{}, timeout int64) bool {
	do, err := cache.Do(ctx, "set", key, value, "nx", "ex", timeout)
	if err != nil {
		g.Log().Errorf(ctx, "redis lock err, err=%+v", err)
		return false
	}
	return do.Bool()
}

// Snake2Camel snake转驼峰
func Snake2Camel(s string) string {
	var camelString string

	// Split the string by underscores
	words := strings.Split(s, "_")

	// Capitalize the first letter of each word
	for j, word := range words {
		words[j] = strings.Title(word)
	}

	// Join the words with spaces
	camelString = strings.Join(words, "")
	return camelString
}

// BuildAuditEncrypt 构建审核的加密
func BuildAuditEncrypt(ak, sk string, duration int64) string {
	encryStr := fmt.Sprintf("%s:%s:%d", ak, sk, duration)
	encryToken := fmt.Sprintf("%s:%x:%d", ak, md5.Sum([]byte(encryStr)), duration)

	return base64.URLEncoding.EncodeToString([]byte(encryToken))
}

func CheckDecryptAuditAkSk(ctx context.Context, authorization, adminAk, adminSk string) bool {
	authorization = strings.Replace(authorization, "Basic ", "", 1)
	decodeString, err := base64.URLEncoding.DecodeString(authorization)
	if err != nil {
		g.Log().Warningf(ctx, "CheckAdminAkSk1 err=%v", err.Error())
		return false
	}
	matchs := strings.SplitN(string(decodeString), ":", 3)
	if len(matchs) != 3 {
		g.Log().Warningf(ctx, "CheckAdminAkSk2 len=%d", len(matchs))
		return false
	}
	ak, token, reqTime := matchs[0], matchs[1], matchs[2]
	//adminAk := gcfg.Instance().GetString("server.AdminAk")
	//adminSk := gcfg.Instance().GetString("server.AdminSk")
	if ak != adminAk {
		g.Log().Warningf(ctx, "CheckAdminAkSk3 ak=%v, adminAk=%v", ak, adminAk)
		return false
	}

	encryStr := fmt.Sprintf("%s:%s:%s", ak, adminSk, reqTime)
	encryToken := fmt.Sprintf("%x", md5.Sum([]byte(encryStr)))
	if encryToken != token {
		g.Log().Warningf(ctx, "CheckAdminAkSk4 encryToken=%v, token=%v", encryToken, token)
		return false
	}
	atoi, err := strconv.Atoi(reqTime)
	if err != nil {
		g.Log().Warningf(ctx, "CheckAdminAkSk5 err=%v", err.Error())
		return false
	}
	currentTime := time.Now().Unix()
	if int64(atoi)+300 < currentTime {
		g.Log().Warningf(ctx, "CheckAdminAkSk6 time error")
		return false
	}
	return true
}

func createMosaic(src image.Image, size int) image.Image {
	bounds := src.Bounds()
	mosaic := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += size {
		for x := bounds.Min.X; x < bounds.Max.X; x += size {
			col := src.At(x, y)
			for i := 0; i < size && y+i < bounds.Max.Y; i++ {
				for j := 0; j < size && x+j < bounds.Max.X; j++ {
					draw.Draw(mosaic, image.Rect(x+j, y+i, x+j+1, y+i+1), &image.Uniform{col}, image.ZP, draw.Src)
				}
			}
		}
	}

	return mosaic
}

func ChangePic2Mosaic(localFile string) (err error) {
	file, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// 使用DetectContentType函数获取MIME类型
	mimeType := http.DetectContentType(content)
	var src image.Image
	buf := bytes.NewBuffer(content)
	if mimeType == "image/jpeg" {
		src, err = jpeg.Decode(buf)
	} else if mimeType == "image/png" {
		src, err = png.Decode(buf)
	} else {
		//报错
		err = errors.New("mimeType error")
	}

	if err != nil {
		return err
	}

	mosaic := createMosaic(src, 70)

	outFile, err := os.Create(localFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	if mimeType == "image/jpeg" {
		err = jpeg.Encode(outFile, mosaic, &jpeg.Options{Quality: 20})
	} else if mimeType == "image/png" {
		err = png.Encode(outFile, mosaic)
	}
	return err
}

func RandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	lettersLen := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(lettersLen)]
	}
	return string(b)
}
