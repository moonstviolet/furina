package data

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	PropKey_Hp       = "Hp"
	PropKey_HpPct    = "HpPct"
	PropKey_Atk      = "Atk"
	PropKey_AtkPct   = "AtkPct"
	PropKey_Def      = "Def"
	PropKey_DefPct   = "DefPct"
	PropKey_Mastery  = "Mastery"
	PropKey_CritRate = "CritRate"
	PropKey_CritDmg  = "CritDmg"
	PropKey_Recharge = "Recharge"

	PropKey_Pct      = "Pct"
	PropKey_Crit     = "Crit"
	PropKey_DmgBonus = "DmgBonus"

	PropKey_PhysicalDmgBonus = "PhysicalDmgBonus"
	PropKey_PhysicalRes      = "PhysicalRes"
	PropKey_PyroDmgBonus     = "PyroDmgBonus"
	PropKey_PyroRes          = "PyroRes"
	PropKey_ElectroDmgBonus  = "ElectroDmgBonus"
	PropKey_ElectroRes       = "ElectroRes"
	PropKey_HydroDmgBonus    = "HydroDmgBonus"
	PropKey_HydroRes         = "HydroRes"
	PropKey_DendroDmgBonus   = "DendroDmgBonus"
	PropKey_DendroRes        = "DendroRes"
	PropKey_AnemoDmgBonus    = "AnemoDmgBonus"
	PropKey_AnemoRes         = "AnemoRes"
	PropKey_GeoDmgBonus      = "GeoDmgBonus"
	PropKey_GeoRes           = "GeoRes"
	PropKey_CryoDmgBonus     = "CryoDmgBonus"
	PropKey_CryoRes          = "CryoRes"

	PropKey_HealingBonus = "HealingBonus"
	PropKey_HealedBonus  = "HealedBonus"
)

type Property struct {
	Key string // 索引

	Base    float64 // 白值
	Flat    float64 // 小值
	Percent float64 // 百分比
	Plus    float64 // 附加值
	Value   float64 // 总值
	Weight  int     // 权重

	Score  float64 // 评分
	Number float64 // 词条数(最大)
	Count  int     // 强化次数 todo
}

var (
	propIdKeyMap = map[string]string{
		"FIGHT_PROP_ATTACK":            PropKey_Atk,
		"FIGHT_PROP_ATTACK_PERCENT":    PropKey_AtkPct,
		"FIGHT_PROP_CHARGE_EFFICIENCY": PropKey_Recharge,
		"FIGHT_PROP_CRITICAL":          PropKey_CritRate,
		"FIGHT_PROP_CRITICAL_HURT":     PropKey_CritDmg,
		"FIGHT_PROP_DEFENSE":           PropKey_Def,
		"FIGHT_PROP_DEFENSE_PERCENT":   PropKey_DefPct,
		"FIGHT_PROP_ELEC_ADD_HURT":     PropKey_ElectroDmgBonus,
		"FIGHT_PROP_ELEC_SUB_HURT":     PropKey_ElectroRes,
		"FIGHT_PROP_ELEMENT_MASTERY":   PropKey_Mastery,
		"FIGHT_PROP_FIRE_ADD_HURT":     PropKey_PyroDmgBonus,
		"FIGHT_PROP_FIRE_SUB_HURT":     PropKey_PyroRes,
		"FIGHT_PROP_GRASS_ADD_HURT":    PropKey_DendroDmgBonus,
		"FIGHT_PROP_GRASS_SUB_HURT":    PropKey_DendroRes,
		"FIGHT_PROP_HEAL_ADD":          PropKey_HealingBonus,
		"FIGHT_PROP_HEALED_ADD":        PropKey_HealedBonus,
		"FIGHT_PROP_HP":                PropKey_Hp,
		"FIGHT_PROP_HP_PERCENT":        PropKey_HpPct,
		"FIGHT_PROP_ICE_ADD_HURT":      PropKey_CryoDmgBonus,
		"FIGHT_PROP_ICE_SUB_HURT":      PropKey_CryoRes,
		"FIGHT_PROP_PHYSICAL_ADD_HURT": PropKey_PhysicalDmgBonus,
		"FIGHT_PROP_PHYSICAL_SUB_HURT": PropKey_PhysicalRes,
		"FIGHT_PROP_ROCK_ADD_HURT":     PropKey_GeoDmgBonus,
		"FIGHT_PROP_ROCK_SUB_HURT":     PropKey_GeoRes,
		"FIGHT_PROP_WATER_ADD_HURT":    PropKey_HydroDmgBonus,
		"FIGHT_PROP_WATER_SUB_HURT":    PropKey_HydroRes,
		"FIGHT_PROP_WIND_ADD_HURT":     PropKey_AnemoDmgBonus,
		"FIGHT_PROP_WIND_SUB_HURT":     PropKey_AnemoRes,
	}
)

func getPropKeyById(id string) string {
	return propIdKeyMap[id]
}

