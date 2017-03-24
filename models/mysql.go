package models


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"

	//"web/models"
)
// go 1.8
//驱动包
//MySQL：github.com/go-sql-driver/mysql
//对象关系映射
//go get github.com/astaxie/beego/orm
//解耦合
//使用方式：
//	//在模型init函数中注册以下
//		m :=Mysql{}
//		user := new(模型)
//		m.RegisterModel(user)
//	实现以下接口
//		type Parent interface {
//			UpdateOne()
//			InsertOne()
//			Del()
//			QueryOne()
//		}
//	找到RegisterModel() 方法 注册模型
//orm.RegisterModel(new(Articel) , new(UserInfo))

type Parent interface {
	//子类自定义条件查询
	Condition(c ... string) string
	//
	UpdateOne()
	InsertOne()
	Del()
	QueryOne()
	TableName() string
	//ReadByCondition() []interface{}
}

var(
	linkString string
)

type Mysql struct {
}

func init()  {
	dbUser := "root"
	dbName := "yii2basic" //数据库名字
	dbPwd := "123456"
	linkString = fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName)
	//最后一个参数是设置连接池数量
	//自己设置连接数量
	//orm.SetMaxOpenConns("default",30)
	m :=Mysql{}
	m.RegisterModel()
}



//注册模型
func (m *Mysql) RegisterModel() {
	// param1 means table's alias name. default is "default".
	// param2 means run next sql if the current is error.
	// param3 means show all info when running command or not.
	//orm.RegisterDataBase("default", "mysql", dbLink)
	orm.RegisterDataBase("default", "mysql", linkString )
	//需要在这里注册
	orm.RegisterModel(new(Articel) , new(UserInfo) , new(Directory))
	//orm.RegisterModel( new(UserInfo))
	orm.RunSyncdb("default", false, true)
	//m.O = m.getORM()
}

func (m *Mysql) getORM() orm.Ormer {
	return orm.NewOrm()
}
// 调用函数 xx(&xx)
// 因为模型层传递的为指针变量，所以不需要在取&
// 定义函数 xx(xx interface{})
func (m *Mysql) InsertOne(ref interface{}) string {
	_, err := m.getORM().Insert(ref)
	if err == nil{
		//return fmt.Sprintf("%s",err)
		return "seccessd"
	}else{
		//return fmt.Sprintf("%s",err)
		return fmt.Sprintf("%s",err)
	}

}

func (m *Mysql) UpdateById(ref interface{}) string {
	_, err := m.getORM().Update(ref)
	if err == nil{
		//return fmt.Sprintf("%s",err)
		return "seccessd"
	}else{
		//return fmt.Sprintf("%s",err)
		return fmt.Sprintf("%s",err)
	}

}


func (m *Mysql) DelById(ref interface{}) int64 {
	num, err := m.getORM().Delete(ref)
	if err == nil{
		//return fmt.Sprintf("%s",err)
		return num
	}else{
		//return fmt.Sprintf("%s",err)
		return 0
	}

}

func (m *Mysql) ReadById(ref interface{}) {
	 m.getORM().Read(ref)

}

func (m *Mysql) QueryAll(ref Parent , objs interface{}){
	o :=  m.getORM()
	tableName := ref.TableName()
	sql := fmt.Sprintf("SELECT * from %s",tableName)
	// 此处不需要穿点第二参数
	num, err := o.Raw(sql).QueryRows(objs)
	if err == nil {
		//fmt.Sprintf("%d" , len(objs))
	}else{
		fmt.Println("%s %s:",sql , num)
	}
}

func (m *Mysql) ReadByCondition(ref Parent,objs interface{},parentId int)  {
	o :=  m.getORM()
	tableName := ref.TableName()
	//var notes [] Articel
	sql := fmt.Sprintf("SELECT * from %s where parent_id = ?",tableName)
	_, err := o.Raw(sql, 1).QueryRows(objs)
	if err == nil {
		//for i:=0; i<len(notes) ;i++{
		//	fmt.Println("user nums: ", notes[i].ArticelContent)
		//}
	}else{
		fmt.Sprintf("%s " , err)
	}
	// 获取 QuerySeter 对象，user 为表名
	//qs := o.QueryTable("user")
	// 获取 QuerySeter 对象，user 为表名

	//qs.Filter("id", 1) // WHERE id = 1
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("Profile__Age", 18) // 使用字段名和 Field 名都是允许的
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18
	//qs.Filter("profile__age__gte", 18) // WHERE profile.age >= 18
	//qs.Filter("profile__age__in", 18, 20) // WHERE profile.age IN (18, 20)
	//
	//qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000)



}


func (m *Mysql) QueryByCondition(ref Parent,objs interface{} ,c string)  {
	o :=  m.getORM()
	tableName := ref.TableName()
	//var notes [] Articel
	sql := fmt.Sprintf("SELECT * from %s where %s ",tableName,c)
	fmt.Println(sql)
	_, err := o.Raw(sql).QueryRows(objs)
	if err == nil {
		//fmt.Println("===",len(objs))
	}else{
		fmt.Sprintf("%s " , err)
	}

}


