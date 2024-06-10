package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/mt190502/portfolio/src/utils/database"
	"github.com/mt190502/portfolio/src/utils/localizator"
	"github.com/mt190502/portfolio/src/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

var LoginCredential database.Credential

func HandleLogin(ctx *gin.Context) {
	var data = map[string]any{
		"Sitetitle": &profile.Sitetitle,
		"login":     localizator.Data["login"],
		"error":     localizator.Data["error"],
	}

	store := ginsession.FromContext(ctx)
	_, session := store.Get("user")
	if session {
		ctx.Redirect(http.StatusMovedPermanently, "/admin")
		return
	} else {
		ctx.HTML(http.StatusOK, "login.html", data)
	}
}

func HandleLoginPost(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	credential := database.DBConnection.First(&LoginCredential, "email = ?", email)
	if credential.Error != nil {
		logger.Log.Errorf("Failed to get credential from database: %v", credential.Error)
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["invalid"],
			},
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(LoginCredential.Password), []byte(password))
	if err != nil {
		logger.Log.Errorf("Failed to compare password: %v", err)
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["invalid"],
			},
		})
		return
	} else {
		store := ginsession.FromContext(ctx)
		store.Set("user", LoginCredential.Email)
		if store.Save() != nil {
			logger.Log.Errorf("Failed to save session: %v", err)
			ctx.HTML(http.StatusOK, "error.html", map[string]any{
				"error": map[string]string{
					"error": localizator.Data["error"]["error"],
					"title": localizator.Data["error"]["sesssaveerr"],
				},
			})
			return
		}

		ctx.Header("HX-Redirect", "/admin")
		// ctx.Redirect(http.StatusMovedPermanently, "/admin")
	}
}
