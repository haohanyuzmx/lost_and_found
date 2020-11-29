package service

import (
	"fmt"
	"io"
	"lost_and_found/model"
	"lost_and_found/util"
	"mime/multipart"
	"os"
	"time"
)

func SaveFile(file *multipart.FileHeader, Category int, RelationID uint) {
	var dst string
	unix := time.Now().Unix()
	time := util.Int64PtrString(unix)
	switch Category {
	case 0:
		dst = "assets/user/" + time + file.Filename
	case 1:
		dst = "assets/lost_and_found/" + time + file.Filename
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		fmt.Println(err)
		return
	}
	f := model.File{
		Uri:        dst,
		RelationID: RelationID,
		Category:   Category,
	}
	if f.Creat() != nil {
		fmt.Println(err)
		return
	}
}
