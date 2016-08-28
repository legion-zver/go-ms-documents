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
    Paragraph     ParagraphPr        `xml:"pPr"`
    // For text
    TextFormat    TextFormatPr       `xml:"rPr"`
    // Only for type = table
    TableStylePrList []*TableStylePr `xml:"tblStylePr,omitempty"`
}

// TableStylePr - item TblStylePrList
type TableStylePr struct {
    Type    string      `xml:"type,attr"`
}

// ParagraphPr in style (pPr)
type ParagraphPr struct {
    KeepNext                 *XMLIntValue    `xml:"keepNext,omitempty"`
    KeepLines                *XMLIntValue    `xml:"keepLines,omitempty"`
    PageBreakBefore          *XMLIntValue    `xml:"pageBreakBefore,omitempty"`
    WidowControl             *XMLIntValue    `xml:"widowControl,omitempty"`
    SuppressAutoHyphens      *XMLIntValue    `xml:"suppressAutoHyphens,omitempty"`
    Bidi                     *XMLIntValue    `xml:"bidi,omitempty"`
    Jc                       *XMLValue       `xml:"jc,omitempty"`
    OutlineLevel             *XMLIntValue    `xml:"outlineLvl,omitempty"`
    Shadow                   *ShadowValue    `xml:"shd,omitempty"`
    Spacing                  *SpacingValue   `xml:"spacing,omitempty"`
    Ind                      *IndValue       `xml:"ind,omitempty"`
}

// TextFormatPr in style (rPr)
type TextFormatPr struct {
    Fonts *TextFontStylePr `xml:"rFonts,omitempty"`
    B          *XMLIntValue     `xml:"b,omitempty"`
    BCS        *XMLIntValue     `xml:"bCs,omitempty"`
    I          *XMLIntValue     `xml:"i,omitempty"`
    ICS        *XMLIntValue     `xml:"iCs,omitempty"`
    Caps       *XMLIntValue     `xml:"caps,omitempty"`
    SmallCaps  *XMLIntValue     `xml:"smallCaps,omitempty"`
    Strike     *XMLIntValue     `xml:"strike,omitempty"`
    DStrike    *XMLIntValue     `xml:"dstrike,omitempty"`
    Outline    *XMLIntValue     `xml:"outline,omitempty"`
    Emboss     *XMLIntValue     `xml:"emboss,omitempty"`
    Imprint    *XMLIntValue     `xml:"imprint,omitempty"`
    Vanish     *XMLIntValue     `xml:"vanish,omitempty"`
    Color      *XMLValue        `xml:"color,omitempty"`
    Spacing    *XMLIntValue     `xml:"spacing,omitempty"`
    W          *XMLIntValue     `xml:"w,omitempty"`
    Kern       *XMLIntValue     `xml:"kern,omitempty"`
    Position   *XMLIntValue     `xml:"position,omitempty"`
    Sz         *XMLIntValue     `xml:"sz,omitempty"`
    SzCS       *XMLIntValue     `xml:"szCs,omitempty"`
    U          *XMLValue        `xml:"u,omitempty"`
    Br         *XMLValue        `xml:"br,omitempty"`
    VertAlign  *XMLValue        `xml:"vertAlign,omitempty"`
    Lang       *XMLValue        `xml:"lang,omitempty"`
}

// TextFontStylePr in TextFormatPr
type TextFontStylePr struct {
    ASCII    string `xml:"ascii,attr"`
    CS       string `xml:"cs,attr"`
    HANSI    string `xml:"hAnsi,attr"`
    EastAsia string `xml:"eastAsia,attr"`
}

// ShadowValue in styly
type ShadowValue struct {
    Value string `xml:"val,attr"`
    Color string `xml:"color,attr"`
    Fill  string `xml:"fill,attr"`
}

// SpacingValue in styly
type SpacingValue struct {
    Before    int64  `xml:"before,attr"`
    After     int64  `xml:"after,attr"`
    Line      int64  `xml:"line,attr"`
    LineRule  string `xml:"lineRule,attr"`
}

// IndValue in styly
type IndValue struct {
    Left        int64  `xml:"left,attr"`
    Right       int64  `xml:"right,attr"`
    FirstLine   int64  `xml:"firstLine,attr"`
}