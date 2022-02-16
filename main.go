package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", UploadSingleFileHandler)
	e.POST("/upload/bulk", UploadBulkFileHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
