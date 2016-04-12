package main

/*
	export GOPATH=$HOME/work
	Instalación 
    go get github.com/gorilla/mux
    go get github.com/gorilla/websocket
*/
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
    "sync"
)

var Users = make(map[User]string)
var UsersRWMutex sync.RWMutex

type User struct {
    websocket *websocket.Conn
}

func main() {
    mux := mux.NewRouter()
    cssHandler := http.FileServer(http.Dir("./css/"))
    jsHandler := http.FileServer(http.Dir("./js/"))
    
    mux.HandleFunc("/", HomeHandler).Methods("GET")
    mux.HandleFunc("/ws", web_socket)

    http.Handle("/", mux)
    http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
    http.Handle("/js/", http.StripPrefix("/js/", jsHandler))

    log.Println("Server running on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func web_socket(w http.ResponseWriter, r *http.Request){
    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
    if err != nil {
        log.Println(err)
        return
    }
    user := create_user(ws, "Eduardo")
    add_user(user)
    for{
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            return
        }
        send_echo(messageType, message)
    }
}

func create_user(ws *websocket.Conn, usuario string) User{
    return User{ websocket: ws}
}

func add_user(user User){
    UsersRWMutex.Lock()
    defer UsersRWMutex.Unlock()
    Users[user] = ""
}

func send_echo(messageType int, message []byte) {
    UsersRWMutex.RLock()
    defer UsersRWMutex.RUnlock()
    log.Println("To send messages")
    for user, _ := range Users {
        if err := user.websocket.WriteMessage(messageType, message); err != nil {
            return
        }
    }
}
