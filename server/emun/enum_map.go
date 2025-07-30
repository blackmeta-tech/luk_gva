/*
* email: oyblog@qq.com
* Author:  oy
* Date:    2021/5/27 下午2:42
 */
package enum

import (
	"fmt"
	"sync"
)

var enumMapInstance *enumMap
var once sync.Once

type enumMap struct {
	Data map[string]SortEnum
}

// 单列模式 协程安全
func GetEnumMapInstance() *enumMap {
	once.Do(func() {
		enumMapInstance = new(enumMap)
	})
	return enumMapInstance
}

// 判断是否已挂载过
func (this *enumMap) IsMounted(key string) bool {
	if this.Data == nil {
		return false
	}
	if _, ok := this.Data[key]; ok {
		return true
	}
	return false
}

// 挂载，如果挂载已存在的key会抛panic
func (this *enumMap) Mount(key string, value SortEnum) {
	if this.Data == nil {
		this.Data = map[string]SortEnum{}
	}
	if _, ok := this.Data[key]; ok {
		panic(fmt.Sprintf("该key已存在请使用别的key，key:%s", key))
	}
	this.Data[key] = value

}

//强制挂载，如果存在会覆盖
func (this *enumMap) ForceMount(key string, value SortEnum) {
	if this.Data == nil {
		this.Data = map[string]SortEnum{}
	}
	this.Data[key] = value
}

//卸载
func (this *enumMap) UnMount(key string) {
	delete(this.Data, key)
}

//获取数据
func (this *enumMap) GetData(keys []string) map[string]SortEnum {
	result := map[string]SortEnum{}
	if len(keys) == 0 {
		return this.Data
	}
	for _, key := range keys {
		if value, ok := this.Data[key]; ok {
			result[key] = value
		} else {
			fmt.Println("key不存在：", key)
		}
	}
	return result
}
