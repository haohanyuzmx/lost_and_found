package model

import (
	"fmt"
	"strconv"
	"testing"
)

func TestUser_Create(t *testing.T) {
	//f:=File{
	//	Uri:        "assets/1_1.png",
	//	RelationID: 1,
	//	Category:   0,
	//}
	//fmt.Println(f.Creat())
	u := User{
		Username: "first",
		Password: "first",
	}

	fmt.Println(u.CheckByUP())
}
func TestPushComment(t *testing.T) {
	for i := 0; i < 50; i++ {
		si := strconv.Itoa(i)
		laf := LostAndFound{
			UserID:      1,
			Description: si,
		}
		laf.Post()
	}
}
