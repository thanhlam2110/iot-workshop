package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/thanhlam/iot-workshop/service"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()

	/*e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "IOT API FOR ORGANIZATION1")
	})*/
	//new
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})
	e.Logger.Fatal(e.StartAutoTLS(":443"))

	//new

	//USER MANAGEMENT
	//<--------------------------- USER MANAGEMENT ------------------------------>
	e.POST("/api/sso/user/register", service.CreateUser)
	//<--------------------------- THINGS MANAGEMENT ------------------------------>
	//e.POST("/api/sso/things/register", service.CreateThings)
	//<--------------------------- CHANELS MANAGEMENT ------------------------------>
	//e.POST("/api/sso/chanels/register", service.CreateChanels)
	//<--------------------------- MAP THINGS CHANELS ------------------------------>
	//e.POST("/api/sso/mapthingtochanel/register", service.CreateMapThingChanel)
	//<--------------------------- PUSH MESSAGE ------------------------------>
	e.POST("/api/mqtt/pushMessage", service.PushMessage)
	//<---------------------------SSO ------------------------------>
	//e.POST("/api/sso/requestToken", service.RequestSSOTokenv2)
	//e.POST("/api/sso/parseToken", service.ParseSSOToken)
	//<---------------------------SSO ------------------------------>
	//e.POST("/api/sso/publicchanel/register", service.CreateOrganizationPublicChanel)
	//e.Logger.Fatal(e.Start(":1323"))
	e.Logger.Fatal(e.StartAutoTLS(":443"))

}
