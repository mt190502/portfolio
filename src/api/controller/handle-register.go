package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mt190502/portfolio/src/utils/database"
	"github.com/mt190502/portfolio/src/utils/localizator"
	"github.com/mt190502/portfolio/src/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

var RegisterCredential database.Credential

func HandleRegister(ctx *gin.Context) {
	siteTitle := database.DBConnection.First(&profile)

	var data = map[string]any{
		"Sitetitle": &profile.Sitetitle,
		"register":  localizator.Data["register"],
		"error":     localizator.Data["error"],
	}
	if siteTitle.Error != nil {
		logger.Log.Errorf("Failed to get site title from database: %v", siteTitle.Error)
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["internal"] + ": " + siteTitle.Error.Error(),
			},
		})
		return
	}

	user := database.DBConnection.First(&RegisterCredential)
	if user.Error == nil {
		ctx.Redirect(http.StatusFound, "/admin")
		return
	}

	ctx.HTML(http.StatusOK, "register.html", data)

}

func HandleRegisterPost(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	confirm := ctx.PostForm("confirm")

	RegisterCredential.Email = email

	if len(password) == 0 {
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["emptypass"],
			},
		})
		return
	}

	if password != confirm {
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["notmatchpass"],
			},
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Errorf("Failed to hash password: %v", err)
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["hasherr"] + ": " + err.Error(),
			},
		})
		return
	}
	RegisterCredential.Password = string(hash)

	result := database.DBConnection.Create(&RegisterCredential)
	if result.Error != nil {
		logger.Log.Errorf("Failed to create register credential: %v", result.Error)
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["acccreateerr"] + ": " + result.Error.Error(),
			},
		})
		return
	}

	ctx.Header("HX-Redirect", "/login")
	// ctx.Redirect(http.StatusMovedPermanently, "/login")
}
