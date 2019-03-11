package echoserver

import (
	"sync"

	"github.com/labstack/echo"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
)

var (
	eServer *server.Server
	once    sync.Once
)

// InitServer Initialize the service
func InitServer(manager oauth2.Manager) *server.Server {
	once.Do(func() {
		eServer = server.NewDefaultServer(manager)
	})
	return eServer
}

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c echo.Context) error {
	return eServer.HandleAuthorizeRequest(c.Response().Writer, c.Request())
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c echo.Context) error {
	return eServer.HandleTokenRequest(c.Response().Writer, c.Request())
}
