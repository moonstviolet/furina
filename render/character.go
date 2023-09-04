package render

import (
	"furina/config"
	"furina/data"
)

const (
	TalentKeyA = "a"
	TalentKeyE = "e"
	TalentKeyQ = "q"
)

type CharacterView struct {
	Uid           string // uid
	Cid           int    // 角色Id
	UpdateAt      string
	Name          string
	Element       string // 元素
	WeaponType    string // 武器类型
	Level         int    // 等级
	Constellation int    // 命座
	ConstellCount []int
	TalentList    []data.Talent
	PropertyList  []PropertyView
	Weapon        data.Weapon
	ArtifactStat  ArtifactStatView
	ArtifactList  []ArtifactView
	Version       config.VersionConfig
}

func getCharacterViewData(c data.Character) CharacterView {
	cv := CharacterView{
		Uid:           c.Uid,
		Cid:           c.Cid,
		UpdateAt:      formatTime(c.UpdateAt),
		Name:          c.Name,
		Element:       c.Element,
		WeaponType:    c.WeaponType,
		Level:         c.Level,
		Constellation: c.Constellation,
		ConstellCount: []int{1, 2, 3, 4, 5, 6},
		TalentList:    []data.Talent{c.TalentMap[TalentKeyA], c.TalentMap[TalentKeyE], c.TalentMap[TalentKeyQ]},
		Weapon:        c.Weapon,
		ArtifactStat:  getArtifactStatView(c.ArtifactStat),
		Version:       config.GetConfig().Version,
	}
	for _, key := range PropertyViewKeyList {
		propertyKey := key
		if key == data.PropKey_DmgBonus {
			propertyKey = c.GetDmgBonusKey()
		}
		cv.PropertyList = append(
			cv.PropertyList, getPropertyView(key, c.PropertyMap[propertyKey], c.PropertyWeight[propertyKey]),
		)
	}
	for _, v := range c.ArtifactList {
		cv.ArtifactList = append(cv.ArtifactList, getArtifactView(v))
	}
	return cv
}
