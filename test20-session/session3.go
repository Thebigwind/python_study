package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "s1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(session)
	session.Values["name"] = "spuerWang"
	session.Save(r, w)
}

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/session", MyHandler)
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}
