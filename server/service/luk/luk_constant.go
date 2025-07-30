package luk

import (
	"encoding/json"
	"fmt"
	"sort"

	common "github.com/flipped-aurora/gin-vue-admin/server/config"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun"
)

type LukConstantService struct {
	Keys []string `json:"keys" form:"keys"`
}

type DataResponseEnum struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type DataResponseEnums []DataResponseEnum

func (data DataResponseEnums) MarshalJSON() ([]byte, error) {
	if len(data) == 0 {
		return json.Marshal([]int{})
	} else {
		temp := []DataResponseEnum{}
		temp = data
		return json.Marshal(temp)
	}
}

func (data DataResponseEnums) Len() int {
	return len(data)
}

func (data DataResponseEnums) Less(i, j int) bool {
	iLabel := data[i].Label
	jLabel := data[j].Label
	labels := []string{iLabel, jLabel}
	sort.Strings(labels)
	return labels[0] == iLabel
}

func (data DataResponseEnums) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func (d *LukConstantService) GetLukConstant() (list map[string]DataResponseEnums) {
	list = make(map[string]DataResponseEnums)
	if len(d.Keys) == 0 {
		return
	}
	enumMapInstance := enum.GetEnumMapInstance()
	data := enumMapInstance.GetData(d.Keys)
	for key, enum := range data {
		allMember := enum.AllMember()
		keysSort := enum.KeysSort()
		if len(keysSort) == 0 { //没有指定顺序的
			for value, key2 := range allMember {
				if in, _ := common.InCollection(key2, enum.HiddenKeys()); in {
					continue
				}
				item := DataResponseEnum{
					Label: key2,
					Value: value,
				}
				if _, ok := list[key]; ok {
					list[key] = append(list[key], item)
				} else {
					list[key] = []DataResponseEnum{item}
				}
			}
			sort.Sort(list[key]) //按字符串排序，sort.sort()对数值字符串排序有坑只会按第一位数字排
		} else { //有指定顺序的
			keysNum := len(keysSort)
			memberNum := len(allMember)
			if keysNum != memberNum {
				fmt.Println("key的数量和map中的成员数量不一致")
			}
			reverseMap := map[string]interface{}{}
			for key, value := range allMember {
				reverseMap[value] = key
			}
			for _, key2 := range keysSort {
				if value, ok := reverseMap[key2]; ok {
					if in, _ := common.InCollection(key2, enum.HiddenKeys()); in {
						continue
					}
					item := DataResponseEnum{
						Label: key2,
						Value: value,
					}
					if _, ok := list[key]; ok {
						list[key] = append(list[key], item)
					} else {
						list[key] = []DataResponseEnum{item}
					}
				} else {
					fmt.Printf("key:%s不存在\n", key2)
				}
			}
		}
	}
	return
}
