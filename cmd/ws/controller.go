package main


//
import (
// 	"encoding/json"
	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/websocket"
// 	"github.com/markelca/toggles/flags"
// 	"github.com/markelca/toggles/storage"
)
//
// var broadcast = make(chan Command)
// var clients = make(map[*websocket.Conn]bool)
// var upgrader = websocket.Upgrader{
//     CheckOrigin: customUpgrader,
// }
//
// // Just for local envs. It should not return always true on production applications.
// // https://pkg.go.dev/github.com/gorilla/websocket?utm_source=godoc#hdr-Origin_Considerations
// var customUpgrader = func(r *http.Request) bool {
//     return true
// }
//
// type WSController struct {
//     flagService flags.FlagService
//     cacheClient storage.CacheClient
// }
//
// func (ws WSController) Init(host string) {
//     http.HandleFunc("/", ws.handleWebSocket)
//     go handleMessages(ws)
//     log.Printf("Starting server on %v...",host)
//     log.Fatal(http.ListenAndServe(host, nil))
// }
//
//
type Command struct {
    Command string `json:"command"`
    Data interface{} `json:"data"`
    broadcast bool
    emmiter *Client
}

type Status int

const (
	StatusSuccess Status = 200
	StatusCreated Status = 201

	StatusInternalServerError Status = 500

    StatusBadRequest    Status = 400
    StatusNotFound      Status = 404
	StatusConflict      Status = 409
)

type Response struct {
    Status Status `json:"status"`
    Value interface{} `json:"value"`
}

type ClientResponse struct {
    Response
    Client *Client
}

// func (ws WSController) Update(cmd *Command) Response {
//     flag,err := ParseFlag(cmd.Data)
//     if err != nil {
//         return Response{StatusInternalServerError,err}
//     }
//     err = ws.flagService.Update(flag.Name,flag.Value)
//     if err != nil {
//         if err == flags.ErrFlagNotFound {
//             return Response{StatusNotFound,err}
//         }
//         return Response{StatusInternalServerError, err}
//     }
//
//     return Response{StatusCreated,nil}
// }
//
func (cmd *Command) Run(/*ws WSController*/) Response {
    var str string
    switch cmd.Command {
        case "get":
            str = "running get"
            // return ws.Get(cmd)
        case "create":
            str = "running create"
            // return ws.Create(cmd)
        case "update":
            str = "running update"
            // return ws.Update(cmd)
        case "delete":
            str = "running delete"
            // return ws.Delete(cmd)
        default:
            msg := fmt.Sprintf("Invalid command (%v)",cmd.Command) 
            return Response{StatusBadRequest,msg}
    }
    return Response{StatusSuccess,str}
}
//
// func (ws WSController) Get(c *Command) Response {
//         if c.Data == nil {
//             flags,err := ws.flagService.List()
//             if err != nil {
//                 return Response{StatusInternalServerError,err}
//             }             
//             return Response{StatusSuccess,flags}
//         }
//         key := c.Data.(string)
//         value,err := ws.flagService.Get(key)
//         if err != nil {
//             if err == flags.ErrFlagNotFound{
//                 return Response{StatusNotFound,err}
//             }
//             return Response{StatusInternalServerError,err}
//         }
//         return Response{StatusSuccess,value}
// }
//
// func ParseFlag(data interface{}) (*flags.Flag,error) {
//     jsonBody,err := json.Marshal(data)
//     if err != nil {
//         return nil,err
//     }
//     var flag flags.Flag
//     if err = json.Unmarshal(jsonBody, &flag); err != nil {
//         return nil,err
//     }
//     return &flag,nil
// }
//
// func (ws WSController) Create(cmd *Command) Response {
//     flag,err := ParseFlag(cmd.Data)
//     if err != nil {
//         return Response{StatusInternalServerError,err}
//     }
//     err = ws.flagService.Create(*flag)
//     if err != nil {
//         if err == flags.ErrFlagAlreadyExists {
//             return Response{StatusConflict,err}
//         }
//         return Response{StatusInternalServerError,err}
//     }
//     return Response{StatusCreated,flag}
// }
//
// func (ws WSController) Delete(cmd *Command) Response {
//     key := fmt.Sprintf("%v",cmd.Data)
//     err := ws.flagService.Delete(key)
//     if err != nil {
//         if err == flags.ErrFlagNotFound {
//             return Response{StatusNotFound,nil}
//         }
//         return Response{StatusInternalServerError,nil}
//     }
//     return Response{StatusSuccess,nil}
// }
//
// func (ws WSController) handleWebSocket(w http.ResponseWriter, r *http.Request) {
//     conn, err := upgrader.Upgrade(w, r, nil)
//
//     if err != nil {
//         log.Println("Error upgrading connection:", err)
//         return
//     }
//     defer conn.Close()
//
//     clients[conn] = true
//
//     for {
//         var cmd Command
//         err := conn.ReadJSON(&cmd)
//         if err != nil {
//             log.Println("Error reading message:", err)
//             delete(clients, conn)
//             break
//         }
//         cmd.emmiter = conn
//         broadcast <- cmd
//     }
// }
//
// func handleMessages(ws WSController) {
//     for {
//         cmd := <-broadcast
//         response := cmd.Run(ws)
//
//         if cmd.broadcast {
//             log.Println("(Broadcasted)")
//             for conn := range clients {
//                 err := conn.WriteJSON(response)
//                 if err != nil {
//                     log.Println("Error writing message:", err)
//                     conn.Close()
//                     delete(clients, conn)
//                 }
//             }
//         } else {
//             log.Println("(NOT Broadcasted)")
//             err := cmd.emmiter.WriteJSON(response)
//             if err != nil {
//                 log.Println("Error writing message:", err)
//                 cmd.emmiter.Close()
//                 delete(clients, cmd.emmiter)
//             }
//         }
//     }
// }
