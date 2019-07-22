package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/minami14/speech-recognition-server/assets/web"
	"github.com/minami14/speech-recognition-server/hub"
)

func main() {
	f, err := web.Assets.Open("/web/index.html")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4096)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	html := string(buf[:n])

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := fmt.Fprintf(writer, html); err != nil {
			log.Println(err)
		}
	})

	h := hub.NewHub()

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		h.Register(conn)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
