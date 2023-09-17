package data

import (
	"fmt"
	"log"
	"math"
	"sort"
)

const (
	ArtifactMaxScore       = 66.0
	ArtifactPerNumberScore = 46.62 / 6
	ArtifactRatingCount    = 9
)

var (
	ArtifactRatingStr = [ArtifactRatingCount]string{"D", "C", "B", "A", "S", "SS", "SSS", "ACE", "ACE²"}
	propGrowthMap     = map[string][4]float64{
		PropKey_Hp:       {209.13, 239, 268.88, 298.75},
		PropKey_HpPct:    {0.048, 0.0466, 0.0525, 0.0583},
		PropKey_Atk:      {13.62, 15.56, 17.51, 19.45},
		PropKey_AtkPct:   {0.048, 0.0466, 0.0525, 0.0583},
		PropKey_Def:      {16.2, 18.52, 20.83, 23.15},
		PropKey_DefPct:   {0.051, 0.0583, 0.0656, 0.0729},
		PropKey_Mastery:  {16.32, 18.65, 20.98, 23.31},
		PropKey_CritRate: {0.027, 0.031, 0.035, 0.0389},
		PropKey_CritDmg:  {0.0544, 0.0622, 0.0699, 0.0777},
		PropKey_Recharge: {0.0453, 0.0518, 0.0583, 0.0648},

		PropKey_HealingBonus:     {3: 0.0449},
		PropKey_PhysicalDmgBonus: {3: 0.0729},
		PropKey_PyroDmgBonus:     {3: 0.0583},
		PropKey_ElectroDmgBonus:  {3: 0.0583},
		PropKey_HydroDmgBonus:    {3: 0.0583},
		PropKey_DendroDmgBonus:   {3: 0.0583},
		PropKey_AnemoDmgBonus:    {3: 0.0583},
		PropKey_GeoDmgBonus:      {3: 0.0583},
		PropKey_CryoDmgBonus:     {3: 0.0583},
	}
)

func getPropGrowthMax(key string) float64 {
	v, ok := propGrowthMap[key]
	if !ok {
		log.Fatalln("invalid key")
	}
	return v[3]
}

func getPropGrowthAverage(key string) float64 {
	v, ok := propGrowthMap[key]
	if !ok {
		log.Fatalln("invalid key")
	}
	return (v[0] + v[1] + v[2] + v[3]) / 4
}

func (c *Character) CalArtifactScore() {
	c.ArtifactStat = ArtifactStat{
		StatNumber: map[string]float64{},
	}
	for idx := range c.ArtifactList {
		c.ArtifactList[idx].CalScore(c.PropertyWeight, c.PropertyMap)
		c.ArtifactStat.Score += c.ArtifactList[idx].Score
		for k, v := range c.ArtifactList[idx].StatNumber {
			c.ArtifactStat.StatNumber[k] += v
		}
	}
	c.ArtifactStat.Desc = fmt.Sprintf("评分规则：%s-通用", c.Name)
	c.ArtifactStat.Rating = getArtifactRating(c.ArtifactStat.Score / 5)
}

func (a *Artifact) CalScore(w MSI, p map[string]Property) {
	a.Score = 0
	a.StatNumber = map[string]float64{}
	weight := MSI{}
	for k, v := range w {
		if k == PropKey_Hp || k == PropKey_Atk || k == PropKey_Def {
			if p != nil {
				weight[k] = int((getPropGrowthMax(k) / p[k].Base) * (float64(v) / getPropGrowthMax(k+PropKey_Pct)))
			} else {
				// 没有指定角色时按1/3权重计算
				weight[k] = v / 3
			}
			weight[k+PropKey_Pct] = v
		} else {
			weight[k] = v
		}
	}
	for i, prop := range a.AppendPropList {
		if prop.Key == PropKey_Hp || prop.Key == PropKey_Atk || prop.Key == PropKey_Def {
			// 显示的权值, 大小值统一
			prop.Weight = weight[prop.Key+PropKey_Pct]
		} else {
			prop.Weight = weight[prop.Key]
		}
		// 展示词条, 取average
		prop.Number = prop.Value / getPropGrowthAverage(prop.Key)
		if weight[prop.Key] > 0 {
			key, number := prop.Key, prop.Number
			if key == PropKey_Hp || key == PropKey_Atk || key == PropKey_Def {
				number *= float64(weight[key]) / float64(weight[key+PropKey_Pct])
			} else if key == PropKey_HpPct || key == PropKey_AtkPct || key == PropKey_DefPct {
				key = key[:len(key)-3]
			} else if key == PropKey_CritRate || key == PropKey_CritDmg {
				key = PropKey_Crit
			}
			a.StatNumber[key] += number
		}
		// 计算分数, 取max
		calNumber := prop.Value / getPropGrowthMax(prop.Key)
		prop.Count = int(math.Ceil(calNumber - 0.1))
		prop.Score = ArtifactPerNumberScore * calNumber * float64(weight[prop.Key]) / 100

		a.AppendPropList[i] = prop
		a.Score += prop.Score
	}

	// 对权重进行排序, 计算理论最高分
	var (
		s        []int
		maxScore float64
		cnt      int
	)
	for k, v := range weight {
		// 排除圣遗物主词条和不可能出现的副词条
		if k == a.MainProp.Key || isDmgBonus(k) || k == PropKey_HealingBonus {
			continue
		}

		s = append(s, v)
	}
	sort.Slice(
		s, func(i, j int) bool {
			return s[i] > s[j]
		},
	)
	for _, v := range s {
		if cnt == 0 {
			maxScore += ArtifactPerNumberScore * 6 * float64(v) / 100
		} else if cnt > 3 {
			break
		} else {
			maxScore += ArtifactPerNumberScore * float64(v) / 100
		}
		cnt++
	}

	// 主词条附加分
	if a.Type == ArtifactType_Sands || a.Type == ArtifactType_Goblet || a.Type == ArtifactType_Circlet {
		key := a.MainProp.Key
		addScore := ArtifactPerNumberScore * a.MainProp.Value / getPropGrowthMax(key) * float64(weight[key]) / 100 * 0.25
		a.Score += addScore
		maxScore += addScore
	}
	// 评分归一化
	if maxScore != 0 {
		a.Score = a.Score / maxScore * ArtifactMaxScore
	}
	a.Rating = getArtifactRating(a.Score)
}

func getArtifactRating(score float64) string {
	r := int(score / 7)
	if r >= ArtifactRatingCount {
		r = ArtifactRatingCount - 1
	}
	if r < 0 {
		r = 0
	}
	return ArtifactRatingStr[r]
}
