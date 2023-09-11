package data

import (
	"furina/logger"
	"log"
	"strconv"
	"sync"
	"time"
)

const (
	Element_Anemo   = "anemo"
	Element_Cryo    = "cryo"
	Element_Dendro  = "dendro"
	Element_Electro = "electro"
	Element_Geo     = "geo"
	Element_Hydro   = "hydro"
	Element_Pyro    = "pyro"

	Weapon_Bow      = "bow"
	Weapon_Catalyst = "catalyst"
	Weapon_Claymore = "claymore"
	Weapon_Polearm  = "polearm"
	Weapon_Sword    = "sword"
)

type Talent struct {
	Key           string
	Constellation int // +3命座
	BaseLevel     int // 原始等级
	Level         int // 实际等级
}

type Weapon struct {
	Id         int
	Name       string
	Level      int
	Quality    int
	Refinement int
}

type Character struct {
	Id             string              `bson:"_id"`
	Uid            string              // 游戏id
	UpdateAt       time.Time           // 数据更新时间
	Cid            int                 // 角色id
	Name           string              // 名称
	Quality        int                 // 星级
	Element        string              // 元素
	WeaponType     string              // 武器类型
	Constellation  int                 // 命座
	Level          int                 // 等级
	TalentMap      map[string]Talent   // 天赋
	PropertyMap    map[string]Property // 战斗属性
	PropertyWeight MSI                 // 属性权重
	Weapon         Weapon
	ArtifactStat   ArtifactStat // 圣遗物总评
	ArtifactList   []Artifact   // 圣遗物列表
}

var (
	CharacterCache     = map[string]Character{} // 缓存
	CharacterCacheLock = sync.RWMutex{}
)

func GetCharacter(uid string, cid int) (c Character) {
	id := getIdByPrefixAndKey(uid, strconv.Itoa(cid))
	CharacterCacheLock.RLock()
	c, ok := CharacterCache[id]
	CharacterCacheLock.RUnlock()
	if ok {
		return
	}
	err := getDB().Get(TableNameCharacter, id, &c)
	if err != nil {
		logger.Error("get character", "error", err)
	}
	if c.Cid != 0 {
		CharacterCacheLock.Lock()
		CharacterCache[id] = c
		CharacterCacheLock.Unlock()
	}
	return
}

func putCharacter(c Character) {
	CharacterCacheLock.Lock()
	defer CharacterCacheLock.Unlock()
	CharacterCache[c.Id] = c
	err := getDB().Put(TableNameCharacter, c.Id, c)
	if err != nil {
		logger.Error("put user", "error", err)
	}
}

func getCharacterList(e *EnkaData) (list []Character) {
	for _, info := range e.AvatarInfoList {
		meta := getCharacterMetaById(info.AvatarID)
		c := Character{
			Id:             getIdByPrefixAndKey(e.Uid, strconv.Itoa(info.AvatarID)),
			Uid:            e.Uid,
			Cid:            info.AvatarID,
			Name:           meta.Name,
			Quality:        meta.Quality,
			Element:        meta.Element,
			WeaponType:     meta.WeaponType,
			Constellation:  len(info.TalentIdList),
			Level:          stringToInt(info.PropMap[PropType_Level].Val),
			TalentMap:      getTalentMapByAvatarInfo(&info, &meta),
			PropertyMap:    getPropertyMapByAvatarInfo(&info, &meta),
			PropertyWeight: meta.PropertyWeight,
			Weapon:         getWeaponByEquipList(info.EquipList),
			ArtifactList:   getArtifactListByEquipList(info.EquipList),
		}
		c.CalArtifactScore()
		list = append(list, c)
	}
	return
}

func getWeaponByEquipList(l []Equip) Weapon {
	for _, e := range l {
		if e.Flat.ItemType == ITEM_WEAPON {
			w := Weapon{
				Id:      e.ItemId,
				Name:    getWeaponNameById(e.ItemId),
				Level:   e.Weapon.Level,
				Quality: e.Flat.RankLevel,
			}
			for _, v := range e.Weapon.AffixMap {
				w.Refinement = v + 1
				break
			}
			return w
		}
	}
	return Weapon{}
}

func getTalentMapByAvatarInfo(info *AvatarInfo, meta *CharacterMeta) map[string]Talent {
	m := map[string]Talent{}
	for id, baseLevel := range info.SkillLevelMap {
		var (
			key          = meta.TalentId[id]
			level        = baseLevel
			constell, ok = meta.TalentConstell[key]
		)
		if ok && len(info.TalentIdList) >= constell {
			level += 3
		}
		m[key] = Talent{
			Key:           key,
			Constellation: constell,
			BaseLevel:     baseLevel,
			Level:         level,
		}
	}
	return m
}

func (c *Character) setPropertyWeight(w MSI) {
	c.PropertyWeight = w
	c.CalArtifactScore()
}

func stringToInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatalln(err)
	}
	return int(num)
}
