package data

const (
	EQUIPType_Flower  = "EQUIP_BRACER"   // 花
	EQUIPType_Plume   = "EQUIP_NECKLACE" // 羽
	EQUIPType_Sands   = "EQUIP_SHOES"    // 沙
	EQUIPType_Goblet  = "EQUIP_RING"     // 杯
	EQUIPType_Circlet = "EQUIP_DRESS"    // 头

	ArtifactType_Flower  = 1 // 花
	ArtifactType_Plume   = 2 // 羽
	ArtifactType_Sands   = 3 // 沙
	ArtifactType_Goblet  = 4 // 杯
	ArtifactType_Circlet = 5 // 头

)

var (
	ArtifactTypeMap = map[string]int{
		EQUIPType_Flower:  ArtifactType_Flower,
		EQUIPType_Plume:   ArtifactType_Plume,
		EQUIPType_Sands:   ArtifactType_Sands,
		EQUIPType_Goblet:  ArtifactType_Goblet,
		EQUIPType_Circlet: ArtifactType_Circlet,
	}
)

type Artifact struct {
	Name           string // 名称
	Set            string // 套装
	Type           int    // 部位
	Quality        int    // 星级
	Level          int    // 等级
	StatNumber     map[string]float64
	Score          float64 // 评分
	Rating         string  // 等级
	MainProp       Property
	AppendPropList []Property
}

type ArtifactStat struct {
	Desc       string // 描述
	StatNumber map[string]float64
	Score      float64 // 评分
	Rating     string  // 评级
}

func getArtifactListByEquipList(l []Equip) (list []Artifact) {
	for _, e := range l {
		if e.Flat.ItemType == ITEM_RELIQUARY {
			a := Artifact{
				Set:     ArtifactSetNameMap[e.Flat.SetNameTextMapHash],
				Type:    ArtifactTypeMap[e.Flat.EquipType],
				Quality: e.Flat.RankLevel,
				Level:   e.Reliquary.Level - 1,
				MainProp: Property{
					Key:   getPropKeyById(e.Flat.ReliquaryMainstat.MainPropId),
					Value: e.Flat.ReliquaryMainstat.StatValue,
				},
			}
			a.Name = ArtifactMetaMap[a.Set].SetNameList[a.Type-1]
			if isPctPropperty(a.MainProp.Key) {
				a.MainProp.Value /= 100
			}
			for i, v := range e.Flat.ReliquarySubstats {
				a.AppendPropList = append(
					a.AppendPropList, Property{
						Key:   getPropKeyById(v.AppendPropId),
						Value: v.StatValue,
					},
				)
				if isPctPropperty(a.AppendPropList[i].Key) {
					a.AppendPropList[i].Value /= 100
				}
			}
			list = append(list, a)
		}
	}
	return
}

func getArtifactMainPropListByType(t int) (list []string) {
	switch t {
	case ArtifactType_Sands:
		list = []string{
			PropKey_HpPct,
			PropKey_AtkPct,
			PropKey_DefPct,
			PropKey_Mastery,
			PropKey_Recharge,
		}
	case ArtifactType_Goblet:
		list = []string{
			PropKey_HpPct,
			PropKey_AtkPct,
			PropKey_DefPct,
			PropKey_Mastery,
			PropKey_PhysicalDmgBonus,
			PropKey_PyroDmgBonus,
			PropKey_ElectroDmgBonus,
			PropKey_HydroDmgBonus,
			PropKey_DendroDmgBonus,
			PropKey_AnemoDmgBonus,
			PropKey_GeoDmgBonus,
			PropKey_CryoDmgBonus,
		}
	case ArtifactType_Circlet:
		list = []string{
			PropKey_HpPct,
			PropKey_AtkPct,
			PropKey_DefPct,
			PropKey_Mastery,
			PropKey_CritRate,
			PropKey_CritDmg,
			PropKey_HealingBonus,
		}
	}
	return
}
