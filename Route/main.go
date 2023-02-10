package main

import "Route/router"

func main() {
	srv := router.InitRoute()
	srv.Run(":8080")
}
