package word

import (
    "encoding/xml"
)

// ContentTypes - [ContentTypes].xml
type ContentTypes struct {
    XMLName     xml.Name        `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
    Defaults    []Default       `xml:"Override"`
    Overrides   []Override      `xml:"Default"`
}

// Override - CT Override
type Override struct {
	PartName    string `xml:",attr"`
	ContentType string `xml:",attr"`
}

// Default - CT Default
type Default struct {
	Extension   string `xml:",attr"`
	ContentType string `xml:",attr"`
}

// MakeDefaultContentTypes - create default content types
func MakeDefaultContentTypes() *ContentTypes {
    types := new(ContentTypes)
    types.Overrides = make([]Override, 9)
	types.Defaults = make([]Default, 10)

    // Defaults
    types.Defaults[0].Extension     = "xml"
    types.Defaults[0].ContentType   = "application/xml"
    types.Defaults[1].Extension     = "rels"
    types.Defaults[1].ContentType   = "application/vnd.openxmlformats-package.relationships+xml"
    types.Defaults[2].Extension     = "jpeg"
    types.Defaults[2].ContentType   = "image/jpg"
    types.Defaults[3].Extension     = "png"
    types.Defaults[3].ContentType   = "image/png"
    types.Defaults[4].Extension     = "bmp"
    types.Defaults[4].ContentType   = "image/bmp"
    types.Defaults[5].Extension     = "gif"
    types.Defaults[5].ContentType   = "image/gif"
    types.Defaults[6].Extension     = "tif"
    types.Defaults[6].ContentType   = "image/tif"
    types.Defaults[7].Extension     = "pdf"
    types.Defaults[7].ContentType   = "application/pdf"
    types.Defaults[8].Extension     = "mov"
    types.Defaults[8].ContentType   = "application/movie"
    types.Defaults[9].Extension     = "vml"
    types.Defaults[9].ContentType   = "application/vnd.openxmlformats-officedocument.vmlDrawing"
    types.Defaults[9].Extension     = "xlsx"
    types.Defaults[9].ContentType   = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

    // Overrides
    types.Overrides[0].PartName     = "/docProps/core.xml"
    types.Overrides[0].ContentType  = "application/vnd.openxmlformats-package.core-properties+xml"
    types.Overrides[1].PartName     = "/docProps/app.xml"
    types.Overrides[1].ContentType  = "application/vnd.openxmlformats-officedocument.extended-properties+xml"
    types.Overrides[2].PartName     = "/word/document.xml"
    types.Overrides[2].ContentType  = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"
    types.Overrides[3].PartName     = "/word/settings.xml"
    types.Overrides[3].ContentType  = "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"
    types.Overrides[4].PartName     = "/word/fontTable.xml"
    types.Overrides[4].ContentType  = "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"
    types.Overrides[5].PartName     = "/word/styles.xml"
    types.Overrides[5].ContentType  = "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"
    types.Overrides[6].PartName     = "/word/header1.xml"
    types.Overrides[6].ContentType  = "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml"
    types.Overrides[7].PartName     = "/word/footer1.xml"
    types.Overrides[7].ContentType  = "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml"
    types.Overrides[8].PartName     = "/word/theme/theme1.xml"
    types.Overrides[8].ContentType  = "application/vnd.openxmlformats-officedocument.theme+xml"
    
    return types;
}