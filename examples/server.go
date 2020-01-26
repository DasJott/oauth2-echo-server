package main

import (
	"net/http"

	echoserver "github.com/dasjott/oauth2-echo-server"
	"github.com/labstack/echo"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	echoserver.InitServer(manager)
	echoserver.SetAllowGetAccessRequest(true)
	echoserver.SetClientInfoHandler(server.ClientFormHandler)

	e := echo.New()

	auth := e.Group("/oauth2")
	{
		auth.GET("/token", echoserver.HandleTokenRequest)
	}

	api := e.Group("/api")
	{
		api.Use(echoserver.TokenHandler())
		api.GET("/test", func(c echo.Context) error {
			ti := c.Get(echoserver.DefaultConfig.TokenKey)
			if ti != "" {
				return c.JSON(http.StatusOK, ti)
			}
			return echo.NewHTTPError(http.StatusNotFound, "token not found")
		})
	}

	e.Start(":9096")
}
