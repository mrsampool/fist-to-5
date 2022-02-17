package main

import (
	"fmt"
	"github.com/mrsampool/fist-to-5/server"
)

func main() {
	router := server.SetupServer()
	err := router.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
