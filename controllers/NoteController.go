package controllers

import(
	"github.com/astaxie/beego"
	"note/models"
	"fmt"
)

type NoteController struct {
	beego.Controller
}



func (c  *NoteController) URLMapping() {
	//c.Mapping("Add",c.Add)
	c.Mapping("Add", c.Add)

}

//进入首页面

// @router /note/addnote/:key [get post head put delete options *]
func (this *NoteController) Add() {
	//需要加上网关 判断是否来自本站的请求


	//第一次访问 将 用户 与 input[type=hidden] 记录到 session中 ,做绑定
	//判断 提交hidden域 是否与 session 中的一致 一样则不做保存

	//因为此处是Ajax请求 ， 我知道 绑定formId ==> 模型名 所以需要使用原生代码与模型对接
	//
	note :=models.Articel{}
	note.ArticelContent = this.Input().Get("noteContent")
	note.ArticelTitle = this.Input().Get("noteTitle")
	msg :=note.BindState()
	var mystruct string
	if msg =="succeed"{
		note.InsertOne()
		mystruct ="{state:'"+ msg +"',operation:'插入成功^.^'}"
	}else{
		mystruct ="{state:'"+ msg +"',operation:'插入失败^.^'}"
	}


	this.Data["json"] = &mystruct
	//this.Ctx.WriteString("===")
	this.ServeJSON()
}





//测试代码
func (this *NoteController) GetList() {
	//var notes [] models.Articel
	articel := models.Articel{}
	notes := articel.ReadByCondition(1)
	for i:=0; i<len(notes) ;i++{
		fmt.Println("user nums: ", notes[i].ArticelContent)
	}
	this.Ctx.WriteString("我的哥")
}

type Obj interface {
}

//type user struct{
//	Id bson.ObjectId `bson:"_id"`
//	Name string `bson:"name"`
//	Age string `bson:"age"`
//}


func (this *NoteController) Test(){
	//
	//m := models.MongoDB{"user"}
	//result := new(user)
	////m.FindOne(bson.M{"name":"中12尉1212"},&result)
	//
	//state := m.FindOneByID("58cdf2b85711801cd062de63",&result)
	////state:=m.Remove(bson.M{"name":"中尉"})
	//u := user{}
	//u.Id = bson.NewObjectId()
	//u.Name ="嚒嚒大"
	//u.Age = "12"
	////state := m.InsertOne(u)
	//
	//this.Ctx.WriteString("w 我      "+state+" : \r\n")
	//fmt.Print(result.Id)
	//this.Ctx.WriteString("w 我       :"+a)

}

func (this *NoteController) Test1()  {
	//session, err := mgo.Dial("127.0.0.1:27017")
	//if err !=nil {
	//	fmt.Print(err)
	//}
	//c := session.DB("lagou").C("user")
	//result :=  new(user)
	////err = c.Find(bson.M{"name": "aaa12"}).One(&result)
	////58cda847a96559da0e4db145
	//err = c.Find(bson.M{"name" : "aaa12"}).One(&result)
	////
	////err = c.Insert(bson.M{"name":"2222"})
	//
	////err = c.Insert(&result)
	//
	////获取所有的doc
	////var users []user
	////err = c.Find(nil).All(&users)
	////for _,person :=range users{
	////	fmt.Print(person.Name+"=================="+person.Age+"\r\n")
	////}
	//
	////更新
	////u := user{Name:"updata",Age:"1313"}
	////为找到会 err = not found
	//err = c.Update(bson.M{"name": "aaa"}, bson.M{"name": "updata","age":"1313"})
	//
	//
	//fmt.Print(err)
	//this.Ctx.WriteString("w 我       age:"+result.Age +"nimeifu!!!")
}