// websockets.go
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		file, err := os.Open("music/output.opus")
		if err != nil {
			return
		}
		// Read message from browser
		buffer := make([]byte, 100)
		_, _, err = conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		// Print the message to the console
		file.Read(buffer)
		for {
			// Write message back to browser
			_, err := file.Read(buffer)
			if err != nil {
				break
			}
			fmt.Println(string(buffer))
			if err = conn.WriteMessage(websocket.BinaryMessage, buffer); err != nil {
				return
			}
			time.Sleep(time.Millisecond * 2)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
