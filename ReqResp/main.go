package main

import "ReqRep/router"

func main() {
	r := router.Router()
	r.Run("0.0.0.0:8080")
}
