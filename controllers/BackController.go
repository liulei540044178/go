package controllers

import (
	"github.com/astaxie/beego"
	"note/models"
	"fmt"
	"strconv"
	"strings"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseRouter struct {
	beego.Controller
	//i18n.Locale
	//user    models.User
	//isLogin bool
}

func (this *BaseRouter) Prepare(){
	//记录起始时间
	//this.Data["PageStartTime"] = time.now()

	// Setting properties.
	//this.Data["AppDescription"] = utils.AppDescription
	//this.Data["AppKeywords"] = utils.AppKeywords
	//this.Data["AppName"] = utils.AppName
	//this.Data["AppVer"] = utils.AppVer
	//this.Data["AppUrl"] = utils.AppUrl
	//this.Data["AppLogo"] = utils.AppLogo
	//this.Data["AvatarURL"] = utils.AvatarURL
	//this.Data["IsProMode"] = utils.IsProMode

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

type BackController struct {
	BaseRouter
}

func (this *BackController) NestPrepare() {

	////获取方式：
	////ctx.Input.Session("uid").(int)
	//a := this.GetSession("signup_msg")
	//
	//
	////this.Ctx.WriteString("我做处理了！！！！")
	//this.Ctx.WriteString(a.(models.User).Username)
	//this.Ctx.WriteString(this.Ctx.Input.Param("key"))


	//删除 session
	defer  this.DelSession("signup_msg")




	//可以做权限校验
	//if this.CheckActiveRedirect() {
	//	return
	//}
/*
	// if user isn't admin, then logout user
	if !this.user.IsAdmin {
		models.LogoutUser(&this.Controller)

		// write flash message
		this.FlashWrite("NotPermit", "true")

		this.Redirect("/login", 302)
		return
	}

	// current in admin page
	this.Data["IsAdmin"] = true

	if app, ok := this.AppController.(ModelPreparer); ok {
		app.ModelPrepare()
		return
	}
	*/
}


//   跳转后台主页面 拉取相应数据
func (this *BackController) Index(){
	this.TplName = "articel/index.tpl"
	articel := models.Articel{}
	notes := articel.ReadByCondition(1)
	this.Data["notes"] = notes
	this.Data["FromName"] = "user"
	//this.Layout = "admin/layout.html"

}

func (this *BackController) Query() {
	//执行函数之前需要判断是后登入,以及对文章的控制，是否来自登入者
	fromUser := this.Ctx.Input.Param("username")
	isnull := strings.TrimSpace(fromUser)
	//无效判断
	if isnull !=""{
		//判断用户是否已登入状态
		this.Ctx.WriteString("字符串为空")
		return
	}
	//获取url中配置：id
	idString := this.Ctx.Input.Param(":id")
	//this.Ctx.Input.Bind(&id, "id")  //?id ==123 绑定?后面的参数
	articel := models.Articel{}
	//字符串转换成int
	id,_ := strconv.Atoi(idString)
	articel.Id =id
	articel.QueryOne()

	//需要去字符串中的回车换行和空格
	//否则json解析的时候带了回车的话 eval 函数解析不了 会报 Uncaught SyntaxError: Unexpected token ILLEGAL
	//格式化json字符串:
	json:=fmt.Sprintf("{id:\"%d\",title:\"%s\",content:\"%s\"}",articel.Id,articel.ArticelTitle,articel.ArticelContent)
	str := strings.Replace(json, " ", "", -1)
	// 去除换行符
	str = strings.Replace(json, "\n", "", -1)
	this.Data["json"] = str
	//返回json数据
	this.ServeJSON()
}


//func (this *BaseAdminRouter) CheckActiveRedirect(){
//
//}


//func (this *BaseAdminRouter) LogoutUser(c Controller){
//
//}
//
//func (this *BaseAdminRouter) FlashWrite(msg string , ok bool){
//
//}