package main

import (
	"github.com/gowcar/gboot"
	"github.com/gowcar/gboot-sandbox/generated"
	"github.com/gowcar/gboot-sandbox/internal/user"
	"github.com/gowcar/gboot/pkg/log"
)

func main() {
	exec()
}

func exec() {
	gboot.RegisterAnnotations(generated.AllPackages())
	gboot.StartApplication()
	log.Debug("check user.DefaultName = %v", user.DefaultUserName)
	log.Debug("check user.Timeout = %v", user.Timeout)
}
