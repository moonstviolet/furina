package render

import (
	"errors"
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
	r.GET("/artifact", artifact)
	r.POST("/artifact", artifact)
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
	if !data.CheckUidValid(uid) {
		ctx.String(http.StatusBadRequest, "错误的id")
		return
	}
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
	if !data.CheckUidValid(uid) {
		ctx.String(http.StatusBadRequest, "错误的id")
		return
	}
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
	if !data.CheckUidValid(uid) {
		ctx.String(http.StatusBadRequest, "错误的id")
		return
	}
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

func artifact(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		ctx.HTML(http.StatusOK, "artifact.html", getArtifactTemplate(ArtifactView{}))
	} else if ctx.Request.Method == http.MethodPost {
		var (
			arti           data.Artifact
			propertyWeight = map[string]int{}
		)
		err := func() error {
			t, err := strconv.Atoi(ctx.PostForm("Type"))
			if err != nil {
				return err
			}
			if t < 1 || t > 5 {
				return errors.New("错误的属性")
			}
			arti.Type = t
			arti.Name = []string{"生之花", "死之羽", "时之沙", "空之杯", "理之冠"}[t-1]

			arti.MainProp = data.Property{
				Key: ctx.PostForm("MainPropKey"),
			}
			val, err := strconv.ParseFloat(ctx.PostForm("MainPropValue"), 64)
			if err != nil {
				return err
			}
			arti.MainProp.Value = val

			keyList, valueList := ctx.PostFormArray("AppendPropKey"), ctx.PostFormArray("AppendPropValue")
			if len(keyList) != len(valueList) || len(keyList) > 4 {
				return errors.New("错误的属性")
			}
			for i := 0; i < len(keyList); i++ {
				appendProp := data.Property{
					Key: keyList[i],
				}
				val, err = strconv.ParseFloat(valueList[i], 64)
				if err != nil {
					return err
				}
				appendProp.Value = val
				arti.AppendPropList = append(arti.AppendPropList, appendProp)
			}

			for k, v := range ctx.PostFormMap("PropertyWeight") {
				if len(v) == 0 {
					continue
				}
				weight, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				propertyWeight[k] = weight
			}
			return nil
		}()
		if err != nil {
			logger.Error("圣遗物手动计算参数错误", "error", err)
			ctx.String(http.StatusBadRequest, "错误的请求")
			return
		}
		// todo 校验
		arti.CalScore(propertyWeight, nil)
		ctx.HTML(
			http.StatusOK, "artifact.html",
			getArtifactTemplate(getArtifactView(arti)),
		)
	}
}
