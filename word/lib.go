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