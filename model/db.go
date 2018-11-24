package model

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/HAL-Future-Creation-Exhibition/go-nellow/config"
	_ "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"
)

var db = NewDBConn()

func NewDBConn() *mgo.Database {
	fmt.Println(config.DBConfig)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    strings.Split(config.DBConfig.Host, ","),
		Timeout:  60 * time.Second,
		Username: config.DBConfig.User,
		Password: config.DBConfig.Pass,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Source: "admin",
	})
	if err != nil {
		panic(err)
	}
	return session.DB(config.DBConfig.Database)
}

func GetDBConn() *mgo.Database {
	return db
}
