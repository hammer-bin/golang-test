package main

import (
	"golang-test/test01/example01/turkerweb/WEB1/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
