package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "suslmk", Email: "suslmk@naver.com"}

	rd.JSON(w, http.StatusOK, user)
	/*w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))*/

}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	//log.Print("[LOGGER1] Started")
	err := json.NewDecoder(r.Body).Decode(user) //user에 값을 채워준다.
	if err != nil {
		/*w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)*/
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
	/*w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))*/
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "suslmk", Email: "suslmk@naver.com"}
	/*tmpl, err := template.New("Hello").ParseFiles("D:/workspaceGolang/golang-test/turkerweb/WEB12/templates/hello.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		rd.Text(w, http.StatusInternalServerError, err.Error())
		return
	}*/
	rd.HTML(w, http.StatusOK, "body", user)
	//tmpl.ExecuteTemplate(w, "hello.tmpl", "Suslmk")
}

func main() {
	rd = render.New(render.Options{
		Directory:  "D:/workspaceGolang/golang-test/turkerweb/WEB12/templates",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello2",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	//mux.HandleFunc("/users", getUserInfoHandler).Methods("GET")
	//mux.HandleFunc("/users", addUserHandler).Methods("POST")

	http.ListenAndServe(":3000", n)
}
