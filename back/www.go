package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	mail "github.com/xhit/go-simple-mail/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type Useroi struct {
	ID       int
	Email    string
	Username string
	Password string
	Captcha  string
	Deleteat time.Time
}

type Information struct {
	ID          int
	UserEmail   string
	Topic       string
	Type        string
	Pay         string
	Times       string
	Address     string
	FullAddress string
	Description string
}

type Userinfor struct {
	Useremail string
	Toid      string
	Name      string
	Photo     string
	Age       string
	Gender    string
}

func main() {
	//gin
	r := gin.Default()
	r.Use(Cors())

	//gorm
	dsn := "admin:gexiangliangLl7_@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Useroi{})
	db.AutoMigrate(&Information{})
	db.AutoMigrate(&Userinfor{})
	//登录
	r.POST("/login", func(c *gin.Context) {
		email := c.Query("username")
		password := c.Query("password")
		var user User
		db.Where("Email = ?", email).First(&user)
		if db.Error == nil {
			if user.Password == password {
				c.String(200, "ok")
			}
		} else {
			c.String(200, "fail")
		}

	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "sdsd")

	})

	//注册
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		emailto := c.PostForm("email")
		var newuser Useroi
		newuser.Email = emailto
		newuser.Password = password
		newuser.Username = username
		var user User
		result := db.Where("Email = ?", emailto).First(&user)
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			rand.Seed(time.Now().UnixNano())
			randomInt := rand.Intn(9000) + 1000
			str := strconv.Itoa(randomInt)
			html1 := "<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n<title>Document</title>\n</head>\n\n<body>\n<td class=\"p-80 mpy-35 mpx-15\" bgcolor=\"#212429\" style=\"padding: 80px; \">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n\n<tbody>\n\n<tr style=\"background-color: #eeeee8 ;\">\n<td>\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"title-36 pb-30 c-grey6 fw-b\" style=\"font-size:36px; line-height:42px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 30px; color:#bfbfbf; font-weight:bold;\"><span style=\"color: #77b9ee;\"></span></td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"text-18 c-grey4 pb-30\" style=\"font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:gray; padding-bottom: 30px;\">看起来您正在尝试注册趣丸兼职网站账号。此处是您访问帐户所需的验证码：</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pb-70 mpb-50\" style=\"padding-bottom: 70px;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" bgcolor=\"#17191c\">\n<tbody>\n<tr>\n<td class=\"py-30 px-56\" style=\"padding-top: 30px; padding-bottom: 30px; padding-left: 56px; padding-right: 56px;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td style=\"font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#8f98a0; text-align:center;\">\n请求来自 </td>\n</tr>\n<tr>\n<td style=\"font-size:25px; line-height:30px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#f1f1f1; text-align:center;letter-spacing:1px\">\nxxx兼职网站 </td>\n</tr>\n<tr>\n<td style=\"padding-bottom: 16px\"></td>\n</tr>\n<tr>\n<td class=\"title-48 c-blue1 fw-b a-center\" style=\"font-size:48px; line-height:52px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#3a9aed; font-weight:bold; text-align:center;\">\n"
			html2 := "</td>\n\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pb-30\" style=\"padding-bottom: 30px;\">\n<table width=\"210\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td><br>&nbsp;</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"title-36 pb-30 c-grey6 fw-b\" style=\"font-size:30px; line-height:34px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 20px; color:#000000; font-weight:bold;\">不是您？</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"text-18 c-grey4 pb-30\" style=\"font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:gray; padding-bottom: 30px;\">您会收到这封电子邮件，是由于有人试图登录您的xxx帐户，且提供了<span style=\"color: gray; font-weight: bold;\">正确的邮箱</span>。<br><br> 如果这不是您尝试注册，建议您忽略此信息 。\n<br><br> 此电子邮件包含一个登录代码，您需要用它访问您的帐户。切勿与任何人分享此代码。\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pb-30\" style=\"padding-bottom: 30px;\">\n<table width=\"210\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td><br>&nbsp;</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n</tbody>\n</table>\n\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pt-30\" style=\"padding-top: 30px;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"img\" width=\"3\" bgcolor=\"#3a9aed\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"></td>\n<td class=\"img\" width=\"37\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"></td>\n<td>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"text-16 py-20 c-grey4 fallback-font\" style=\"font-size:16px; line-height:22px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-top: 20px; padding-bottom: 20px; color:gray;\">\n祝您愉快 </td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n\n\n</td>\n</tr>\n\n</tbody>\n</table>\n</td>\n</body>\n\n</html>"
			htmlContent := html1 + str + html2
			newuser.Captcha = str
			newuser.Deleteat = time.Now().Add(time.Hour)
			result = db.Create(&newuser)
			email := mail.NewMSG()
			email.SetFrom("2869842198@qq.com").
				AddTo(emailto).
				SetSubject("测试邮件").
				SetBody(mail.TextHTML, htmlContent)
			//SetBody(mail.TextPlain, str)
			client := mail.NewSMTPClient()
			client.Host = "smtp.qq.com"
			client.Port = 587
			client.Username = "2869842198@qq.com"
			client.Password = "tshgigkclupiddda"
			server, err := client.Connect()
			defer server.Close()
			err = email.Send(server)
			if err != nil {
				c.String(500, "出错了")
				log.Println(err)
			} else {
				c.String(200, "success")
			}
		} else {
			c.String(200, "该邮箱已被使用")
		}
	})

	//验证码检测
	r.POST("/captcha", func(c *gin.Context) {
		cap := c.Query("captcha")
		email := c.Query("email")
		db.Where("deleteat < ?", time.Now()).Delete(&Useroi{})
		var user User
		re := db.Where("email = ?", email).First(&user)
		if re.Error != nil && re.Error == gorm.ErrRecordNotFound {
			var useroi Useroi
			db.Where("email = ?", email).First(&useroi)
			if cap == useroi.Captcha {
				var user User
				user.Username = useroi.Username
				user.Password = useroi.Password
				user.Email = useroi.Email
				result := db.Create(&user)
				if result.Error != nil {
					c.String(200, "出错了")
					//log.Fatal(err)
				} else {
					c.String(200, "success")
				}
			} else {
				c.String(200, "验证码错误")
			}
		} else {
			c.String(200, "账号已存在")
		}

	})

	//创建 兼职信息
	r.POST("/publish", func(c *gin.Context) {
		var information Information
		information.UserEmail = c.Query("useremail")
		information.Topic = c.PostForm("topic")
		information.Type = c.PostForm("type")
		information.Pay = c.PostForm("pay")
		information.Times = c.PostForm("time")
		information.Address = c.PostForm("address")
		information.FullAddress = c.PostForm("fulladdress")
		information.Description = c.PostForm("description")
		result := db.Create(&information)
		if result != nil {
			c.String(500, "错误")
		} else {
			c.String(200, "创建成功")
		}
	})

	//筛选合适信息
	r.GET("/home", func(c *gin.Context) {
		address := c.Query("address")
		types := c.Query("type")
		pay := c.Query("pay")
		infor := []Information{}
		db.Where("address = ?", address).Where("type = ?", types).Where("pay >= ?", pay).Find(&infor)
		jsonData, err := json.Marshal(infor)
		if err != nil {
			c.String(500, "JSON serialization failed")
		} else {
			c.JSON(200, jsonData)
		}
	})

	//获取兼职详细信息
	r.GET("/apply", func(c *gin.Context) {
		var information Information
		id := c.Query("id")
		ID, err := strconv.Atoi(id)
		if err != nil {
			c.String(500, "服务器出错了")
		} else {
			db.Where("ID = ?", ID).First(&information)
			jsonData, err := json.Marshal(information)
			if err != nil {
				c.String(500, "服务器出错了")
			} else {
				c.JSON(200, jsonData)
			}
		}
	})

	//进行申请
	r.POST("/apply", func(c *gin.Context) {
		var userinfor Userinfor
		userinfor.Toid = c.Query("id")
		userinfor.Useremail = c.Query("useremail")
		userinfor.Name = c.PostForm("name")
		userinfor.Photo = c.PostForm("photo")
		userinfor.Age = c.PostForm("age")
		userinfor.Gender = c.PostForm("gender")
		re := db.Create(&userinfor)
		if re != nil {
			c.String(500, "出错了")
		} else {
			c.String(200, "申请成功")
		}
	})
	//获得我的发布
	//r.GET("/person", func(c *gin.Context) {
	//	module := c.Query("module")
	//	useremail := c.Query("useremail")
	//
	//})

	r.Run(":8080")
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
