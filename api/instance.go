package api

import (
	"github.com/allentom/haruka"
	"github.com/allentom/haruka/middleware"
	"github.com/rs/cors"
	"yousyncclient/config"
)

func RunAPIService() {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.UseMiddleware(middleware.NewLoggerMiddleware())
	e.UseMiddleware(middleware.NewPaginationMiddleware("page", "pageSize", 1, 20))
	e.Router.POST("/sync/folder", syncFileHandler)
	e.RunAndListen(config.Instance.Addr)
}
