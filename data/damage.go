package data

import (
	"log"
)

type DamageCharacterInfo struct {
	Level            int
	Base             float64
	DmgBonus         float64
	CritRate         float64
	CritDamage       float64
	ResistanceReduce float64
	DefenceIgnore    float64
	DefenceReduce    float64
	Mastery          float64
}

type DamageMonsterInfo struct {
	Level      int
	Resistance float64
}

type Damage struct {
	CharacterInfo DamageCharacterInfo
	MonsterInfo   DamageMonsterInfo
	ZoneMap       map[string]float64
}

func NewDamage(c DamageCharacterInfo, m DamageMonsterInfo) *Damage {
	return &Damage{
		CharacterInfo: c,
		MonsterInfo:   m,
	}
}

func test()  {

}

func (d *Damage) Cal() (average, max float64) {
	// 基础区
	res := d.CharacterInfo.Base
	log.Println("base", d.CharacterInfo.Base)
	// 增伤区
	res *= 1 + d.CharacterInfo.DmgBonus
	log.Println("dmgBonus", 1+d.CharacterInfo.DmgBonus)
	// 反应区
	// 防御区
	t := (float64(d.CharacterInfo.Level) + 100) /
		((float64(d.CharacterInfo.Level) + 100) +
			(1-d.CharacterInfo.DefenceIgnore)*(1-d.CharacterInfo.DefenceReduce)*(float64(d.MonsterInfo.Level)+100))
	res *= t
	log.Println("defence", t)
	// 抗性区
	if r := d.MonsterInfo.Resistance - d.CharacterInfo.ResistanceReduce; r > 0.75 {
		res *= 1 / (1 + 4*r)
	} else if r >= 0 {
		res *= 1 - r
	} else {
		res *= 1 - r/2
	}
	// 暴击区
	if d.CharacterInfo.CritRate > 1 {
		d.CharacterInfo.CritRate = 1
	}
	average = res * (1 + d.CharacterInfo.CritRate*d.CharacterInfo.CritDamage)
	max = res * (1 + d.CharacterInfo.CritDamage)
	log.Println("average", average)
	log.Println("max", max)
	return
}
