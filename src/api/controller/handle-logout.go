package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/mt190502/portfolio/src/utils/logger"
)

func HandleLogout(ctx *gin.Context) {
	store := ginsession.FromContext(ctx)
	store.Delete("user")
	if store.Save() != nil {
		logger.Log.Errorf("Failed to delete session")
		ctx.HTML(http.StatusOK, "error.html", "Internal server error")
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/")
}
