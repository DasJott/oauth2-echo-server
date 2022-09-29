package main

import (
	"net/http"

	echoserver "github.com/dasjott/oauth2-echo-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/labstack/echo/v4"
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
