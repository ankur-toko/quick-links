package webserver

import (
	"errors"
	"fmt"

	"github.com/ankur-toko/quick-links/config"
	"github.com/ankur-toko/quick-links/core"
	"github.com/labstack/echo"
)

var mycore core.Core

func Start() error {
	var err error
	mycore, err = core.CreateCoreClassObject()

	if err != nil {
		return err
	}
	web_server := echo.New()

	Routes(web_server)
	port, ok := config.Get("app.port")
	if !ok {
		return errors.New(PORT_NOT_DEFINED)
	}
	err = web_server.Start("0.0.0.0:" + port)
	fmt.Println("error starting web server", err)
	return nil
}
