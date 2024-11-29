package cmd

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Main = cMain{}
	rMap = map[string]interface{}{
		"demo": &RDemo{},
	}
)

type cMain struct {
	g.Meta `name:"main" brief:"main"`
}

type cMainInput struct {
	g.Meta `name:"main"`
	Name   string `v:"required" short:"n" name:"name"     brief:"class name"`
	Action string `v:"required" short:"a" name:"action"     brief:"action name"`
}

type cMainOutput struct{}

func (c cMain) Index(ctx context.Context, in cMainInput) (out *cMainOutput, err error) {
	//反射
	if _, ok := rMap[in.Name]; !ok {
		return nil, errors.New("no found struct config")
	}
	rValue := reflect.ValueOf(rMap[in.Name])
	method := rValue.MethodByName(in.Action)
	if !method.IsValid() {
		return nil, errors.New("struct no found func")
	}
	args := []reflect.Value{reflect.ValueOf(ctx)}
	method.Call(args)
	return
}

// RDemo 这是一个测试，-n demo -a Dt即可获取
type RDemo struct {
}

// Dt 给main方法做的一个demo
func (c RDemo) Dt(ctx context.Context) (err error) {
	fmt.Println("i am a test")
	return nil
}
