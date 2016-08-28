package word

import (
    "fmt"
    "errors"
    "archive/zip"
    "encoding/xml"
)

// ReadZipToFile - чтение Zip в файл
func ReadZipToFile(zipReader *zip.ReadCloser, wordFile *File) error {
    if zipReader == nil {
        return errors.New("Zip Reader Invalidate (nil)")
    }
    if wordFile == nil {
        return errors.New("Loaded Word File Invalidate (nil)")
    }
    defer zipReader.Close()
    // Чтение    
    for _, v := range zipReader.File {
        var err error
        if v != nil {
            fmt.Println("File in zip: ", v.Name)
            switch v.Name {
                case "[Content_Types].xml": {
                    wordFile.ContentTypes = new(ContentTypes)
                    err = readXMLFromZipFile(v, wordFile.ContentTypes)
                }                
                case "docProps/core.xml": {
                    wordFile.CoreProperties = new(CoreProperties)
                    err = readXMLFromZipFile(v, wordFile.CoreProperties)
                }
                case "docProps/app.xml": {
                    wordFile.AppProperties = new(AppProperties)
                    err = readXMLFromZipFile(v, wordFile.AppProperties)
                }
                case "word/fontTable.xml": {
                    wordFile.FontTable = new(FontTable)
                    err = readXMLFromZipFile(v, wordFile.FontTable)
                }
                case "word/settings.xml": {
                    wordFile.Settings = new(Settings)
                    err = readXMLFromZipFile(v, wordFile.Settings)
                }
                case "word/styles.xml": {
                    wordFile.Styles = new(Styles)
                    err = readXMLFromZipFile(v, wordFile.Styles)
                }
            }
        }
        if err != nil {
            return err
        }
    }    
    return nil
}

// readXMLFromZipFile 
func readXMLFromZipFile(zipFile *zip.File, out interface{}) error {
    rc, err := zipFile.Open()
    if err != nil {
        return err
    }
    decoder := xml.NewDecoder(rc)
    err = decoder.Decode(out)
    if err != nil {
        return err
    }
    return nil
}