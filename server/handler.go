package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"furina/data"
	"furina/errorCode"
	"furina/logger"
)

func index(ctx *gin.Context) {
	if getUserId(ctx) == "" {
		ctx.HTML(http.StatusOK, "index.html", nil)
		return
	}
	ctx.Redirect(http.StatusFound, "/user/profile")
}

func login(
	req *struct {Uid string},
	resp *struct{},
	ctx *gin.Context,
) *errorCode.RespError {
	uid := req.Uid
	if !data.CheckUidValid(uid) {
		return errorCode.NewError(errorCode.CODE_PARAM_WRONG, "错误的UID")
	}
	ctx.SetCookie("uid", uid, 0, "", "", false, true)
	return nil
}

func logout(ctx *gin.Context) {
	ctx.SetCookie("uid", "", -1, "", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

// 用户
func userProfile(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "profile.html", nil)
}

func UserProfile(
	req *struct{ Update bool },
	resp *struct {
		User       UserProfileView
		UpdateMsg  string
		UpdateList []int
	},
	ctx *gin.Context,
) *errorCode.RespError {
	uid, _ := ctx.Cookie("uid")
	var user data.User
	if req.Update {
		t := data.UpdateUser(uid)
		user = t.User
		resp.UpdateMsg = t.UpdateMsg
		resp.UpdateList = t.UpdateList
	} else {
		user = data.GetUser(uid)
	}
	if resp.UpdateMsg != "" && resp.UpdateMsg != data.UpdateSucceedMsg {
		return nil
	}
	resp.User = getUserProfileView(user)
	return nil
}

// 角色
func characterDetail(ctx *gin.Context) {
	cidStr := ctx.Param("cid")
	_, err := strconv.Atoi(cidStr)
	if err != nil {
		logger.Error("Character Detail", "error", err)
	}
	ctx.HTML(http.StatusOK, "character.html", nil)
}

func CharacterDetail(
	req *struct{ Cid int },
	resp *struct {
		Character CharacterView
	},
	ctx *gin.Context,
) *errorCode.RespError {
	uid := getUserId(ctx)
	char := data.GetCharacter(uid, req.Cid)
	resp.Character = getCharacterViewData(char)
	return nil
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
