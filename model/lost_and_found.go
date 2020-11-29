package model

import (
	"gorm.io/gorm"
)

type LostAndFound struct {
	gorm.Model
	RealName    string //拥有具体信息（以下三项
	PhoneNum    int64
	StuNum      int64
	Where       string //索引详情（以下三项
	Time        int64
	WhatKind    string
	Description string
	UserID      uint      //发帖人
	File        []File    `gorm:"foreignKey:RelationID"`
	Comment     []Comment `gorm:"foreignKey:PostID"`
}

func (laf *LostAndFound) Post() {
	DB.Create(laf).First(laf)

}

func (laf *LostAndFound) Detail() {
	DB.Model(laf).First(laf)
	fs := GetFile(laf.ID, 1)
	laf.File = fs
	comments := GetCommentByPostID(laf.ID, 1)
	laf.Comment = comments
}

func (laf *LostAndFound) Id() uint {
	return laf.ID
}

func (laf *LostAndFound) Name() string {
	return "LostAndFound"
}

func (laf *LostAndFound) GetByTag() (lafs []LostAndFound) {
	DB.Model(laf).Exec("SELECT * FROM `lost_and_founds` WHERE"+
		" (time<? and what_kind like ? and `where` like ?) "+
		"AND `lost_and_founds`.`deleted_at` IS NULL", laf.Time, laf.WhatKind, laf.Where).Find(&lafs)
	for _, found := range lafs {
		found.Detail()
	}
	return
}

func CheckByEffectiveInfo(user User) []LostAndFound {
	rl := CheckByRealName(user.RealName)
	sn := CheckByStuNum(user.StuNum)
	pn := CheckByPhoneNum(user.PhoneNum)
	laf := append(rl, append(sn, pn...)...)
	return laf
}
func CheckByRealName(realname string) (lafs []LostAndFound) {
	DB.Model(&LostAndFound{}).Where(&LostAndFound{
		RealName: realname,
	}).Find(&lafs)
	return lafs
}
func CheckByStuNum(stunum int64) (laf []LostAndFound) {
	DB.Model(&LostAndFound{}).Where(&LostAndFound{
		StuNum: stunum,
	}).Find(&laf)
	return laf
}
func CheckByPhoneNum(phonenum int64) (lafs []LostAndFound) {
	DB.Model(&LostAndFound{}).Where(&LostAndFound{
		PhoneNum: phonenum,
	}).Find(&lafs)
	return lafs
}
func CheckLOFByDesc(start, end int) (lafs []LostAndFound) {
	DB.Model(&LostAndFound{}).Order("created_at desc").Offset(start).Limit(end - start).Find(&lafs)
	for _, laf := range lafs {
		laf.Detail()
	}
	return
}
