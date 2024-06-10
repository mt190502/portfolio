package controller

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/mt190502/portfolio/src/utils/database"
	"github.com/mt190502/portfolio/src/utils/localizator"
)

func HandleRoot(ctx *gin.Context) {
	database.DBConnection.First(&profile)

	profileMap := structs.Map(profile)

	var data = map[string]any{
		"profile": profileMap,
		"index":   localizator.Data["index"],
	}

	ctx.HTML(200, "index.html", data)
}
