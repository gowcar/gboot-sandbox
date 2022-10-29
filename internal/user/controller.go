package user

import (
	"fmt"
	"github.com/gowcar/gboot"
	"github.com/gowcar/gboot/pkg/log"
)

// @Config("user.default_username")
var DefaultUserName string

// @Config("application.name")
var applicationName string

// @Config("application.timeout")
var Timeout int

// @GetMapping("/api/static/config")
func GlobalController(c *gboot.Context) error {
	return c.SendString("Hello, World!")
}

// @RestController
// @RequestMapping("/api/users")
type UserController struct {
	// @Config("user.fetch_size")
	FetchSize int
}

// @GetMapping("/")
func (s *UserController) UserHandler(c *gboot.Context) error {
	log.Debug("applicationName is :%v", applicationName)
	log.Debug("Timeout is :%v", Timeout)
	log.Debug("DefaultUserName is :%v", DefaultUserName)
	log.Debug("FetchSize is :%v", s.FetchSize)

	str := fmt.Sprintf("default fetch size is %v", s.FetchSize)
	return c.SendString(str)
}

func init() {
	DefaultUserName = "Default123"
}
