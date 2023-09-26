package user

import "time"

// 用户信息
type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

// 储存用户注册信息
type Useroi struct {
	ID       int
	Email    string
	Username string
	Password string
	Captcha  string
	Deleteat time.Time
}
