package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFileHandler1(c echo.Context) error {
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	for _, file := range files {
		fmt.Println(file.Filename)
	}
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Files Missing!!! please provide valid files",
		})
	}

	instance, err := UploadBulkImage(files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong!",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    instance,
	})
	// return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))

}
