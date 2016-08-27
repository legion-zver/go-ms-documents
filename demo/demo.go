package main

import (
	"encoding/xml"
	"fmt"	
	"github.com/legion-zver/go-ms-documents/word"
)

func main() {		 				  	
	file := new(word.File)
	err := file.Open("example.docx")
	if err != nil {
		fmt.Println(err)
		return
	}
	if file.Settings != nil {
		data, _ := xml.Marshal(file.Settings)
		fmt.Println(string(data))
	}
}
