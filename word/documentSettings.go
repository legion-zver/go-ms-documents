package word

import (    
    "encoding/xml"
)

// Settings - параметры документа /word/settings.xml
type Settings struct {
    XMLName                       xml.Name          `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main settings"`
    XMLNsMC                       string            `xml:"mc,attr"`
    XMLNsO                        string            `xml:"o,attr"`
    XMLNsV                        string            `xml:"v,attr"`
    View                          XMLValue          `xml:"view"`
    MirrorMargins                 XMLValue          `xml:"mirrorMargins"`
    BordersDoNotSurroundHeader    XMLValue          `xml:"bordersDoNotSurroundHeader"`
    BordersDoNotSurroundFooter    XMLValue          `xml:"bordersDoNotSurroundFooter"`
    RevisionView                  RevisionView      `xml:"revisionView"`
    DefaultTabStop                XMLValue          `xml:"defaultTabStop"`
    AutoHyphenation               XMLValue          `xml:"autoHyphenation"`
    EvenAndOddHeaders             XMLValue          `xml:"evenAndOddHeaders"`
    BookFoldPrinting              XMLValue          `xml:"bookFoldPrinting"`
    NoLineBreaksAfter             XMLValue          `xml:"noLineBreaksAfter"`
    NoLineBreaksBefore            XMLValue          `xml:"noLineBreaksBefore"`
    Compat                        Compat            `xml:"compat"`
    ClrSchemeMapping              ClrSchemeMapping  `xml:"clrSchemeMapping"`
}

// RevisionView in settings
type RevisionView struct {
    Markup      int `xml:"markup,attr"`
    Comments    int `xml:"comments,attr"`
    InsDel      int `xml:"insDel,attr"`
    Formatting  int `xml:"formatting,attr"`
}

// ClrSchemeMapping in settings
type ClrSchemeMapping struct {
    Bg1     string `xml:"bg1,attr"`
    T1      string `xml:"t1,attr"`    
    Bg2     string `xml:"bg2,attr"`
    T2      string `xml:"t2,attr"`

    Accent1 string `xml:"accent1,attr"`
    Accent2 string `xml:"accent2,attr"`
    Accent3 string `xml:"accent3,attr"`
    Accent4 string `xml:"accent4,attr"`
    Accent5 string `xml:"accent5,attr"`
    Accent6 string `xml:"accent6,attr"`

    Hyperlink           string `xml:"hyperlink,attr"`
    FollowedHyperlink   string `xml:"followedHyperlink,attr"`
}

// Compat in settings
type Compat struct {
    CompatSetting *XMLValue `xml:"compatSetting,omitempty"`
}