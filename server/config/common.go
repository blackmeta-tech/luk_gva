package config

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/wumansgy/goEncrypt"
)

const (
	FORMAT_DATE_CST     = "2006-01-02"
	FORMAT_DATETIME_CST = "2006-01-02 15:04:05"
	FORMAT_DATEHOUR_CST = "2006-01-02 15:00:00"
	EMPTY_USERADDRESS   = "0x0000000000000000000000000000000000000000" //空地址

)

/*
* 判断元素是否在集合中
* @param value       元素
* @param collection  集合
 */
func InCollection(value interface{}, collection interface{}) (inCollection bool, err error) {
	collectionValueOf := reflect.ValueOf(collection)
	collectionTypeOf := collectionValueOf.Type()
	collectionKind := collectionTypeOf.Kind()
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			return
		}
	}()
	switch collectionKind {
	case reflect.Array, reflect.Slice:
		for i := 0; i < collectionValueOf.Len(); i++ {
			itemValue := collectionValueOf.Index(i).Interface()
			if itemValue == value {
				inCollection = true
				return
			}
		}
	}
	inCollection = value == collection
	return
}

//加密
func Encryption(str string) (encryptionStr string, err error) {
	key := []byte("asd12345")
	if len(str) < 2 {
		fmt.Println(" 输入参数")
		return
	}
	plaintext := []byte(str)
	fmt.Println("明文为：", string(plaintext))
	// 传入明文和自己定义的密钥，密钥为8字节
	cryptText, err := goEncrypt.DesCbcEncrypt(plaintext, key, []byte{})
	if err != nil {
		return
	}
	encryptionStr = base64.StdEncoding.EncodeToString(cryptText)
	fmt.Println("DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错，8字节 如果解密秘钥错误解密后的明文会为空
	newplaintext, err := goEncrypt.DesCbcDecrypt(cryptText, key, []byte{}) //解密得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		return
	}
	fmt.Println("DES的CBC模式解密完:", string(newplaintext))
	return
}

func ToString(obj interface{}) string {
	if m, ok := obj.(string); ok {
		return m
	}
	if m, ok := obj.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	if m, ok := obj.(int); ok {
		return strconv.Itoa(m)
	}
	if m, ok := obj.(int64); ok {
		return strconv.FormatInt(m, 10)
	}
	if m, ok := obj.(bool); ok {
		return strconv.FormatBool(m)
	}

	return ""
}

func TimeToDate(timestamp int64, timeLayout string) string {
	if timeLayout == "" {
		timeLayout = FORMAT_DATE_CST
	}
	date := time.Unix(timestamp, 0).Format(timeLayout)
	return date
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func InUintArray(v uint, arr []uint) bool {
	flag := false
	for _, val := range arr {
		if val == v {
			flag = true
			break
		}
	}
	return flag
}

func InStringArray(v string, arr []string) bool {
	flag := false
	for _, val := range arr {
		if val == v {
			flag = true
			break
		}
	}
	return flag
}

//获取后一个小时时间
func GetNextHours(t string) (d string) {
	formatTime, _ := time.Parse(FORMAT_DATEHOUR_CST, t)
	h, _ := time.ParseDuration("1h")
	d = formatTime.Add(h).Format(FORMAT_DATEHOUR_CST)
	return
}

//获取上个月开始以及结束时间
func GetLastMonth() (start, end string) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start = thisMonth.AddDate(0, -1, 0).Format(FORMAT_DATE_CST)
	end = thisMonth.AddDate(0, 0, -1).Format(FORMAT_DATE_CST)
	return
}

//获取前一个小时时间
func GetLastHours(t string) (d string) {
	formatTime, _ := time.Parse(FORMAT_DATEHOUR_CST, t)
	h, _ := time.ParseDuration("-1h")
	d = formatTime.Add(h).Format(FORMAT_DATEHOUR_CST)
	return
}

//@brief:填充明文
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

//@brief:去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//@brief:AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
