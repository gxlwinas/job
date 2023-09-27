package common

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func Setemail(emailto string) (bool, string) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(9000) + 1000
	str := strconv.Itoa(randomInt)
	html1 := "<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n<title>Document</title>\n</head>\n\n<body>\n<td class=\"p-80 mpy-35 mpx-15\" bgcolor=\"#212429\" style=\"padding: 80px; \">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n\n<tbody>\n\n<tr style=\"background-color: #eeeee8 ;\">\n<td>\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"title-36 pb-30 c-grey6 fw-b\" style=\"font-size:36px; line-height:42px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 30px; color:#bfbfbf; font-weight:bold;\"><span style=\"color: #77b9ee;\"></span></td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"text-18 c-grey4 pb-30\" style=\"font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:gray; padding-bottom: 30px;\">看起来您正在尝试注册趣丸兼职网站账号。此处是您访问帐户所需的验证码：</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pb-70 mpb-50\" style=\"padding-bottom: 70px;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" bgcolor=\"#17191c\">\n<tbody>\n<tr>\n<td class=\"py-30 px-56\" style=\"padding-top: 30px; padding-bottom: 30px; padding-left: 56px; padding-right: 56px;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td style=\"font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#8f98a0; text-align:center;\">\n请求来自 </td>\n</tr>\n<tr>\n<td style=\"font-size:25px; line-height:30px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#f1f1f1; text-align:center;letter-spacing:1px\">\nxxx兼职网站 </td>\n</tr>\n<tr>\n<td style=\"padding-bottom: 16px\"></td>\n</tr>\n<tr>\n<td class=\"title-48 c-blue1 fw-b a-center\" style=\"font-size:48px; line-height:52px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#3a9aed; font-weight:bold; text-align:center;\">\n"
	html2 := "</td>\n\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pb-30\" style=\"padding-bottom: 30px;\">\n<table width=\"210\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td><br>&nbsp;</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"title-36 pb-30 c-grey6 fw-b\" style=\"font-size:30px; line-height:34px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 20px; color:#000000; font-weight:bold;\">不是您？</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"text-18 c-grey4 pb-30\" style=\"font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:gray; padding-bottom: 30px;\">您会收到这封电子邮件，是由于有人试图登录您的xxx帐户，且提供了<span style=\"color: gray; font-weight: bold;\">正确的邮箱</span>。<br><br> 如果这不是您尝试注册，建议您忽略此信息 。\n<br><br> 此电子邮件包含一个登录代码，您需要用它访问您的帐户。切勿与任何人分享此代码。\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pb-30\" style=\"padding-bottom: 30px;\">\n<table width=\"210\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td><br>&nbsp;</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n\n</table>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n</tbody>\n</table>\n\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"pt-30\" style=\"padding-top: 30px;\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"img\" width=\"3\" bgcolor=\"#3a9aed\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"></td>\n<td class=\"img\" width=\"37\" style=\"font-size:0pt; line-height:0pt; text-align:left;\"></td>\n<td>\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tbody>\n<tr>\n<td class=\"text-16 py-20 c-grey4 fallback-font\" style=\"font-size:16px; line-height:22px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-top: 20px; padding-bottom: 20px; color:gray;\">\n祝您愉快 </td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n</td>\n</tr>\n</tbody>\n</table>\n\n\n</td>\n</tr>\n\n</tbody>\n</table>\n</td>\n</body>\n\n</html>"
	htmlContent := html1 + str + html2
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
		return false, str
	} else {
		return true, str
	}
}
