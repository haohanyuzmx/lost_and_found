package model

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model `json:"-"`
	UserID     uint
	PostID     uint `json:"-"`
	Category   int  `json:"-"` //1为lostAndfound，2为
	Comment    string
}

func PushComment(uid, pid uint, category int, comment string) {
	DB.Create(&Comment{UserID: uid, PostID: pid, Category: category, Comment: comment})
}

func GetCommentByPostID(postid uint, category int) []Comment {
	var cs []Comment
	err := DB.Where(&Comment{PostID: postid, Category: category}).Find(&cs).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return cs
}
