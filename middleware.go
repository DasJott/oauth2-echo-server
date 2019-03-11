package echoserver

import (
	"github.com/labstack/echo"
)

// Config for middleware
type Config struct {
	// keys stored in the context
	TokenKey string
	// defines a function to skip middleware.Returning true skips processing
	// the middleware.
	Skipper func(echo.Context) bool
}

var (
	// DefaultConfig is the default middleware config.
	DefaultConfig = Config{
		TokenKey: "token",
		Skipper: func(_ echo.Context) bool {
			return false
		},
	}
)

// TokenHandler gets the token from request using default config
func TokenHandler() echo.MiddlewareFunc {
	return TokenHandlerWithConfig(&DefaultConfig)
}

// TokenHandlerWithConfig gets the token from request with given config
func TokenHandlerWithConfig(cfg *Config) echo.MiddlewareFunc {
	tokenKey := cfg.TokenKey
	if tokenKey == "" {
		tokenKey = DefaultConfig.TokenKey
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cfg.Skipper != nil && cfg.Skipper(c) {
				return next(c)
			}
			ti, err := eServer.ValidationBearerToken(c.Request())
			if err != nil {
				return err
			}

			c.Set(tokenKey, ti)
			return next(c)
		}
	}
}
