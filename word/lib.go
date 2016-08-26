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
                    err = readContentTypesToFile(v, wordFile)
                }                
                case "docProps/core.xml": {
                    err = readCorePropertiesToFile(v, wordFile)
                }
                case "docProps/app.xml": {
                    err = readAppPropertiesToFile(v, wordFile)
                }
            }
        }
        if err != nil {
            return err
        }
    }    
    return nil
}

// Чтение контент типов в документе
func readContentTypesToFile(zipFile *zip.File, wordFile *File)  error {
    wordFile.ContentTypes = new(ContentTypes)
    rc, err := zipFile.Open()
    if err != nil {
        return err
    }
    decoder := xml.NewDecoder(rc)
    err = decoder.Decode(wordFile.ContentTypes)
    if err != nil {
        return err
    }
    return nil
}

// Чтение параметров Core
func readCorePropertiesToFile(zipFile *zip.File, wordFile *File)  error {
    wordFile.CoreProperties = new(CoreProperties)
    rc, err := zipFile.Open()
    if err != nil {
        return err
    }
    decoder := xml.NewDecoder(rc)
    err = decoder.Decode(wordFile.CoreProperties)
    if err != nil {
        return err
    }
    return nil
}

// Чтение параметров App
func readAppPropertiesToFile(zipFile *zip.File, wordFile *File)  error {
    wordFile.AppProperties = new(AppProperties)
    rc, err := zipFile.Open()
    if err != nil {
        return err
    }
    decoder := xml.NewDecoder(rc)
    err = decoder.Decode(wordFile.AppProperties)
    if err != nil {
        return err
    }
    return nil
}