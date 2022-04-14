package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toy/terraman/action"
)

var iaas = map[string]*IaaS{}

type IaaS struct {
	csp string `json:"csp"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	})
}

func rest() {
	mux := http.NewServeMux()

	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(wr).Encode(iaas)
			fmt.Println("user : ", iaas)
			action.GetTodo() // 신호를 받았을 때 액션 처리
		case http.MethodPost:
			var infra IaaS
			json.NewDecoder(r.Body).Decode(&infra)
			iaas[infra.csp] = &infra
			json.NewEncoder(wr).Encode(infra)
		}
	})

	mux.Handle("/", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":8080", mux)
}

func main() {

	fmt.Println("Start process...")
	rest() // API 서버 동작

	fmt.Println("End process...")
}
