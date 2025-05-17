package main

import (
	"fmt"
	"module/router"
	"net/http"
)

func main() {
    r := router.Router()
	fmt.Println("Started server")
	http.ListenAndServe(":4000", r)
}
