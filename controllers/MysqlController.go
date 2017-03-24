package controllers


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"note/models"
)

//驱动包
//MySQL：github.com/go-sql-driver/mysql
//对象关系映射
//go get github.com/astaxie/beego/orm
type MysqlController struct {
	beego.Controller
}

//
type User struct {
	Id int
	Name string
}


func (c *MysqlController) Add(){
	//o := orm.NewOrm()
	//m := models.Mysql{}
	user := models.UserInfo{Username: "newslene",Password:"123455",Age:12,IsAdmin:true}
	//m.InsertOne(&user)
	user.InsertOne()
	//user
	// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	c.Ctx.WriteString("excute successd")
}

func (c *MysqlController) Update(){
	user := models.UserInfo{Id:1,Username: "newslene1212",Password:"1234512125",Age:12,IsAdmin:true}
	//m.InsertOne(&user)
	user.UpdateOne()
	c.Ctx.WriteString("excute successd")
}

func (c *MysqlController) Del(){
	user := models.UserInfo{Id:2,Username: "newslene1212",Password:"1234512125",Age:12,IsAdmin:true}
	//m.InsertOne(&user)
	user.Del()
	c.Ctx.WriteString("excute successd")
}


func (c *MysqlController) Read(){
	user := models.UserInfo{}
	user.Id = 1
	user.QueryOne()
	///fmt.Println(o.Update(user))

	//fmt.Println(o.Delete(user))
	c.Ctx.WriteString(user.Username)
}

