package data

import (
	"encoding/json"
	"log"
	"os"
)

type CharacterMeta struct {
	Cid            int
	Name           string
	Quality        int
	Element        string
	WeaponType     string
	TalentId       map[int]string
	TalentConstell map[string]int
	PropertyWeight map[string]int
}

type ArtifactMeta struct {
	Name        string
	SetNameList []string
}

var (
	CharacterIdToNameMap = map[int]string{}          // id->name
	CharacterMetaMap     = map[int]CharacterMeta{}   // id->characterMeta
	WeaponIdToNameMap    = map[int]string{}          // id->name
	ArtifactSetNameMap   = map[string]string{}       // hash->name
	ArtifactMetaMap      = map[string]ArtifactMeta{} // name->artifactMeta
)

func init() {
	// 角色名称
	err := readDataFromFile(getCharacterNameMapFile(), &CharacterIdToNameMap)
	if err != nil {
		log.Fatalln(err)
	}
	// 角色属性权重
	weightMap := map[string]map[string]int{} // Name->Weight
	err = readDataFromFile(getCharacterPropertyWeightMapFile(), &weightMap)
	for name, weight := range weightMap {
		for k, v := range weight {
			if isPropWeightKey(k) == false || v < 0 || v > 100 {
				log.Fatalf("%s属性权重配置有误, %v: %v\n", name, k, v)
			}
		}
	}
	if err != nil {
		log.Fatalln(err)
	}
	// 角色基础信息
	for id, name := range CharacterIdToNameMap {
		CharacterMetaMap[id] = func() CharacterMeta {
			var mt struct {
				Star       int            `json:"star"`
				Elem       string         `json:"elem"`
				Weapon     string         `json:"weapon"`
				TalentId   map[int]string `json:"talentId"`
				TalentCons map[string]int `json:"talentCons"`
			}
			err = readDataFromFile(getCharacterDataFileByName(name), &mt)
			if err != nil {
				log.Fatalln(err)
			}
			cm := CharacterMeta{
				Cid:            id,
				Name:           name,
				Quality:        mt.Star,
				Element:        mt.Elem,
				WeaponType:     mt.Weapon,
				TalentId:       map[int]string{},
				TalentConstell: map[string]int{},
				PropertyWeight: weightMap[name],
			}
			for k, v := range mt.TalentId {
				cm.TalentId[k] = v
			}
			for k, v := range mt.TalentCons {
				cm.TalentConstell[k] = v
			}
			return cm
		}()
	}
	// 武器名称
	err = readDataFromFile(getWeaponNameMapFile(), &WeaponIdToNameMap)
	if err != nil {
		log.Fatalln(err)
	}
	// 圣遗物套装名称
	err = readDataFromFile(getArtifactSetNameMapFile(), &ArtifactSetNameMap)
	if err != nil {
		log.Fatalln(err)
	}
	// 圣遗物基础信息
	m := map[string]struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Sets map[int]struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"sets"`
	}{}
	err = readDataFromFile(getArtifactDataFile(), &m)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range m {
		ArtifactMetaMap[v.Name] = ArtifactMeta{
			Name: v.Name,
			SetNameList: []string{
				v.Sets[1].Name,
				v.Sets[2].Name,
				v.Sets[3].Name,
				v.Sets[4].Name,
				v.Sets[5].Name,
			},
		}
	}
}

func getCharacterMetaById(id int) CharacterMeta {
	return CharacterMetaMap[id]
}

func getWeaponNameById(id int) string {
	return WeaponIdToNameMap[id]
}

func readDataFromFile(path string, v any) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}
