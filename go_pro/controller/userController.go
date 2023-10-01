package controller

import (
	"github.com/gin-gonic/gin"
	"go_pro/common"
	"go_pro/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
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

	db.Where("Email=?", email).First(&user)
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

	is, captcha := common.Setemail(email)

	if !is {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "发送验证码失败",
		})
		return
	}

	newuser := model.Newuser{
		Username: username,
		Password: string(hasedPassword),
		Email:    email,
		Captcha:  captcha,
	}
	newuser.Deleteat = time.Now().Add(time.Minute * 20)

	db.Create(&newuser)

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"captacha": captcha,
		"message":  "请输入验证码",
	})
}

func Captcha(c *gin.Context) {
	db := common.GetDB()
	db.Where("Deleteat < ?", time.Now()).Delete(&model.Newuser{})
	email := c.PostForm("email")
	captcha := c.PostForm("captcha")

	var newuser model.Newuser
	db.Where("Email = ?", email).First(&newuser)
	if captcha != newuser.Captcha {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "验证码错误",
		})
		return
	}
	user := model.User{
		Username: newuser.Username,
		Password: newuser.Password,
		Email:    newuser.Email,
	}

	res := db.Create(&user)
	if res != nil {
		errorMessage := res.Error.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()

	email := c.PostForm("email")
	password := c.PostForm("password")

	var user model.User
	db.Where("Email=?", email).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱不存在",
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
		log.Printf("token generate error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "系统异常",
		})

		return
	}
	c.Set("token", token)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "登录成功",
	})

}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{"user": user},
	})
}

func JobPublic(c *gin.Context) {
	db := common.GetDB()
	var job model.Job
	c.Bind(&job)
	if err := db.Create(&job).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "兼职已发布"})
}

func PrinvatePublic(c *gin.Context) {
	db := common.GetDB()
	usernemail, _ := c.Params.Get("useremail")
	var jobList []model.Job
	if err := db.Where("Useremail = ?", usernemail).Find(&jobList).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error})
		return
	}
	c.JSON(http.StatusOK, jobList)

}

func JobInfo(c *gin.Context) {
	db := common.GetDB()
	var jobList []model.Job
	address := c.Query("address")
	wages := c.Query("wages")
	jobtype := c.Query("jobtype")
	if err := db.Where("Address = ?", address).Where("Wages >= ?", wages).Where("Jobtype = ?", jobtype).Find(&jobList).Error; err != nil {
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
	var job model.Job
	if err := db.Where("ID=?", id).First(&job).Error; err != nil {
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
	if err := db.Where("id=?", id).Delete(model.Job{}).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "工作已删除"})
}

func Link(c *gin.Context) {
	db := common.GetDB()
	jobid, ok := c.Params.Get("jobid")
	id, _ := strconv.Atoi(jobid)
	applyemail := c.PostForm("applicant")
	joblink := model.Joblink{
		Jobid:      id,
		Applyemail: applyemail,
	}
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "未获取到job"})
		return
	}
	db.Create(&joblink)
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "申请成功"})

}

func Privatelink(c *gin.Context) {
	db := common.GetDB()
	useremail, err := c.Params.Get("useremail")
	if !err {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "未获取到账户信息"})
		return
	}
	var joblink []model.Joblink
	db.Where("Applyemail = ?", useremail).Find(&joblink)
	c.JSON(http.StatusOK, joblink)
}
