package render

import (
	"sort"

	"furina/data"
)

type StatNumberView struct {
	Key    string
	value  float64
	Number string
}

type ArtifactStatView struct {
	Desc       string           // 描述
	Score      string           // 评分
	Rating     string           // 评级
	StatNumber []StatNumberView // 词条
}

type ArtifactView struct {
	Name           string // 名称
	Set            string // 套装
	Type           int    // 部位
	Level          int    // 等级
	Score          string
	Rating         string
	MainProp       PropertyView
	AppendPropList []PropertyView
}

var (
	ArtifactPropertyViewDescMap = map[string]string{
		data.PropKey_Hp:               "小生命",
		data.PropKey_HpPct:            "大生命",
		data.PropKey_Atk:              "小攻击",
		data.PropKey_AtkPct:           "大攻击",
		data.PropKey_Def:              "小防御",
		data.PropKey_DefPct:           "大防御",
		data.PropKey_Mastery:          "元素精通",
		data.PropKey_CritRate:         "暴击率",
		data.PropKey_CritDmg:          "暴击伤害",
		data.PropKey_Recharge:         "充能效率",
		data.PropKey_PhysicalDmgBonus: "物伤加成",
		data.PropKey_PyroDmgBonus:     "火伤加成",
		data.PropKey_ElectroDmgBonus:  "雷伤加成",
		data.PropKey_HydroDmgBonus:    "水伤加成",
		data.PropKey_DendroDmgBonus:   "草伤加成",
		data.PropKey_AnemoDmgBonus:    "风伤加成",
		data.PropKey_GeoDmgBonus:      "岩伤加成",
		data.PropKey_CryoDmgBonus:     "冰伤加成",
		data.PropKey_HealingBonus:     "治疗加成",
	}
)

func getArtifactStatView(as data.ArtifactStat) ArtifactStatView {
	asv := ArtifactStatView{
		Desc:   as.Desc,
		Score:  formatFloat(as.Score),
		Rating: as.Rating,
	}
	sum := 0.0
	for k, v := range as.StatNumber {
		asv.StatNumber = append(
			asv.StatNumber, StatNumberView{Key: PropertyViewDescMap[k], value: v, Number: formatFloat(v)},
		)
		sum += v
	}
	asv.StatNumber = append(
		asv.StatNumber, StatNumberView{Key: "总计", value: sum, Number: formatFloat(sum)},
	)
	sort.Slice(
		asv.StatNumber, func(i, j int) bool {
			return asv.StatNumber[i].value > asv.StatNumber[j].value
		},
	)
	return asv
}

func getArtifactView(a data.Artifact) ArtifactView {
	av := ArtifactView{
		Name:     a.Name,
		Set:      a.Set,
		Type:     a.Type,
		Level:    a.Level,
		Score:    formatFloat(a.Score),
		Rating:   a.Rating,
		MainProp: getArtifactPropertyView(a.MainProp),
	}
	for _, v := range a.AppendPropList {
		av.AppendPropList = append(av.AppendPropList, getArtifactPropertyView(v))
	}
	return av
}

func getArtifactPropertyView(p data.Property) PropertyView {
	pv := PropertyView{
		Desc:   ArtifactPropertyViewDescMap[p.Key],
		Weight: p.Weight,
		Number: formatFloat(p.Number),
		Count:  p.Count,
	}
	switch p.Key {
	case data.PropKey_Hp, data.PropKey_Atk, data.PropKey_Def, data.PropKey_Mastery:
		pv.Value = formatFloat(p.Value)
	default:
		pv.Value = formatPercent(p.Value)
	}
	return pv
}
