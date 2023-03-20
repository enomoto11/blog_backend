package main

import (
	"blog/api/router"
)

func main() {
	r, entClient := router.InitControllers()
	defer entClient.Close()

	r.Run()
}
