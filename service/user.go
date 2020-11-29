package service

import "lost_and_found/model"

func CreatUser(username, password string) error {
	u := model.User{
		Username: username,
		Password: password,
	}
	return u.Create()
}
func CheckUser(username, password string) (model.User, error) {
	u := model.User{
		Username: username,
		Password: password,
	}
	u, err := u.CheckByUP()
	return u, err
}
func GetUser(id uint) model.User {
	var u model.User
	u.ID = id
	u.GetUser()
	u.GetPortrait()
	return u
}
func SaveUser(user *model.User) {
	_ = user.Update()
	user.GetUser()
}
func FindMyThing(u model.User) map[uint]string {
	info := model.CheckByEffectiveInfo(u)
	data := make(map[uint]string)
	for _, found := range info {
		data[found.ID] = found.Description
	}
	return data
}
