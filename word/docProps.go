package word

import (
    "time"
    "encoding/xml"
)

// CoreProperties - Document properties
type CoreProperties struct {            
    XMLName            xml.Name    `xml:"coreProperties"`
    XMLNsCP            string      `xml:"xmlns:cp,attr"`
    XMLNsDC            string      `xml:"xmlns:dc,attr"`
    XMLNsDCT           string      `xml:"xmlns:dcterms,attr"`
    XMLNsDCM           string      `xml:"xmlns:dcmitype,attr,omitempty"`    
    XMLNsXSI           string      `xml:"xmlns:xsi,attr"`    
    Creator            string      `xml:"dc:creator,omitempty"`
    LastModifiedBy     string      `xml:"cp:lastModifiedBy,omitempty"`
    Revision           int         `xml:"cp:revision,omitempty"`
    Created           *time.Time   `xml:"dcterms:created,omitempty"`
    Modified          *time.Time   `xml:"dcterms:modified,omitempty"`
}

// AppProperties - Extended properties
type AppProperties struct {
    XMLName              xml.Name    `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Properties"`
    XMLNsVT              string      `xml:"vt,attr"`
    TotalTime            int64       `xml:",omitempty"`
    Pages                int64       `xml:",omitempty"`
    Words                int64       `xml:",omitempty"`
    Characters           int64       `xml:",omitempty"`
    Application          string      `xml:",omitempty"`
    DocSecurity          int64       `xml:",omitempty"`
    Lines                int64       `xml:",omitempty"`
    Paragraphs           int64       `xml:",omitempty"`
    ScaleCrop            bool        `xml:",omitempty"`
    Company              string      `xml:",omitempty"`
    LinksUpToDate        bool        `xml:",omitempty"`
    CharactersWithSpaces int64       `xml:",omitempty"`
    SharedDoc            bool        `xml:",omitempty"`
    HyperlinksChanged    bool        `xml:",omitempty"`
    AppVersion           float64     `xml:",omitempty"`

    // Реализовать:
    // HeadingPairs
    // TitlesOfParts
}

// MakeDefaultCoreProperties - create default core properties
func MakeDefaultCoreProperties() *CoreProperties {
    props := new(CoreProperties)    
    props.XMLNsDC  = "http://purl.org/dc/elements/1.1/"
    props.XMLNsDCT = "http://purl.org/dc/terms/"    
    props.XMLNsXSI = "http://www.w3.org/2001/XMLSchema-instance"
    //props.XMLNsDCM = "http://purl.org/dc/dcmitype/"
    return props
}

// MakeDefaultAppProperties - create default app properties
func MakeDefaultAppProperties() *AppProperties {
    props := new(AppProperties)
    props.XMLNsVT = "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"
    return props;
}