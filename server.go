package main

/*
	export GOPATH=$HOME/work
	Instalación 
	go get github.com/gorilla/websocket
*/
import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    cssHandler := http.FileServer(http.Dir("./css/"))
    jsHandler := http.FileServer(http.Dir("./js/"))

    http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
    http.Handle("/js/", http.StripPrefix("/js/", jsHandler))

    r.HandleFunc("/", HomeHandler)
    
    http.Handle("/", r)
    log.Println("Server running on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
    http.ServeFile(rw, r, "index.html")
}