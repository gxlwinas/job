package main

// import (
// 	"net/http"
// 	"regexp"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jinzhu/gorm"
// 	"golang.org/x/crypto/bcrypt"
// )

// type jobtype uint8
// type duration uint8

// const (
// 	Short duration = iota
// 	Medium
// 	Long
// )

// const (
// 	Physical jobtype = iota
// 	Mental
// )

// type User struct {
// 	gorm.Model
// 	Name            string `gorm:"varchar(20);not null"`
// 	Username        string
// 	Password        string `gorm:"size:255;not null"`
// 	Email           *string
// 	isAdministrator bool `gorm:"default:false"`
// 	isPublisher     bool `gorm:"default:false"`
// 	Brithday        *time.Time
// }

// type PartTimeJob struct {
// 	gorm.Model
// 	PublisherID uint
// 	Jobtype     jobtype
// 	Duration    duration
// 	Wages       string
// 	isEffective bool
// }

// func InitDB() (db *gorm.DB) {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
// 	db, err := gorm.Open("mysql", dsn)
// 	if err != nil {
// 		panic("failed to connect database, err: " + err.Error())
// 	}
// 	db.AutoMigrate(&User{})
// 	return db
// }

// func VerifyEmailFormat(email string) bool {
// 	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
// 	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

// 	reg := regexp.MustCompile(pattern)
// 	return reg.MatchString(email)
// }

// func main() {
// 	db := InitDB()
// 	defer db.Close()

// 	ginServer := gin.Default()
// 	ginServer.GET("/get", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello,it's get!")
// 	})
// 	ginServer.POST("/post", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello,it's post!")
// 	})
// 	//注册
// 	ginServer.POST("/register", func(c *gin.Context) {
// 		username := c.PostForm("username")
// 		password := c.PostForm("password")
// 		email := c.PostForm("email")

// 		if len(username) == 0 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "用户名不能为空",
// 			})
// 			return
// 		}

// 		if len(password) < 6 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "密码不能少于6位",
// 			})
// 			return
// 		}

// 		if !VerifyEmailFormat(email) {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "邮箱格式不正确",
// 			})
// 			return
// 		}

// 		var user User
// 		db.Where("username=?", username).First(&user)
// 		if user.ID != 0 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "用户名已存在",
// 			})
// 			return
// 		}
// 		db.Where("email=?", email).First(&user)
// 		if user.ID != 0 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "该邮箱已被注册",
// 			})
// 			return
// 		}

// 		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 		if err != nil {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    500,
// 				"message": "密码加密错误",
// 			})
// 			return
// 		}

// 		newUser := User{
// 			Username: username,
// 			Password: string(hasedPassword),
// 			Email:    &email,
// 		}

// 		db.Create(&newUser)

// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    200,
// 			"message": "注册成功",
// 		})
// 	})
// 	//登录
// 	ginServer.POST("/login", func(c *gin.Context) {
// 		username := c.PostForm("username")
// 		password := c.PostForm("password")

// 		if len(username) == 0 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "请输入用户名",
// 			})
// 			return
// 		}

// 		if len(password) < 6 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "密码不能少于6位",
// 			})
// 			return
// 		}

// 		var user User
// 		db.Where("username=?", username).First(&user)
// 		if user.ID == 0 {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "用户名不存在",
// 			})
// 			return
// 		}
// 		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 			c.JSON(http.StatusUnprocessableEntity, gin.H{
// 				"code":    422,
// 				"message": "密码错误",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    200,
// 			"message": "登录成功",
// 		})
// 	})
// 	ginServer.Run(":8080")
// }
