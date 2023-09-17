package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	EnkaUrlFormat = "https://enka.network/api/uid/%s"
	UserAgent     = "furina"
)

const (
	PropType_Level = "4001"

	FightPropType_HPBase     = "1"
	FightPropType_HPFlat     = "2"
	FightPropType_HPPercent  = "3"
	FightPropType_HP         = "2000"
	FightPropType_ATKBase    = "4"
	FightPropType_ATKFlat    = "5"
	FightPropType_ATKPercent = "6"
	FightPropType_ATK        = "2001"
	FightPropType_DEFBase    = "7"
	FightPropType_DEFFlat    = "8"
	FightPropType_DEFPercent = "9"
	FightPropType_DEF        = "2002"

	FightPropType_CRITRate         = "20"
	FightPropType_CRITDMG          = "22"
	FightPropType_EnergyRecharge   = "23"
	FightPropType_ElementalMastery = "28"

	FightPropType_PhysicalDMGBonus = "30"
	FightPropType_PyroDMGBonus     = "40"
	FightPropType_ElectroDMGBonus  = "41"
	FightPropType_HydroDMGBonus    = "42"
	FightPropType_DendroDMGBonus   = "43"
	FightPropType_AnemoDMGBonus    = "44"
	FightPropType_GeoDMGBonus      = "45"
	FightPropType_CryoDMGBonus     = "46"

	//10    Base SPD
	//11    SPD%
	//26    Healing Bonus
	//27    Incoming Healing Bonus
	//29    Physical RES
	//50    Pyro RES
	//51    Electro RES
	//52    Hydro RES
	//53    Dendro RES
	//54    Anemo RES
	//55    Geo RES
	//56    Cryo RES
	//70    Pyro Enegry Cost
	//71    Electro Energy Cost
	//72    Hydro Energy Cost
	//73    Dendro Energy Cost
	//74    Anemo Energy Cost
	//75    Cryo Energy Cost
	//76    Geo Energy Cost
	//80    Cooldown reduction
	//81    Shield Strength
	//1000    Current Pyro Energy
	//1001    Current Electro Energy
	//1002    Current Hydro Energy
	//1003    Current Dendro Energy
	//1004    Current Anemo Energy
	//1005    Current Cryo Energy
	//1006    Current Geo Energy
	//1010    Current HP
	//2003    SPD
	//3025    Elemental reaction CRIT Rate
	//3026    Elemental reaction CRIT DMG
	//3027    Elemental reaction (Overloaded) CRIT Rate
	//3028    Elemental reaction (Overloaded) CRIT DMG
	//3029    Elemental reaction (Swirl) CRIT Rate
	//3030    Elemental reaction (Swirl) CRIT DMG
	//3031    Elemental reaction (Electro-Charged) CRIT Rate
	//3032    Elemental reaction (Electro-Charged) CRIT DMG
	//3033    Elemental reaction (Superconduct) CRIT Rate
	//3034    Elemental reaction (Superconduct) CRIT DMG
	//3035    Elemental reaction (Burn) CRIT Rate
	//3036    Elemental reaction (Burn) CRIT DMG
	//3037    Elemental reaction (Frozen (Shattered)) CRIT Rate
	//3038    Elemental reaction (Frozen (Shattered)) CRIT DMG
	//3039    Elemental reaction (Bloom) CRIT Rate
	//3040    Elemental reaction (Bloom) CRIT DMG
	//3041    Elemental reaction (Burgeon) CRIT Rate
	//3042    Elemental reaction (Burgeon) CRIT DMG
	//3043    Elemental reaction (Hyperbloom) CRIT Rate
	//3044    Elemental reaction (Hyperbloom) CRIT DMG
	//3045    Base Elemental reaction CRIT Rate
	//3046    Base Elemental reaction CRIT DMG

	ITEM_WEAPON    = "ITEM_WEAPON"
	ITEM_RELIQUARY = "ITEM_RELIQUARY"
)

type EnkaData struct {
	EnkaPlayerInfo `json:"playerInfo"`
	AvatarInfoList []AvatarInfo `json:"avatarInfoList"`
	TTL            int          `json:"ttl"` // 缓存刷新秒数
	Uid            string       `json:"uid"`
}

