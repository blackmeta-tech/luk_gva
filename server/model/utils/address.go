package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Address string //自定义

func (t *Address) UnmarshalJSON(data []byte) (err error) {
	str := ""
	if err = json.Unmarshal(data, &str); err != nil {
		return
	}

	if err = t.CheckAddress(str); err != nil {
		return err
	}
	// 转成小写
	*t = Address(strings.ToLower(str))
	return
}

func (t Address) MarshalJSON() (data []byte, err error) {

	if err = t.CheckAddress(string(t)); err != nil {
		return
	}
	data = make([]byte, 0, t.len()+2)
	data = append(data, '"')
	data = append(data, t...)
	data = append(data, '"')
	return
}

func (t *Address) len() int {
	return len(string(*t))
}

func (t *Address) String() string {
	return strings.ToLower(string(*t))
}

//  实现 Value 方法，写入数据库时会调用该方法  不能用指针
func (t Address) Value() (driver.Value, error) {
	str := strings.ToLower(string(t))
	return str, nil
}

// 实现 Scan 方法，读取数据库时会调用该方法
func (t *Address) Scan(v interface{}) error {
	switch v.(type) {
	case string:
		value := v.(string)
		*t = Address(strings.ToLower(value))
		return nil
	case int:
		value := v.(int)
		*t = Address(strings.ToLower(string(value)))
		return nil
	case []byte:
		value := string(v.([]byte))
		*t = Address(strings.ToLower(value))
		return nil
	default:
		return fmt.Errorf("can not convert %v to Address", v)
	}
}

func (t *Address) CheckAddress(addr string) error {
	if !strings.Contains(addr, "0x") {
		return errors.New(" address error")
	}
	// if len(addr) != 42 {
	// 	return errors.New(" address lenght error")
	// }

	return nil
}
