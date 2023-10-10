package data

import (
	"fmt"
	"furina/logger"
	"sync"
	"time"
)

const (
	UpdateSucceedMsg             = "获取角色面板数据成功"
	UpdateInternalServerErrorMsg = "获取角色面板数据失败, 服务不可用"
	UpdateIdNotExistErrorMsg     = "获取角色面板数据失败, ID错误"
)

type CharacterInfo struct {
	Cid           int
	Name          string
	Quality       int
	Constellation int
}

type User struct {
	Uid           string `bson:"_id"`
	UpdateAt      time.Time
	TTL           int
	Name          string
	CharacterList []CharacterInfo
}

var (
	UserCache     = map[string]User{} // 缓存
	UserCacheLock = sync.RWMutex{}
)

func GetUser(uid string) (user User) {
	user = getUserCache(uid)
	if user.Uid != "" {
		return
	}
	err := getDB().Get(TableNameUser, uid, &user)
	if err != nil {
		logger.Error("get user", "error", err)
		return
	}
	if user.Uid != "" {
		setUserCache(user)
	}
	return
}

type UpdateUserResp struct {
	User       User
	UpdateMsg  string
	UpdateList []int
}

func UpdateUser(uid string) (data UpdateUserResp) {
	user := getUserCache(uid)
	if t := time.Now().Sub(user.UpdateAt).Seconds(); t < float64(user.TTL) {
		data.UpdateMsg = fmt.Sprintf("更新失败, %.fs后才可再次更新数据", float64(user.TTL)-t)
		return
	}
	enkaData, err := getEnkaData(uid)
	if err != nil {
		logger.Error("UpdateUser", "get enka data error", err)
		data.UpdateMsg = UpdateInternalServerErrorMsg
		return
	}
	if enkaData.Uid == "" {
		logger.Error("UpdateUser", "uid错误", uid)
		data.UpdateMsg = UpdateIdNotExistErrorMsg
		return
	}

	var (
		now  = time.Now()
		has  = map[int]int{}
		list = getCharacterList(enkaData)
	)
	user.Uid = enkaData.Uid
	user.TTL = enkaData.TTL
	user.Name = enkaData.Nickname
	user.UpdateAt = now
	for i, c := range user.CharacterList {
		has[c.Cid] = i
	}
	fmt.Println(len(user.CharacterList))
	for _, c := range list {
		c.UpdateAt = now
		putCharacter(c)
		ci := CharacterInfo{
			Cid:           c.Cid,
			Name:          c.Name,
			Quality:       c.Quality,
			Constellation: c.Constellation,
		}
		if i, ok := has[c.Cid]; ok {
			user.CharacterList[i] = ci
		} else {
			user.CharacterList = append(user.CharacterList, ci)
		}
		data.UpdateList = append(data.UpdateList, c.Cid)
	}
	fmt.Println(len(user.CharacterList))
	putUser(user)
	data.User = user
	data.UpdateMsg = UpdateSucceedMsg
	return
}

func getUserCache(uid string) (user User) {
	UserCacheLock.RLock()
	defer UserCacheLock.RUnlock()
	user = UserCache[uid]
	user.CharacterList = make([]CharacterInfo, len(user.CharacterList))
	copy(user.CharacterList, UserCache[uid].CharacterList)
	return
}

func setUserCache(user User) {
	UserCacheLock.Lock()
	defer UserCacheLock.Unlock()
	UserCache[user.Uid] = user
}

func putUser(user User) {
	UserCacheLock.Lock()
	defer UserCacheLock.Unlock()
	UserCache[user.Uid] = user
	logger.Info("putUser", "updateTime", time.Now().String())
	err := getDB().Put(TableNameUser, user.Uid, user)
	if err != nil {
		logger.Error("put user", "error", err)
	}
}

func CheckUidValid(uid string) bool {
	if len(uid) != 9 {
		return false
	}
	for _, v := range uid {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
