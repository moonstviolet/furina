package render

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"

	"furina/config"
	"furina/data"
	"furina/logger"
)

func Routes(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(config.GetBaseDir(), "render", "templates", "*"))
	r.GET("/", index)
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.GET("/user/:uid/profile", userProfile)
	r.POST("/user/:uid/update", userUpdate)
	r.GET("/user/:uid/character/:cid", characterDetail)
}

func index(ctx *gin.Context) {
	uid, err := ctx.Cookie("uid")
	if err != nil || uid == "" {
		ctx.HTML(http.StatusOK, "index.html", nil)
		return
	}
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/user/%s/profile", uid))
}

func login(ctx *gin.Context) {
	uid := ctx.Request.FormValue("uid")
	if uid == "" {
		ctx.Redirect(http.StatusFound, "/")
	}
	ctx.SetCookie("uid", uid, 0, "", "", false, true)
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/user/%s/profile", uid))
}

func logout(ctx *gin.Context) {
	ctx.SetCookie("uid", "", -1, "", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

func userProfile(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user := data.GetUser(uid)
	if user.Uid == "" {
		user = data.UpdateUser(uid)
	}
	if user.Uid == "" {
		ctx.String(http.StatusOK, "错误的id")
		return
	}
	data.CleanUserUpdateMsg(uid)
	ctx.HTML(http.StatusOK, "profile.html", getUserProfileView(user))
}

func userUpdate(ctx *gin.Context) {
	uid := ctx.Param("uid")
	data.UpdateUser(uid)
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/user/%s/profile", uid))
}

func characterDetail(ctx *gin.Context) {
	uidStr, cidStr := ctx.Param("uid"), ctx.Param("cid")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		logger.Error("Character Detail", "error", err)
	}
	ctx.HTML(
		http.StatusOK, "character.html", getCharacterViewData(data.GetCharacter(uidStr, cid)),
	)
}

type ArtifactCalReq struct {
	Type int
}

func artifact(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		ctx.HTML(http.StatusOK, "artifact.html", nil)
	} else if ctx.Request.Method == http.MethodPost {
		fmt.Println(ctx.PostForm("type"))
	}
}
