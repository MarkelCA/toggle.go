package main

import (
	"log"
	"net/http"
)

// Just for local envs. It should not return always true on production applications.
// https://pkg.go.dev/github.com/gorilla/websocket?utm_source=godoc#hdr-Origin_Considerations
var customUpgrader = func(r *http.Request) bool { 
    return true
}

type Command struct {
    command string
    data interface{}
}

func InitWS(host string) {
    http.HandleFunc("/", handleWebSocket)
    go handleMessages()
    log.Printf("Starting server on %v...",host)
    log.Fatal(http.ListenAndServe(host, nil))

}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println("Error upgrading connection:", err)
        return
    }
    defer conn.Close()

    clients[conn] = true

    for {
        var msg interface{}
        err := conn.ReadJSON(&msg)
        // msg.Time = JSONTime(time.Now())
        log.Printf("debug: %v", msg)
        log.Printf("Message received: %v", msg)
        if err != nil {
            log.Println("Error reading message:", err)
            delete(clients, conn)
            break
        }
        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        for conn := range clients {
            err := conn.WriteJSON(msg)
            if err != nil {
                log.Println("Error writing message:", err)
                conn.Close()
                delete(clients, conn)
            }
        }
    }
}
