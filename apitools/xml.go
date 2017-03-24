package apitools

import (
	"encoding/xml"
	"fmt"
	"os"
)


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