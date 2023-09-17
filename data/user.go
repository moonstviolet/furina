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
)

type CharacterInfo struct {
	Cid             int
	Name            string
	Quality         int
	Constellation   int
	UpdatedRecently bool `json:"-"`
}

type User struct {
	Uid           string `bson:"_id"`
	UpdateAt      time.Time
	UpdateMsg     string `json:"-" bson:"-"`
	TTL           int
	Name          string
	NewUpdate     bool `json:"-" bson:"-"`
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
	} else {
		user = UpdateUser(uid)
	}
	return
}

func UpdateUser(uid string) (user User) {
	user = getUserCache(uid)
	if t := time.Now().Sub(user.UpdateAt).Seconds(); t < float64(user.TTL) {
		user.UpdateMsg = fmt.Sprintf("更新失败, %.fs后才可再次更新数据", t)
		setUserCache(user)
		return
	}
	enkaData, err := getEnkaData(uid)
	if err != nil {
		logger.Error("get enka data", "error", err)
		user.UpdateMsg = UpdateInternalServerErrorMsg
		setUserCache(user)
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
	user.UpdateMsg = UpdateSucceedMsg
	if len(list) > 0 {
		user.NewUpdate = true
	}
	for i, c := range user.CharacterList {
		has[c.Cid] = i
	}
	for _, c := range list {
		c.UpdateAt = now
		putCharacter(c)
		ci := CharacterInfo{
			Cid:             c.Cid,
			Name:            c.Name,
			Quality:         c.Quality,
			Constellation:   c.Constellation,
			UpdatedRecently: true,
		}
		if i, ok := has[c.Cid]; ok {
			user.CharacterList[i] = ci
		} else {
			user.CharacterList = append(user.CharacterList, ci)
		}
	}
	putUser(user)
	return
}

func CleanUserUpdateMsg(uid string) {
	user := getUserCache(uid)
	user.UpdateMsg = ""
	user.NewUpdate = false
	for i := range user.CharacterList {
		user.CharacterList[i].UpdatedRecently = false
	}
	setUserCache(user)
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
	err := getDB().Put(TableNameUser, user.Uid, user)
	if err != nil {
		logger.Error("put user", "error", err)
	}
}
