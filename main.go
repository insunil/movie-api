package main

import (
	"fmt"
	"module/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Started")

}
