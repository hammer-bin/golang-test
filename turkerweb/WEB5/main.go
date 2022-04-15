package main

import (
	"fmt"
	"golang-test/test01/example01/turkerweb/WEB5/myapp"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":3000", myapp.NewHandler())
	if err != nil {
		return
	} else {
		fmt.Println(err)
	}
}
