package models
//
//import (
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//	//"fmt"
//)
//const (
//	URL string= "127.0.0.1:27017"
//	DB  string = "website"
//	//C  string = "article"
//)
//var (
//	mgoSession *mgo.Session
//
//)
////定义接口  作为传递参数使用
//
//
type MongoDB struct {
	ColName string
}
////mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
////session, err := mgo.Dial("127.0.0.1:27017")
////if err !=nil {
////fmt.Print(err)
////}
////c := session.DB("lagou").C("user")
//
////连接池
//func (m *MongoDB)getSession() *mgo.Session {
//	if mgoSession == nil {
//		var err error
//		mgoSession, err = mgo.Dial(URL)
//		if err != nil {
//			panic(err) //直接终止程序运行
//		}
//	}
//	//最大连接池默认为4096
//	return mgoSession.Clone()
//}
//
//
////func witchCollection(collection string, s func(*mgo.Collection) error) error {
////	session := getSession()
////	defer session.Close()
////	c := session.DB(dataBase).C(collection)
////	return s(c)
////}
//
//func  (m *MongoDB) SetCollectionName(colName  string)  {
//	//m.colName = colName
//}
//
//func (m *MongoDB) getCollection(s func(*mgo.Collection) error) error{
//	session := m.getSession()
//	//获取时 从连接池中得到一个对象， 关闭时 将对象返回给连接池 而不是真正意义上的断开资源
//	//
//	defer session.Close()
//	c := session.DB(DB).C(m.ColName)
//	return s(c)
//}
//
//
////err = c.Insert(bson.M{"name":"2222"})
////err = c.Insert(&result)
//func (m *MongoDB) InsertOne(obj interface{}) string {
//	query := func(c *mgo.Collection) error {
//		return  c.Insert(obj)
//	}
//	err := m.getCollection(query)
//	if err !=nil {
//		return "false"
//	}
//	return "success"
//}
//
//
//
////result :=  new(user)
////err = c.Find(bson.M{"name": "aaa12"}).One(&result)
//
////注意这里要传递引用 否则会发生 函数内赋值操作
//func (m *MongoDB) FindOne(obj interface{},result interface{}) interface{} {
//	query := func(c *mgo.Collection) error {
//		return  c.Find(obj).One(result)
//	}
//	//fmt.Print(returnObj.Name)
//
//	err := m.getCollection(query)
//	if err !=nil {
//		return "false"
//	}
//	return result
//}
////by id
////func GetPersonById(id string) *Person {
////	objid := bson.ObjectIdHex(id)
////	person := new(Person)
////	query := func(c *mgo.Collection) error {
////		return c.FindId(objid).One(&person)
////	}
////	witchCollection("person", query)
////	return person
////}
//func (m *MongoDB) FindOneByID(id string , result interface{}) string  {
//	//需要将stringId 58cdf2b85711801cd062de63 做哈希转换
//	objid := bson.ObjectIdHex(id)
//	query := func(c *mgo.Collection) error {
//		return  c.FindId(objid).One(result)
//	}
//	//fmt.Print(returnObj.Name)
//
//	err := m.getCollection(query)
//	if err !=nil {
//		return "false"
//	}
//	return "success"
//}
//
//
//
////更新  先删除数据 后添加数据
////u := user{Name:"updata",Age:"1313"}
////为找到会 err = not found
////err = c.Update(bson.M{"name": "aaa"}, bson.M{"name": "updata","age":"1313"})
//func (m *MongoDB) UpdateOne(obj interface{}, newObj interface{}) string{
//	query := func(c *mgo.Collection) error {
//		return  c.Update(obj,newObj)
//	}
//	//fmt.Print(returnObj.Name)
//
//	err := m.getCollection(query)
//	if err !=nil {
//		return "false"
//	}
//	return "success"
//}
//
//
////每次删除一条
////func (c *Collection) Remove(selector interface{}) error
//func (m *MongoDB) Remove(obj interface{}) string{
//	query := func(c *mgo.Collection) error {
//		return  c.Remove(obj)
//	}
//	//fmt.Print(returnObj.Name)
//
//	err := m.getCollection(query)
//	if err !=nil {
//		return "false"
//	}
//	return "success"
//}