func getPropertyMapByAvatarInfo(info *AvatarInfo, meta *CharacterMeta) map[string]Property {
	m := map[string]Property{
		PropKey_Hp: {
			Key:     PropKey_Hp,
			Base:    info.FightPropMap[FightPropType_HPBase],
			Flat:    info.FightPropMap[FightPropType_HPFlat],
			Percent: info.FightPropMap[FightPropType_HPPercent],
			Value:   info.FightPropMap[FightPropType_HP],
			Plus:    info.FightPropMap[FightPropType_HP] - info.FightPropMap[FightPropType_HPBase],
		},
		PropKey_Atk: {
			Key:     PropKey_Atk,
			Base:    info.FightPropMap[FightPropType_ATKBase],
			Flat:    info.FightPropMap[FightPropType_ATKFlat],
			Percent: info.FightPropMap[FightPropType_ATKPercent],
			Value:   info.FightPropMap[FightPropType_ATK],
			Plus:    info.FightPropMap[FightPropType_ATK] - info.FightPropMap[FightPropType_ATKBase],
		},
		PropKey_Def: {
			Key:     PropKey_Def,
			Base:    info.FightPropMap[FightPropType_DEFBase],
			Flat:    info.FightPropMap[FightPropType_DEFFlat],
			Percent: info.FightPropMap[FightPropType_DEFPercent],
			Value:   info.FightPropMap[FightPropType_DEF],
			Plus:    info.FightPropMap[FightPropType_DEF] - info.FightPropMap[FightPropType_DEFBase],
		},
		PropKey_Mastery: {
			Key:   PropKey_Mastery,
			Value: info.FightPropMap[FightPropType_ElementalMastery],
		},
		PropKey_CritRate: {
			Key:   PropKey_CritRate,
			Base:  0.05,
			Value: info.FightPropMap[FightPropType_CRITRate],
		},
		PropKey_CritDmg: {
			Key:   PropKey_CritDmg,
			Base:  0.50,
			Value: info.FightPropMap[FightPropType_CRITDMG],
		},
		PropKey_Recharge: {
			Key:   PropKey_Recharge,
			Base:  1,
			Value: info.FightPropMap[FightPropType_EnergyRecharge],
		},
		PropKey_PhysicalDmgBonus: {
			Key:   PropKey_PhysicalDmgBonus,
			Value: info.FightPropMap[FightPropType_PhysicalDMGBonus],
		},
		PropKey_PyroDmgBonus: {
			Key:   PropKey_PyroDmgBonus,
			Value: info.FightPropMap[FightPropType_PyroDMGBonus],
		},
		PropKey_ElectroDmgBonus: {
			Key:   PropKey_ElectroDmgBonus,
			Value: info.FightPropMap[FightPropType_ElectroDMGBonus],
		},
		PropKey_HydroDmgBonus: {
			Key:   PropKey_HydroDmgBonus,
			Value: info.FightPropMap[FightPropType_HydroDMGBonus],
		},
		PropKey_DendroDmgBonus: {
			Key:   PropKey_DendroDmgBonus,
			Value: info.FightPropMap[FightPropType_DendroDMGBonus],
		},
		PropKey_AnemoDmgBonus: {
			Key:   PropKey_AnemoDmgBonus,
			Value: info.FightPropMap[FightPropType_AnemoDMGBonus],
		},
		PropKey_GeoDmgBonus: {
			Key:   PropKey_GeoDmgBonus,
			Value: info.FightPropMap[FightPropType_GeoDMGBonus],
		},
		PropKey_CryoDmgBonus: {
			Key:   PropKey_CryoDmgBonus,
			Value: info.FightPropMap[FightPropType_CryoDMGBonus],
		},
	}
	return m
}

func isPctPropperty(p string) bool {
	for _, v := range []string{
		PropKey_HpPct,
		PropKey_AtkPct,
		PropKey_DefPct,
		PropKey_CritRate,
		PropKey_CritDmg,
		PropKey_Recharge,
		PropKey_PhysicalDmgBonus,
		PropKey_PyroDmgBonus,
		PropKey_ElectroDmgBonus,
		PropKey_HydroDmgBonus,
		PropKey_DendroDmgBonus,
		PropKey_AnemoDmgBonus,
		PropKey_GeoDmgBonus,
		PropKey_CryoDmgBonus,
		PropKey_HealingBonus,
	} {
		if p == v {
			return true
		}
	}
	return false
}

func isPropWeightKey(p string) bool {
	for _, v := range []string{
		PropKey_Hp,
		PropKey_Atk,
		PropKey_Def,
		PropKey_Mastery,
		PropKey_CritRate,
		PropKey_CritDmg,
		PropKey_Recharge,
		PropKey_PhysicalDmgBonus,
		PropKey_PyroDmgBonus,
		PropKey_ElectroDmgBonus,
		PropKey_HydroDmgBonus,
		PropKey_DendroDmgBonus,
		PropKey_AnemoDmgBonus,
		PropKey_GeoDmgBonus,
		PropKey_CryoDmgBonus,
		PropKey_HealingBonus,
	} {
		if p == v {
			return true
		}
	}
	return false
}

func isDmgBonus(p string) bool {
	for _, v := range []string{
		PropKey_PhysicalDmgBonus,
		PropKey_PyroDmgBonus,
		PropKey_ElectroDmgBonus,
		PropKey_HydroDmgBonus,
		PropKey_DendroDmgBonus,
		PropKey_AnemoDmgBonus,
		PropKey_GeoDmgBonus,
		PropKey_CryoDmgBonus,
	} {
		if p == v {
			return true
		}
	}
	return false
}

func (c *Character) GetDmgBonusKey() string {
	for k := range c.PropertyWeight {
		if isDmgBonus(k) {
			return k
		}
	}
	return cases.Title(language.Und).String(c.Element) + PropKey_DmgBonus
}
