package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"time"
)

// Result member 변수는 대문자로 시작해야 한다.
type Result struct {
	Status    string    `json:"status"`
	Csp       string    `json:"csp"`
	CreatedAt time.Time `json:"created_at"`
}

var rd *render.Render

func receiveDataHandler(w http.ResponseWriter, r *http.Request) {
	result := new(Result)
	err := json.NewDecoder(r.Body).Decode(result)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	result.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, result)
	log.Printf(result.Csp)
}

func main() {
	rd = render.New()
	mux := pat.New()
	mux.Post("/result", receiveDataHandler)
	log.Println("start Server")
	http.ListenAndServe(":3001", mux)
}
