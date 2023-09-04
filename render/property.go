package render

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"time"

	"furina/data"
)

type PropertyView struct {
	Desc   string // 名称
	Key    string
	Value  string // 值
	Base   string // 值
	Plus   string // 值
	Number string // 词条数
	Weight int    // 权重
	Count  int    // 强化次数
}

var (
	PropertyViewKeyList = []string{
		data.PropKey_Hp,
		data.PropKey_Atk,
		data.PropKey_Def,
		data.PropKey_Mastery,
		data.PropKey_CritRate,
		data.PropKey_CritDmg,
		data.PropKey_Recharge,
		data.PropKey_DmgBonus,
	}
	PropertyViewDescMap = map[string]string{
		data.PropKey_Hp:       "生命值",
		data.PropKey_Atk:      "攻击力",
		data.PropKey_Def:      "防御力",
		data.PropKey_Mastery:  "元素精通",
		data.PropKey_CritRate: "暴击率",
		data.PropKey_CritDmg:  "暴击伤害",
		data.PropKey_Crit:     "双暴",
		data.PropKey_Recharge: "充能效率",
		data.PropKey_DmgBonus: "伤害加成",
	}
)

func getPropertyView(key string, p data.Property, w int) PropertyView {
	pv := PropertyView{
		Desc: PropertyViewDescMap[key],
		Key:  key,
	}
	switch key {
	case data.PropKey_Hp, data.PropKey_Atk, data.PropKey_Def:
		pv.Base = formatFloat(p.Base)
		pv.Plus = formatFloat(p.Plus)
		pv.Value = formatFloat(p.Value)
	case data.PropKey_Mastery:
		pv.Value = formatFloat(p.Value)
	case data.PropKey_CritRate, data.PropKey_CritDmg, data.PropKey_Recharge, data.PropKey_DmgBonus:
		pv.Value = formatPercent(p.Value)
	default:
		pv.Value = formatFloat(p.Value)
	}
	pv.Weight = w
	return pv
}

func formatFloat(f float64) string {
	return message.NewPrinter(language.Und).Sprint(number.Decimal(f, number.Scale(1)))
}

func formatPercent(f float64) string {
	return message.NewPrinter(language.Und).Sprint(number.Percent(f, number.Scale(1)))
}

func formatTime(t time.Time) string {
	return t.Format("01-02 15:04")
}
