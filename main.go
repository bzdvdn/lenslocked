package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Welcome</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get information, please send an email <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	// mux := &http.ServeMux{} //default handler ServerMux
	// mux.HandleFunc("/", handlerFunc)
	// gmux := mux.NewRouter()
	// router := httprouter.New() // init router
	// router.GET("/hello/:name", Hello)
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	// http.HandleFunc("/", handlerFunc)
	// gmux.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", router) // if you have router or mux, you need to include this
}
