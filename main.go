package main

import (
	"github.com/HAL-Future-Creation-Exhibition/go-nellow/router"
)

func main() {
	r := router.GetRouter()
	r.Run()
}
