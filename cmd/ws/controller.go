package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/markelca/toggles/flags"
)

var broadcast = make(chan Command)
var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
    CheckOrigin: customUpgrader,
}

// Just for local envs. It should not return always true on production applications.
// https://pkg.go.dev/github.com/gorilla/websocket?utm_source=godoc#hdr-Origin_Considerations
var customUpgrader = func(r *http.Request) bool { 
    return true
}

type Command struct {
    Command string `json:"command"`
    Data interface{} `json:"data"`
}

func (c Command) Run(flagService flags.FlagService) (string,error) {
    switch c.Command {
        case "get":
            key := c.Data.(string)
            value,err := flagService.Get(key)
            if err != nil {
                return "",err
            }
            return strconv.FormatBool(value),nil
        case "create":
        case "update":
        case "list":
        default:
            return "wrong!",nil
            
    }
    return "wrong2!",nil
}

type WSController struct {
    flagService flags.FlagService
}

func (ws WSController) Init(host string) {
    http.HandleFunc("/", handleWebSocket)
    go handleMessages(ws.flagService)
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
        var cmd Command
        err := conn.ReadJSON(&cmd)
        // cmd.Time = JSONTime(time.Now())
        log.Printf("debug: %v", cmd)
        log.Printf("Message received: %v", cmd)
        if err != nil {
            log.Println("Error reading message:", err)
            delete(clients, conn)
            break
        }
        broadcast <- cmd
    }
}

func handleMessages(flagService flags.FlagService) {
    for {
        cmd := <-broadcast
        for conn := range clients {
            response,_ := cmd.Run(flagService)
            err := conn.WriteJSON(response)
            if err != nil {
                log.Println("Error writing message:", err)
                conn.Close()
                delete(clients, conn)
            }
        }
    }
}
