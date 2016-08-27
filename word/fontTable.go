package word

import (    
    "encoding/xml"
)

// FontTable - таблица шрифтов
type FontTable struct {
    XMLName      xml.Name   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main fonts"`
    Fonts      []*Font      `xml:"font"`
}

// Font - шрифт
type Font struct {
    Name    string   `xml:"name,attr"`
    Charset XMLValue `xml:"charset"`
    Family  XMLValue `xml:"family"`
    Pitch   XMLValue `xml:"pitch"`
}