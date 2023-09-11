package render

import (
	"fmt"
	"furina/config"
	"furina/data"
	"furina/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

func Routes(r *gin.Engine) {
	r.LoadHTMLGlob(filepath.Join(config.GetBaseDir(), "render", "templates", "*"))
	r.GET("/", index)
	r.GET("/user/:uid/profile", userProfile)
	r.POST("/user/:uid/update", userUpdate)
	r.GET("/user/:uid/character/:cid", characterDetail)
	r.GET("/artifact", artifact)
	r.POST("/artifact", artifact)
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", data.UserList)
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
