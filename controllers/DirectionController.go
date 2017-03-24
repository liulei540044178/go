package controllers
import (
	"github.com/astaxie/beego"
	"note/models"
	"encoding/xml"
	"fmt"
	"os"
	//"note/apitools"
	//"bee/vendor/github.com/lib/pq/oid"
)

type DirectionController struct{
	beego.Controller
}
func CreateXML(){
	type Directory struct {
		XMLName   xml.Name `xml:"Directory"`
		Id        int      `xml:"id,attr"`
		Item_id	  int      `xml:"Item_id,attr"`
		Revel  	  int      `xml:"Revel,attr"`
		Name      string    `xml:"Directory>Name"`

	}
	v := &Directory{Id: 1, Revel: 1, Item_id: 1, Name: "第一个例子"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)


}

func sortMenus(menus [] models.Directory , parentId int ) string {
	 //arr :=menus
	//slice := make([]models.Directory, 0)
	var objS  string = ""
	for i:=0; i<len(menus);i++{

		if menus[i].Dir_revel != 1 && menus[i].Item_id == parentId{
			//slice = append(slice , menus[i])
			obj := fmt.Sprintf("{'id':%d ,'dir_name':%s}," ,
				menus[i].Id , menus[i].Dir_name )
			objS = objS + obj
			//fmt.Println("       2j====%s" , objS)
		}else{
			//sortMenus(menus , parentId)
		}
		//fmt.Println(1)
	}
	return  objS + "}]"
}

func (m *DirectionController) Test(){
	dir := models.Directory{}
	menus := dir.QueryAll()
	var objs []string
	//var menuList []string
	//slice := make([]models.Directory, 0)
	//const size = len(menus)
	//var slice []interface{}
	//var slice1 []type = make([]type, len)
	for i:=0; i<len(menus);i++{

		if menus[i].Dir_revel == 1 {
			obj := fmt.Sprintf("[{menubar:{'id':%d ,'dir_name':%s},items:[" ,
				menus[i].Id , menus[i].Dir_name )
			objs = append(objs,obj)
			objs = append(objs,sortMenus(menus,menus[i].Id))

			//sortMenus(menus,menus[i].Id)
			//slice[i] = child

		}
	}

	var json string
	for i:=0; i<len(objs);i++{
		json = json + objs[i]
	}
	//fmt.Println(" 1j ==== %s" , objs)

	//fmt.Println("目录层级 ==== ：",len(slice[3]))
	//for i:= 0; i<len(slice); i++{
	//	for j:=0; j<len(slice[i]); j++{
	//		menu:=slice[i][j]
	//		fmt.Println("一级目录====:",menu)
	//	}
	//
	//}

/*
//fmt.Println(len(menus))
size := len(menus)
var menuList[][size] models.Directory
for i:=0; i<len(menus); i++{
	menuList[i] = menus[i]
	condition := fmt.Sprintf("item_id = %d", menus[i].Id)
	items  := dir.GetMenu(condition)
	for j:=0; j<len(items); j++{
		menuList[i]["item"][j] =  items[j]
		fmt.Println(items[j].Dir_name)
	}
}
*/
//m.Ctx.WriteString("====")
//CreateXML()
	m.Data["json"] = json
	m.ServeJSON()
}