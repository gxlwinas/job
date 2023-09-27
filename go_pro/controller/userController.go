package controller

import (
	"go_pro/common"
	"go_pro/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := common.GetDB()

	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	if !common.VerifyEmailFormat(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "无效的电子邮箱",
		})
		return
	}

	var user model.User

	db.Where("email=?", email).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "该邮箱已被注册",
		})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}

	newUser := model.User{
		Username: username,
		Password: string(hasedPassword),
		Email:    email,
	}

	db.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "请输入验证码",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()

	var requestUser model.User
	c.Bind(&requestUser)
	username := requestUser.Username
	password := requestUser.Password

	if len(username) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "请输入用户名",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	var user model.User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不存在",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "系统异常",
		})
		log.Printf("token generate error: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "登录成功",
	})

	c.Set("token", token)

}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{"user": user},
	})
}

func JobPublic(c *gin.Context) {
	db := common.GetDB()
	var job model.JobJson
	c.Bind(&job)
	if err := db.Create(&job).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "兼职已发布到数据库"})
}

func JobInfo(c *gin.Context) {
	db := common.GetDB()
	var jobList []model.JobJson
	if err := db.Find(&jobList).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobList)
}

func JobUpdata(c *gin.Context) {
	db := common.GetDB()
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "未获取到id"})
		return
	}
	var job model.JobJson
	if err := db.Where("id=?", id).First(&job).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error()})
		return
	}
	c.Bind(&job)
	if err := db.Save(&job).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "工作详情已修改", "job": job})

}

func JobDelete(c *gin.Context) {
	db := common.GetDB()
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "未获取到id"})
		return
	}
	if err := db.Where("id=?", id).Delete(model.JobJson{}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "工作已删除"})
}
