package main

import (
	"fmt"
	"github.com/gowcar/gboot"
	"github.com/gowcar/gboot-sandbox/generated"
	"github.com/gowcar/gboot-sandbox/internal/user"
	"github.com/gowcar/gboot/pkg/log"
	"github.com/gowcar/gboot/pkg/proxy"
	"reflect"
)

func main() {
	//exec()
	exec1()
}

type Hello interface {
	SayHello() string
}

type HelloImpl struct{}

func (h *HelloImpl) SayHello() string {
	log.Debug("some thing")
	return "Hello"
}

type Hello1 struct {
	SayHello func() string
}

func exec1() {
	gboot.RegisterAnnotations(generated.AllPackages())
	gboot.StartApplication()
	proxy1 := proxy.InvocationProxy.NewProxyInstance(&HelloImpl{}, func(obj any, method proxy.InvocationMethod, args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("This is a proxy function")}
	}).(*HelloImpl)
	fmt.Println(proxy1.SayHello())
	log.Debug("check user.DefaultName = %v", user.DefaultUserName)
	log.Debug("check user.Timeout = %v", user.Timeout)
}

func exec() {
	gboot.RegisterAnnotations(generated.AllPackages())
	gboot.StartApplication()
	log.Debug("check user.DefaultName = %v", user.DefaultUserName)
	log.Debug("check user.Timeout = %v", user.Timeout)
}
