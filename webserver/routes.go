package webserver

import (
	"net/http"

	"github.com/ankur-toko/quick-links/core/models"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) error {
	e.GET("/save", Save)
	e.GET("/:key", Get)
	return nil
}

func Save(c echo.Context) error {
	k := c.QueryParam("key")
	u := c.QueryParam("url")

	err := mycore.SaveQuickLink(models.QuickLink{Key: k, URL: u})

	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "saved successfull")
	}

	return err
}

func Get(c echo.Context) error {
	k := c.Param("key")
	ql := mycore.GetQuickLink(k)
	c.JSON(http.StatusOK, ql.ToJSON())
	return nil
}
