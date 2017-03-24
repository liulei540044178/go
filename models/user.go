package models

import(
	"fmt"
)

var(
	//包私有变量
	tool Mysql
)



// form中使用id绑定类名
//如果要忽略一个字段，有两种办法，一是：字段名小写开头，二是：form 标签的值设置为 -
type  UserInfo struct {
	Id int `form:-`
	//这里的类 只是绑定表单数据`orm:"size(100)"`
	//变量名小写 为私有
	// 绑定需要些绑定数据注释 `form xxxx`
	IsAdmin  bool `form:-`
	Username string `form:"username"  orm:"size(25)" valid:"Required;Match(/^Bee.*/)"`  // Name 不能为空并且以 Bee 开头
	Password string `form:"password"`
	//别忘了绑定form 属性
	Age int `form:"age" valid:"Min(1)"`
}


func (m *UserInfo) checkUser() bool{
	//if()
	return true
}


// 解耦合   ===》　变得比较繁琐了
func (m *UserInfo) InsertOne() {
	state := tool.InsertOne(m)
	fmt.Printf(state)
}


func (m *UserInfo) UpdateOne() {
	state := tool.UpdateById(m)
	fmt.Printf(state)
}


func (m *UserInfo) Del() {
	state := tool.DelById(m)
	fmt.Sprintf("asdf %d",state)
}

func (m *UserInfo) QueryOne() {
	tool.ReadById(m)
}


func (m *UserInfo) Condition(c ... string) string  {
	return "a"
}