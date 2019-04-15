package models

//声明一个Token用的列表
type Auth struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//鉴别函数
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.Id > 0 {
		return true
	}
	return false
}
