package main

import (
	"WEB16/app"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	//n.Use(negroni.NewStatic(http.Dir("./turkerweb/WEB16/public/")))
	n.UseHandler(m)

	log.Println("Start App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}

}
