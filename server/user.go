package server

import (
	"furina/data"
)

type UserProfileView struct {
	Uid           string
	UpdateAt      string
	CharacterList []data.CharacterInfo
}

func getUserProfileView(user data.User) UserProfileView {
	return UserProfileView{
		Uid:           user.Uid,
		UpdateAt:      formatTime(user.UpdateAt),
		CharacterList: user.CharacterList,
	}
}
