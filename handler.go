package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func UploadSingleFileHandler(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("failed to get file: %v", err)
		return c.JSON(http.StatusBadRequest, &Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    fmt.Sprintf("failed to get file: %v", strings.Split(err.Error(), "\n")[0]),
		})
	}
	if file == nil {
		return c.JSON(http.StatusBadRequest, &Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Files Missing!!! please provide valid files",
		})
	}
	instance, err := UploadSingleFile(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
		})
	}
	log.Printf("uploaded successfully %s files with fields name=%s and email=%s.", file.Filename, name, email)
	return c.JSON(http.StatusOK, &Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       instance,
	})
}

func UploadBulkFileHandler(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, &Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Files Missing!!! please provide valid files",
		})
	}

	instance, err := UploadBulkFile(files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    err.Error(),
		})
	}
	log.Printf("uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email)
	return c.JSON(http.StatusOK, &Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       instance,
	})
}
