package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	 //内嵌一个beego控制器 相当于继承
	//其中包括 Init、Prepare、Post、Get、Delete、Head 等方法。
	// 我们可以通过重写的方式来实现这些方法，而我们上面的代码就是重写了 Get 方法。
	beego.Controller
}

//func (c *MainController) Get() {
//	//设置传递参数
//	c.Data["Website"] = "beego.me"
//	//模板解析字符串{{.Email}}
//	c.Data["Email"] = "astaxie@gmail.com"
//	//设置模板
//	//this.TplName 就是需要渲染的模板，这里指定了 index.tpl，如果用户不设置该参数，
//	// 那么默认会去到模板目录的 Controller/<方法名>.tpl 查找
//	c.TplName = "index.tpl"
//
//
//}

func (c *MainController) Get(){
	//a := beego.AppConfig.String("mysqluser")
	//c.Ctx.WriteString(a)
	c.TplName = "login/login.tpl"


	//o := orm.NewOrm()
	//
	//user := models.User{Username: "slene" , Password:"12345" }

	// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	//user.Username = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	//u := models.User{Id: user.Id}
	//err := o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)

	// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//c.Ctx.WriteString("jehhe")
}


