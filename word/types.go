package word

// <item val="val"/>

// XMLValueInfo - информация о значении
type XMLValueInfo struct {
    Name     string `xml:"name,attr,omitempty"`
    URI      string `xml:"uri,attr,omitempty"`
    Language string `xml:"lang,attr,omitempty"`    
}

// XMLValue - одиночное значение
type XMLValue struct {
    XMLValueInfo
    Value    string `xml:"val,attr"`
}

// XMLBoolValue - одиночное int значение
type XMLBoolValue struct {
    XMLValueInfo
    Value    bool `xml:"val,attr"`
}

// XMLIntValue - одиночное int значение
type XMLIntValue struct {
    XMLValueInfo
    Value    int64 `xml:"val,attr"`
}

// XMLFloatValue - одиночное float значение
type XMLFloatValue struct {
    XMLValueInfo
    Value    float64 `xml:"val,attr"`
}