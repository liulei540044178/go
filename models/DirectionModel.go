package models
import(
	//"fmt"
)
type Directory struct {
	Id int
	Dir_name string
	Item_id int
	Dir_revel int
}


// 解耦合   ===》　变得比较繁琐了
func (m *Directory) InsertOne() {
	tool.InsertOne(m)
}


func (m *Directory) UpdateOne() {
	tool.UpdateById(m)
}


func (m *Directory) Del() {
	tool.DelById(m)

}

func (m *Directory) QueryOne() {
	tool.ReadById(m)
}

func (m *Directory) GetList() string{

	return ""
}

func (m *Directory) TableName() string{
	return "directory"
}

func (m *Directory) QueryAll() []Directory {
	var dirs [] Directory
	tool.QueryAll(m ,&dirs)
	return  dirs
}

//返回本类数组
func (m *Directory) ReadByCondition( parentId int) []Directory {

	var notes [] Directory
	tool.ReadByCondition(m , &notes ,parentId)
	return  notes
}

func (m *Directory) GetMenu(c string) []Directory {
	var menus [] Directory
	tool.QueryByCondition(m,&menus,c)
	return  menus
}

func (m *Directory) GetMenuItem() {

}

func (m *Directory) Condition(c ... string) string  {
	var query string
	if c[0] == ""{
		//return"1"
	}else{
		query = c[0]

	}
	return query
}