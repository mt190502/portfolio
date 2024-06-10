package controller

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/mt190502/portfolio/src/utils/database"
	"github.com/mt190502/portfolio/src/utils/localizator"
	"golang.org/x/crypto/bcrypt"
)

func ChangeValues(column string, val1 string, val2 string, ctx *gin.Context) {
	database.DBConnection.First(&credential)

	if len(val1) == 0 {
		return
	}

	if val1 != val2 {
		ctx.HTML(http.StatusOK, "error.html", map[string]any{
			"error": map[string]string{
				"error": localizator.Data["error"]["error"],
				"title": localizator.Data["error"]["valnotmatch"],
			},
		})
		return
	}

	if column == "email" {
		database.DBConnection.Model(&database.Credential{}).Where("id = ?", 1).Update(column, val2)
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(val1), bcrypt.DefaultCost)
		if err != nil {
			ctx.HTML(http.StatusOK, "error.html", map[string]any{
				"error": map[string]string{
					"error": localizator.Data["error"]["error"],
					"title": localizator.Data["error"]["hasherr"] + ": " + err.Error(),
				},
			})
			return
		}
		database.DBConnection.Model(&database.Credential{}).Where("id = ?", 1).Update(column, string(hash))
	}

}

func CheckURL(url string) string {
	regex := regexp.MustCompile(`([a-zA-Z0-9._-]+)$`)
	return regex.FindString(url)
}

func HandleAdmin(ctx *gin.Context) {
	store := ginsession.FromContext(ctx)
	_, session := store.Get("user")
	if !session {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	if ctx.Request.Method == "POST" {
		switch ctx.PostForm("type") {
		case "email":
			val1 := ctx.PostForm("new-email")
			val2 := ctx.PostForm("confirm-email")
			ChangeValues("email", val1, val2, ctx)
		case "password":
			val1 := ctx.PostForm("new-password")
			val2 := ctx.PostForm("confirm-password")
			ChangeValues("password", val1, val2, ctx)
		case "settings":
			database.DBConnection.First(&profile)

			var values = map[string]string{
				"Sitetitle": ctx.PostForm("sitename"),
				"Name":      ctx.PostForm("username"),
				"Phone":     ctx.PostForm("phone"),
				"Address":   ctx.PostForm("address"),
				"About":     ctx.PostForm("about"),
				"Email":     ctx.PostForm("email"),
				"Linkedin":  CheckURL(ctx.PostForm("linkedin")),
				"Github":    CheckURL(ctx.PostForm("github")),
				"Instagram": CheckURL(ctx.PostForm("instagram")),
				"Twitter":   CheckURL(ctx.PostForm("twitter")),
				"Facebook":  CheckURL(ctx.PostForm("facebook")),
			}

			for key, value := range values {
				if len(value) == 0 {
					continue
				}

				database.DBConnection.Model(&database.Profile{}).Where("id = ?", 1).Update(key, value)
			}
		}
	}
	var data = map[string]interface{}{
		"Sitetitle": profile.Sitetitle + " " + localizator.Data["admin"]["admin_panel"],
		"admin":     localizator.Data["admin"],
		"error":     localizator.Data["error"],
	}
	ctx.HTML(http.StatusOK, "admin.html", data)

}
