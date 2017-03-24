package routers

import (
	"note/controllers"
	"github.com/astaxie/beego"

)

func init() {


//注解路由
//beego.Include(&controllers.NoteController{})
//beego.Include(&controllers.LogoutController{})


//路由适配器 定义url规则
beego.Router("/", &controllers.MainController{})
//对应一个LoginController
beego.Router("/login/index", &controllers.LoginController{},"GET:Index")
beego.Router("/login/sub", &controllers.LoginController{},"POST:Post")
beego.Router("/login/signup", &controllers.LoginController{},"GET:Signup")

beego.Router("/json", &controllers.LoginController{},"GET:Json")
beego.Router("/xml", &controllers.LoginController{},"GET:Xml")
beego.Router("/javascript", &controllers.LoginController{},"GET:JavaScript")


//后台控制器
//	登入首页面
beego.Router("/backed/index", &controllers.BackController{},"GET:Index")
//beego.Router(“/api/:id([0-9]+)“, &controllers.RController{})
//自定义正则匹配 //匹配 /api/123 :id = 123
//对note进行操作
beego.Router("/note/query/:username/?:id", &controllers.BackController{},"GET:Query")

//目录表的操作
beego.Router("/dir/test", &controllers.DirectionController{},"GET:Test")




	//beego.Router("/bakcrouter",&controllers.)
//开启自动匹配 url = login/index  ==> LoginController Index 方法 注意大小写
//beego.AutoRouter(&controllers.LoginController{})

//beego.Router("/noteop/addnote",&controllers.)
//beego.Router("/note/index", &controllers.NoteController{} , "GET:Index")

beego.Router("/note/addnote1", &controllers.NoteController{},"GET:Add")
beego.Router("/note/addnote", &controllers.NoteController{} , "POST:Add")
//	注意大小写  小写方法不能配置
beego.Router("/note/mog", &controllers.NoteController{} , "GET:Test")
beego.Router("/note/list", &controllers.NoteController{} , "GET:GetList")


beego.Router("/mysql/add", &controllers.MysqlController{} , "GET:Add")
beego.Router("/mysql/u", &controllers.MysqlController{} , "GET:Update")
beego.Router("/mysql/d", &controllers.MysqlController{} , "GET:Del")
beego.Router("/mysql/r", &controllers.MysqlController{} , "GET:Read")


}