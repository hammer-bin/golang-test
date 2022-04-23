package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
	"web16-2/app"
)

func main() {
	//m := app.MakeHandler(flag.Args())
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
