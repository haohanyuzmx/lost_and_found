package service

import (
	"fmt"
	"gorm.io/gorm"
	"lost_and_found/model"
)

type Post interface {
	Post()
	Detail()
	Name() string
	Id() uint
}
type PostInfo struct {
	Type        int    `json:"type"`
	RealName    string `json:"real_name"`
	PhoneNum    int64  `json:"phone_num"`
	StuNum      int64  `json:"stu_num"`
	Where       string `json:"where"`
	Time        int64  `json:"time"`
	WhatKind    string `json:"what_kind"`
	Description string `json:"description"`
	UserID      uint
}
type PostList struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	UserID      uint
	Portrait    string `json:"portrait"`
	Picture     string `json:"picture"`
}
type Tag struct {
	Where    string `json:"where"`
	Time     int64  `json:"time"`
	WhatKind string `json:"what_kind"`
}

func GetPostFromPI(info PostInfo) Post {
	if info.Type == 1 {
		return &model.LostAndFound{
			RealName:    info.RealName,
			PhoneNum:    info.PhoneNum,
			StuNum:      info.StuNum,
			Where:       info.Where,
			Time:        info.Time,
			WhatKind:    info.Where,
			Description: info.Description,
			UserID:      info.UserID,
		}
	} else {
		return &model.NoticeForLost{
			Description: info.Description,
			UserID:      info.UserID,
		}
	}

}

func NewPost(kind int, id uint) Post {
	if kind == 1 {
		return &model.LostAndFound{
			Model: gorm.Model{ID: id},
		}
	} else {
		return &model.NoticeForLost{
			Model: gorm.Model{ID: id},
		}
	}
}

func Talk(pi PostInfo) uint {
	t := GetPostFromPI(pi)
	t.Post()
	return t.Id()
}

func Comment(uid, pid uint, kind int, comment string) {
	model.PushComment(uid, pid, kind, comment)
}

func ShowPostDetail(kind int, id uint) Post {
	post := NewPost(kind, id)
	post.Detail()
	return post
}

func FindPosts(start, end, kind int) []PostList {
	var ps []PostList
	if kind == 1 {
		lafs := model.CheckLOFByDesc(start, end)
		for _, laf := range lafs {
			u := model.User{Model: gorm.Model{ID: laf.UserID}}
			u.GetPortrait()
			var picture string
			if len(laf.File) < 1 {
				picture = ""
			} else {
				picture = laf.File[0].Uri
			}
			pl := PostList{
				Id:          laf.ID,
				Description: laf.Description,
				UserID:      laf.UserID,
				Portrait:    u.Portrait.Uri,
				Picture:     picture,
			}
			ps = append(ps, pl)
		}
	} else {
		nfls := model.CheckLOFByDesc(start, end)
		for _, nfl := range nfls {
			u := model.User{Model: gorm.Model{ID: nfl.UserID}}
			u.GetPortrait()
			pl := PostList{
				Id:          nfl.ID,
				Description: nfl.Description,
				UserID:      nfl.UserID,
				Portrait:    u.Portrait.Uri,
				Picture:     nfl.File[1].Uri,
			}
			ps = append(ps, pl)
		}
	}
	return ps
}

func SelectPost(t Tag) []model.LostAndFound {
	laf := model.LostAndFound{
		WhatKind: t.WhatKind,
		Where:    `%` + t.Where + `%`,
		Time:     t.Time,
	}
	fmt.Println(t)
	if t.Where == "all" {
		laf.Where = "%%"
	}
	if t.WhatKind == "all" {
		laf.WhatKind = "%%"
	}
	fmt.Println(laf)
	lafs := laf.GetByTag()
	return lafs
}
