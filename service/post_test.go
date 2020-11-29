package service

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestComment(t *testing.T) {
	/*pi:=PostInfo{
		Type:        1,
		RealName:    "fist",
		PhoneNum:    0,
		StuNum:      0,
		Where:       "四教",
		Time:        time.Now().Unix(),
		WhatKind:    "书",
		Description: "教室找到的",
	}*/
	pi := Tag{
		Where:    "all",
		Time:     time.Now().Unix(),
		WhatKind: "手机",
	}
	bytes, _ := json.Marshal(pi)
	fmt.Println(string(bytes))
}
