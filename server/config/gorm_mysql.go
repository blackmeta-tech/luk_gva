package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// func (m *Mysql) Dsn() string {
// 	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
// }


func (m *Mysql) Dsn() string {
	if m.Password == "" {
		// 没有密码时，不拼接冒号和密码
		return m.Username + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
