package word

import (
    "encoding/xml"
)

// Styles - word/styles.xml
type Styles struct {
    // Header
    XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main styles"`
    XMLNsMC     string `xml:"mc,attr"`
    XMLNsW14    string `xml:"w14,attr"`
    Ignorable   string `xml:",attr"`

    // Styles
    Styles    []*Style `xml:"style"`
}

// Style in styles
type Style struct {
    Type    string      `xml:"type,attr"`
    Default string      `xml:"default,attr"`
    StyleID string      `xml:"styleId,attr"`
    Name    XMLValue    `xml:"name"`
    Next    XMLValue    `xml:"next"`

    // For paragraph 
    ParagraphPr ParagraphPr          `xml:"pPr"`
    // Only for type = table
    TableStylePrList []*TableStylePr `xml:"tblStylePr,omitempty"`
}

// TableStylePr - item TblStylePrList
type TableStylePr struct {
    Type    string      `xml:"type,attr"`
}

// ParagraphPr in style
type ParagraphPr struct {
    KeepNext                 *XMLIntValue    `xml:"keepNext,omitempty"`
    KeepLines                *XMLIntValue    `xml:"keepLines,omitempty"`
    PageBreakBefore          *XMLIntValue    `xml:"pageBreakBefore,omitempty"`
    WidowControl             *XMLIntValue    `xml:"widowControl,omitempty"`
    SuppressAutoHyphens      *XMLIntValue    `xml:"suppressAutoHyphens,omitempty"`
    Bidi                     *XMLIntValue    `xml:"bidi,omitempty"`
    Jc                       *XMLIntValue    `xml:"jc,omitempty"`
}