type EnkaPlayerInfo struct {
	Nickname             string `json:"nickname"`             // 昵称
	Level                int    `json:"level"`                // 等级
	Signature            string `json:"signature"`            // 签名
	WorldLevel           int    `json:"worldLevel"`           // 世界等级
	NameCardId           int    `json:"nameCardId"`           // 资料名片 ID
	FinishAchievementNum int    `json:"finishAchievementNum"` // 已解锁成就数
	TowerFloorIndex      int    `json:"towerFloorIndex"`      // 本期深境螺旋层数
	TowerLevelIndex      int    `json:"towerLevelIndex"`      // 本期深境螺旋间数
	ShowAvatarInfoList   []struct {
		AvatarId  int `json:"avatarId"`  // 角色 ID
		Level     int `json:"level"`     // 角色等级
		CostumeId int `json:"costumeId"` // 角色衣装 ID
	} `json:"showAvatarInfoList"` // 展示角色列表
	ShowNameCardIdList []int `json:"showNameCardIdList"` // 展示名片 ID 列表
	ProfilePicture     struct {
		AvatarId int `json:"avatarId"`
	} `json:"profilePicture"` // 头像角色 ID
}

type AvatarInfo struct {
	AvatarID     int   `json:"avatarId"`     // 角色 ID
	TalentIdList []int `json:"talentIdList"` // 命之座 ID 列表, 如果未解锁任何命之座则此数据不存在
	PropMap      map[string]struct {
		Type int    `json:"type"` // 属性类型
		Ival string `json:"ival"`
		Val  string `json:"val"` // 属性值
	} `json:"propMap"` // 角色属性列表
	FightPropMap           map[string]float64 `json:"fightPropMap"`           // 角色战斗属性 Map
	SkillDepotId           int                `json:"skillDepotId"`           // 角色天赋 ID
	InherentProudSkillList []int              `json:"inherentProudSkillList"` // 固定天赋列表 ID
	SkillLevelMap          map[int]int        `json:"SkillLevelMap"`          // 天赋等级
	EquipList              []Equip            `json:"EquipList"`              // 装备列表：武器、圣遗物
	FetterInfo             struct {
		ExpLevel int `json:"expLevel"` // 角色好感等级
	} `json:"fetterInfo"`
}

type Equip struct {
	ItemId int `json:"itemId"` // 装备 ID
	Weapon struct {
		Level        int         `json:"level"`        // 武器等级
		PromoteLevel int         `json:"promoteLevel"` // 武器突破等级
		AffixMap     map[int]int `json:"affixMap"`     // 武器精炼等级 [0-4], key: Cid
	} `json:"weapon"` //武器基本信息, 只有武器有
	Reliquary struct {
		Level            int   `json:"level"`            // 圣遗物等级 [1-21]
		MainPropId       int   `json:"mainPropId"`       // 圣遗物主属性 ID
		AppendPropIdList []int `json:"appendPropIdList"` // 圣遗物副属性 ID 列表
	} `json:"reliquary"` // 圣遗物基本信息, 只有圣遗物有
	Flat struct {
		NameTextMapHash    string `json:"nameTextMapHash"`    // 装备名的哈希值
		SetNameTextMapHash string `json:"setNameTextMapHash"` // 圣遗物套装的名称的哈希值
		RankLevel          int    `json:"rankLevel"`          // 装备稀有度
		ReliquaryMainstat  struct {
			MainPropId string  `json:"mainPropId"` // 装备属性名称
			StatValue  float64 `json:"statValue"`  // 属性值
		} `json:"reliquaryMainstat"` // 圣遗物主属性
		ReliquarySubstats []struct {
			AppendPropId string  `json:"appendPropId"` // 装备属性名称
			StatValue    float64 `json:"statValue"`    // 属性值
		} `json:"reliquarySubstats"` // 圣遗物副属性列表
		WeaponStats []struct {
			AppendPropId string  `json:"appendPropId"` // 装备属性名称
			StatValue    float64 `json:"statValue"`    // 属性值
		} `json:"weaponStats"`

		ItemType  string `json:"itemType"`  // 装备类别
		Icon      string `json:"icon"`      // 装备图标名称
		EquipType string `json:"equipType"` // 圣遗物类型
	} `json:"flat"` // 装备详细信息
}

func getEnkaData(uid string) (data *EnkaData, err error) {
	url := fmt.Sprintf(EnkaUrlFormat, uid)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", UserAgent)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusBadRequest:
		err = errors.New("Wrong UID format")
	case 404:
		err = errors.New("Player does not exist (MHY server said that)")
	case 424:
		err = errors.New("Game maintenance / everything is broken after the game update")
	case 429:
		err = errors.New("Rate-limited (either by my server or by MHY server)")
	case 500:
		err = errors.New("General server error")
	case 503:
		err = errors.New("I screwed up massively")
	default:
		err = errors.New(resp.Status)
	}
	if resp.StatusCode != http.StatusOK {
		return
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var res EnkaData
	err = json.Unmarshal(bytes, &res)
	return &res, err
}
