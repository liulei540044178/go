package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	beego.Controller
}


func (c *LoginController) Index(){

	//c.Ctx.WriteString("===index")
	//不设置此参数 会自动在controllers目录下寻找模板文件
	c.TplName = "login/login.tpl"
	//c.Ctx.Output.Body([]byte("ok"))
	c.Render()

}
func (c *LoginController) Signup(){

	//c.Ctx.WriteString("===index")
	//不设置此参数 会自动在controllers目录下寻找模板文件
	c.TplName = "login/signup.tpl"
	//c.Ctx.Output.Body([]byte("ok"))
	c.Render()

}



func (this *LoginController) Post(){
	//也可以在该目录下创建一个类 使用下面方式初始化
	//u := User{}
	//访问models下面的User
	u := models.User{}
	if err := this.ParseForm(&u); err != nil {
		//自定义 异常处理
	}else{


		valid := validation.Validation{}
		b, err := valid.Valid(&u)
		if err != nil {
			// handle error
			this.Ctx.WriteString("爆你妹的错~~~~")
		}
		if !b {
			// validation does not pass
			// blabla...
			for _, err := range valid.Errors {
				//log.Println(err.Key, err.Message)
				this.Ctx.WriteString(err.Key+"  ===  "+err.Message)
			}
		}


		//redirect 之前有输出的话 那么将不会跳转
		//设置session信息
		v := this.GetSession("signup_msg")
		if v == nil {
			this.SetSession("signup_msg", u)
			//this.Data["msg"]= u.Username
		} else {
			msg:=this.GetSession("signup_msg")
			//如果ssesion存在 通过GetSession(xx) 获得一个接口对象v
			//v.(类型)  == 那么可以保存类对象
			this.Data["num"] = msg.(models.User)
		}


		this.Redirect("/back/index",302)

	}

}


//func (this *LoginController) GetUrl() {
//	this.Ctx.Output.Body([]byte(this.UrlFor(".Myext")))
//}


//	返回json数据
//	需要配置copyrequestbody = true
func (this *LoginController) Json() {
	objectid := "hehe"
	mystruct := "{\"ObjectId\":\"" + objectid + "\"}"
	this.Data["json"] = mystruct
	this.ServeJSON()
}
//	返回xml数据
func (this *LoginController) Xml() {
	mystruct := "<student><name>小明</name></student>"
	this.Data["xml"] = mystruct
	this.ServeXML()
}

//	返回xml数据
//	调用 ServeJSONP 之后，会设置 content-type 为 application/javascript，然后同时把数据进行 JSON 序列化，然后根据请求的 callback 参数设置 jsonp 输出。
func (this *LoginController) JavaScript() {
	mystruct := "<script>alert(1)</script>"
	this.Data["jsonp"] = mystruct
	this.ServeJSONP()
}
