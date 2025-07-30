/*
* email: oyblog@qq.com
* Author:  oy
* Date:    2021/5/25 下午2:15
 */
package enum

type Enum interface {
	IsMember(key interface{}) bool     //判断是否是枚举成员(要转成对应的类型才能判断成功)
	AllMember() map[interface{}]string //获取枚举里面的所有枚举成员key是值value是名称
	HiddenKeys() []string              //隐藏哪些key
}

type SortEnum interface {
	Enum
	KeysSort() []string //指定key的顺序按指定key的顺序对值进行排序
}

type SortEnumGetText interface {
	SortEnum
	GetText(value interface{}) string //值转文字说明
}
