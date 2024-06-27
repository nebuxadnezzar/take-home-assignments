package route

import (
	"log"
	db "maas/api/database"
	"maas/api/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoWrapperRouter struct {
	*echo.Echo
}

func NewEchoWrapperRouter(userdbpath string) (*EchoWrapperRouter, map[string]interface{}) {
	ec := echo.New()
	custom := map[string]interface{}{}
	memesdb, err := db.NewMemesDatabase()
	if err != nil {
		log.Fatalf("failed to create memes db: %v", err)
	}
	usersdb, err := db.NewUsersDb(userdbpath)
	if err != nil {
		log.Fatalf("failed to create user db: %v", err)
	}
	custom[db.MEMESDB_KEY] = memesdb
	custom[db.USERSDB_KEY] = usersdb
	ec.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("custom", custom)
			return next(c)
		}
	})

	ec.Use(middleware.Logger())

	ec.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, BASE_ROUTE+HEALTH_ROUTE)
	})
	base := ec.Group(BASE_ROUTE)
	base.Group(HEALTH_ROUTE).GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	base.Group(MEMES_ROUTE).GET("", handler.SendMeme)
	base.Group(USER_STATS_ROUTE).GET("", handler.SendStats)

	return &EchoWrapperRouter{ec}, custom
}
