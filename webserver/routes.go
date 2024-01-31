package webserver

import (
	"net/http"

	"github.com/ankur-toko/quick-links/core/models"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) error {
	e.GET("/save", Save)
	e.GET("/favicon.ico", Favicon)
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
		c.String(http.StatusOK, "saved successful")
	}
	return err
}

func Get(c echo.Context) error {
	k := c.Param("key")
	ql := mycore.GetQuickLink(k)

	if ql != nil && ql.URL != "" {
		c.Redirect(302, ql.URL)
		// c.JSON(http.StatusOK, ql.ToJSON())
	} else {
		c.String(404, "404 not found")
	}

	return nil
}

func Favicon(c echo.Context) error {
	return nil
}
