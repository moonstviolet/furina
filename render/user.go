package render

import (
	"fmt"
	"furina/data"
)

type UserProfileView struct {
	Uid           string
	UpdateAt      string
	UpdateMsg     string
	NewUpdate     bool
	CharacterList []data.CharacterInfo
}

func getUserProfileView(user data.User) UserProfileView {
	fmt.Println(user)
	return UserProfileView{
		Uid:           user.Uid,
		UpdateAt:      formatTime(user.UpdateAt),
		UpdateMsg:     user.UpdateMsg,
		NewUpdate:     user.NewUpdate,
		CharacterList: user.CharacterList,
	}
}
