package main

import (
    "fmt"
    "github.com/legion-zver/go-ms-documents/word"
)

func main()  {
    file := new(word.File)
    err := file.Open("example.docx")
    if err != nil {
        fmt.Println(err)
        return
    }
    if file.ContentTypes != nil {
        fmt.Println(*file.ContentTypes)
    }
}