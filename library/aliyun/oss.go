package aliyun

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/syyongx/php2go"
)

//const (
//	//OssBucketImage 定义图片存储服务名字
//	OssBucketImage = "xs-image"
//)

var ossImageConfig = &KeyConfig{}
var ossLocalConfig = &KeyLocalConfig{}

// NewOss 根据名字实例化一个OSS对象
//func NewOss(ctx context.Context, name string, uid ...uint32) (*Oss, error) {
//	var cfg *KeyConfig
//	var target uint32 = 0
//	if len(uid) > 0 {
//		target = uid[0]
//	}
//	switch name {
//	case OssBucketImage:
//		cfg = ossImageConfig
//	}
//	if cfg != nil {
//		client := &Oss{
//			Config:     cfg,
//			BucketName: name,
//			UID:        target,
//		}
//		err := client.init()
//		if err == nil {
//			return client, nil
//		}
//		return nil, err
//	}
//	return nil, gerror.Newf("error bucket name %s", name)
//}

func NewLocalOss(ctx context.Context, ossCfgKey string, uid ...uint64) (*Oss, error) {
	var u uint64
	if len(uid) > 0 {
		u = uid[0]
	}
	ossCfg, err := g.Cfg().Get(ctx, ossCfgKey)
	if err != nil {
		panic(ossCfgKey + "no oss config")
	}
	ossCfg.Struct(ossLocalConfig)
	client := &Oss{
		UID:        u,
		BucketName: ossLocalConfig.Bucket,
		CdnDomain:  ossLocalConfig.CdnDomain,
		Config: &KeyConfig{
			Endpoint: ossLocalConfig.Endpoint,
			Key:      ossLocalConfig.AccessKey,
			Secret:   ossLocalConfig.AccessSecret,
		},
	}
	err = client.init()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Oss 一般情况下，没必要保存这个Oss对象
type Oss struct {
	UID        uint64
	BucketName string
	CdnDomain  string
	Config     *KeyConfig
	client     *oss.Client
	bucket     *oss.Bucket
}

// IsObjectExist 判断文件是否存在
func (serv *Oss) IsObjectExist(objectKey string) (bool, error) {
	return serv.bucket.IsObjectExist(objectKey)
}

// GetObjectToFile 把文件从服务器下载到本地
func (serv *Oss) GetObjectToFile(objectKey, localFile string) error {
	return serv.bucket.GetObjectToFile(objectKey, localFile)
}

// DeleteObject 从服务器删除文件
func (serv *Oss) DeleteObject(objectKey string) error {
	return serv.bucket.DeleteObject(objectKey)
}

// UploadLocalFile 从本地把文件上传到服务器，系统会根据uid自动生成文件名字
func (serv *Oss) UploadLocalFile(localFile string, ext ...string) (string, error) {
	extension := ""
	if len(ext) > 0 {
		extension = ext[0]
	} else {
		extension = serv.getExtension(localFile)
	}
	if len(extension) > 0 {
		extension = "." + extension
	}
	remoteName := ""
	if serv.UID > 0 {
		remoteName = fmt.Sprintf(
			"upload/%s/%d_%s%s",
			php2go.Date("200601/02", time.Now().Unix()),
			serv.UID,
			php2go.Uniqid(""),
			extension,
		)
	} else {
		remoteName = fmt.Sprintf(
			"upload/%s/%s%s",
			php2go.Date("200601/02", time.Now().Unix()),
			php2go.Uniqid(""),
			extension,
		)
	}
	options := []oss.Option{
		oss.ResponseCacheControl("max-age=31536000"),
	}
	err := serv.bucket.PutObjectFromFile(remoteName, localFile, options...)
	return remoteName, err
}

func (serv *Oss) init() error {
	if serv.client != nil {
		return nil
	}
	client, err := oss.New(
		serv.Config.Endpoint,
		serv.Config.Key,
		serv.Config.Secret,
		oss.Timeout(3, 30),
	)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(serv.BucketName)
	if err != nil {
		return err
	}
	serv.client = client
	serv.bucket = bucket
	return nil
}

func (serv *Oss) getExtension(name string) string {
	index := strings.LastIndex(name, ".")
	if index < 0 {
		return ""
	}
	return name[index+1:]
}

func (serv *Oss) FormatOssFileUrl(ossFileUrl string) string {
	if serv.CdnDomain != "" {
		return filepath.Join(serv.CdnDomain, ossFileUrl)
	}

	return filepath.Join(serv.BucketName+"."+serv.Config.Endpoint, ossFileUrl)
}

// GetObject 从服务器删除文件
func (serv *Oss) GetObject(objectKey string, options ...oss.Option) (buffer bytes.Buffer, err error) {
	objectReader, err := serv.bucket.GetObject(objectKey, options...)
	if err != nil {
		return buffer, err
	}
	defer objectReader.Close()

	if _, err = buffer.ReadFrom(objectReader); err != nil {
		return
	}
	return
}
