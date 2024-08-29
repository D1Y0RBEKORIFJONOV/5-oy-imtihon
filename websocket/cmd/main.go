package main

import (
	"ekzamen_5/websocket/app"
	clientgrpcserver "ekzamen_5/websocket/client_grpc_server"
	"ekzamen_5/websocket/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.New()
	grpcCliet, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		log.Fatal(err)
	}
	server := app.NewServer(grpcCliet)

	http.HandleFunc("/ws", server.HandlerNotification)
	log.Println("Listening on :9005")
	log.Fatal(http.ListenAndServe("ws:9005", nil))
}
