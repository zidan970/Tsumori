package main

import (
	"zidan/clean-arch/app/configs"
	"zidan/clean-arch/app/databases"
	"zidan/clean-arch/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbSql := databases.InitDBMysql(cfg)
	databases.InitialMigration(dbSql)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.Logger())
	// // e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// // 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// // }))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	routers.InitRouter(dbSql, e)

	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
