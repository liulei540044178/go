package models



type Articel struct {
	Id int `form:-`
	ArticelTitle string `form:"noteTitle"`
	ArticelContent string `form:"noteContent"`
	IsDelete int8 `form:-` //标识文章是否被删除
	ParentId int //文章所属目录
}


func (this *Articel) validation() bool{
	if this.ArticelTitle == ""|| this.ArticelContent =="" {
		return false
	}else{
		return true
	}

}

func (this *Articel) BindState() string{
	state := this.validation()
	if state {
		return "succeed"
	}else{
		return "faild"
	}

}

// 解耦合   ===》　变得比较繁琐了
func (m *Articel) InsertOne() {
	tool.InsertOne(m)
}


func (m *Articel) UpdateOne() {
	tool.UpdateById(m)
}


func (m *Articel) Del() {
	 tool.DelById(m)

}

func (m *Articel) QueryOne() {
	tool.ReadById(m)
}

func (m *Articel) GetList() string{

	return ""
}

func (m *Articel) TableName() string{

	return "articel"
}


//返回本类数组
func (m *Articel) ReadByCondition( parentId int) []Articel {

	var notes [] Articel
	tool.ReadByCondition(m , &notes ,parentId)
	return  notes
}

func (m *Articel) Condition(c ... string) string  {
	return ""
}