package server

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"furina/config"
	"furina/middleware"
)

func routes(r *gin.Engine) {
	r.Delims("[[", "]]")
	r.LoadHTMLGlob(filepath.Join(config.GetBaseDir(), "server", "templates", "*"))

	r.GET("/", index)
	r.POST("/login", middleware.CreateHandlerFunc(login))

	needLogin := r.Group("")
	needLogin.Use(logined)
	{
		needLogin.GET("/logout", logout)

		needLogin.GET("/user/profile", userProfile)
		needLogin.Handle(http.MethodPost, "/user/profile", middleware.CreateHandlerFunc(UserProfile))

		needLogin.GET("/character/:cid", characterDetail)
		needLogin.Handle(http.MethodPost, "/character", middleware.CreateHandlerFunc(CharacterDetail))
	}

	r.GET("/artifact", artifact)

	r.POST("/artifact", artifact)
}

func getUserId(ctx *gin.Context) string {
	uid, err := ctx.Cookie("uid")
	if err != nil {
		return ""
	}
	return uid
}

func logined(ctx *gin.Context) {
	if getUserId(ctx) == "" {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
}
