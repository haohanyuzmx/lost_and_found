package model

import "gorm.io/gorm"

type File struct {
	gorm.Model `json:"-"`
	Uri        string
	RelationID uint `json:"-"`
	Category   int  `json:"-"` //0为用户，1为lostAndfound，2为
}

func (f *File) Creat() error {
	err := DB.Table("files").Create(&f).Error
	if err != nil {
		return err
	}
	return nil
}
func (f *File) Delete() error {
	if err := DB.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}
func GetFile(relationid uint, category int) (f []File) {
	DB.Where(&File{RelationID: relationid, Category: category}).Find(&f)
	return
}
