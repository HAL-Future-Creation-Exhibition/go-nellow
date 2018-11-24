package service

import (
	"github.com/HAL-Future-Creation-Exhibition/go-nellow/model"
	mgo "gopkg.in/mgo.v2"
)

var User = userService{db.C("users")}

type userService struct {
	C *mgo.Collection
}

func (u *userService) FindAll() (model.Users, error) {
	var users model.Users
	err := u.C.Find(nil).All(&users)
	return users, err
}

func (u *userService) Create(user model.User) error {
	err := u.C.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}
