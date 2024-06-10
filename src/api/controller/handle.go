package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mt190502/portfolio/src/utils/database"
)

var profile database.Profile
var credential database.Credential

func Handle(handler string) func(*gin.Context) {
	database.DBConnection.First(&profile)
	database.DBConnection.First(&credential)
	switch handler {
	case "login":
		return HandleLogin
	case "root":
		return HandleRoot
	case "admin":
		return HandleAdmin
	case "register":
		return HandleRegister
	default:
		return nil
	}
}
