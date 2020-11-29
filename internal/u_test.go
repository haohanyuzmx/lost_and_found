package internal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUpdateReal(t *testing.T) {
	r := RealBase{
		PhoneNum: 177,
		StuNum:   2019,
		RealName: "nihao",
		QQNum:    7352,
	}
	marshal, _ := json.Marshal(r)
	fmt.Println(string(marshal))
}
func TestRegister(t *testing.T) {
	u := UserBase{
		Username: "test1",
		Password: "test1",
	}
	marshal, _ := json.Marshal(u)
	fmt.Println(string(marshal))
}
func TestPushComment(t *testing.T) {
	//info := CommentInfo{
	//	Kind:    1,
	//	PostId:  1,
	//	Comment: "我的",
	//}
	lq := ListRequest{
		Start: 0,
		End:   10,
		Kind:  1,
	}
	marshal, _ := json.Marshal(lq)
	fmt.Println(string(marshal))
}
