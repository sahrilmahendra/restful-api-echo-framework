package controllers

import (
	"echo-rest/helper"
	"echo-rest/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "login success")
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
