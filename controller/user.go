package controller

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/HAL-Future-Creation-Exhibition/go-nellow/model"
	"github.com/HAL-Future-Creation-Exhibition/go-nellow/service"
	"github.com/gin-gonic/gin"
)

var UserController userController

type userController struct{}

func (u *userController) Get(c *gin.Context) {
	users, err := service.User.FindAll()
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "err user findAll",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *userController) Create(c *gin.Context) {
	var req struct {
		IsNellow bool `json:"is_nellow"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(req.IsNellow)
	var user model.User
	user.ID = bson.NewObjectId()
	user.Name = "bcp-guest-00"
	user.Icon = "hogehoge.iconUrl.com"
	err = service.User.Create(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "err user create",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
