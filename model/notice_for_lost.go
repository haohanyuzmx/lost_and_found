package model

import "gorm.io/gorm"

type NoticeForLost struct {
	gorm.Model

	Description string
	UserID      uint
	File        []File    `gorm:"foreignKey:RelationID"`
	Comment     []Comment `gorm:"foreignKey:PostID"`
}

func (nfl *NoticeForLost) Post() {
	DB.Create(nfl).First(nfl)
}

func (nfl *NoticeForLost) Detail() {
	DB.Model(nfl).First(nfl)
	fs := GetFile(nfl.ID, 1)
	nfl.File = fs
	comments := GetCommentByPostID(nfl.ID, 1)
	nfl.Comment = comments
}

func (nfl *NoticeForLost) Id() uint {
	return nfl.ID
}

func (nfl *NoticeForLost) Name() string {
	return "NoticeForLost"
}

func CheckNFLByDesc(start, end int) (nfls []NoticeForLost) {
	DB.Model(&NoticeForLost{}).Order("created_at desc").Offset(start).Limit(end - start).Find(&nfls)
	for _, nfl := range nfls {
		nfl.Detail()
	}
	return
}
