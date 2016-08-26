package word

import (
    "archive/zip"
)

// File - Файл Word
type File struct {
    ContentTypes *ContentTypes
}

// Open (File) - открыть файл
func (f *File) Open(fileName string) error {    
	z, err := zip.OpenReader(fileName)
	if err != nil {
		return err
	}
    err = ReadZipToFile(z, f)    
    if err == nil {
        // Коррекции после удачной загрузки
        f.correctionTest()
    }
    return err
}

// correctionTest - коррекционный тест до сохранения и после загрузки
func (f *File) correctionTest() {
    // Станадартные типы по умолчанию
    if f.ContentTypes == nil {
        f.ContentTypes = MakeDefaultContentTypes()
    }
}
