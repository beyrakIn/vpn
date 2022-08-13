package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	m "service/middleware"
)

var (
	logger   = m.Logger{}
	config   = &Config{}
	cmd      = "create-vpn"
	userPath = "user.json"
)

func init() {
	openJsonFile()
}

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.LogErr(errors.New("Recovered in /add handler " + fmt.Sprintf("%v", err)))
			}
		}()

		var d User
		err := c.ShouldBindJSON(&d)
		logger.LogErr(err)

		auth := config.Users[0]

		if d.Username != auth.Username || d.Password != auth.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		result := command()

		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	}
}

func command() string {
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		logger.LogErr(err)
	}
	return string(out)
}

func openJsonFile() {
	file, err := os.ReadFile(userPath)
	if err != nil {
		logger.LogErr(err)
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		logger.LogErr(err)
	}

}

type Config struct {
	Users []User `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
