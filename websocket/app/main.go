package app

import (
	"context"
	clientgrpcserver "ekzamen_5/websocket/client_grpc_server"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	clients     map[string]*websocket.Conn
	clientsLock sync.Mutex
	redisClient *redis.Client
	grpcClient  clientgrpcserver.ServiceClient
}

func NewServer(gprcClient clientgrpcserver.ServiceClient) *Server {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Conn errr: %v", err)
	}
	server := &Server{
		grpcClient:  gprcClient,
		clients:     make(map[string]*websocket.Conn),
		redisClient: redisClient,
	}
	go server.listenForNotifications()
	return server
}

func (s *Server) HandlerNotification(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	id := r.Header.Get("id")
	if id == "" {
		log.Println("id  not found ")
		return
	}

	s.clientsLock.Lock()
	s.clients[id] = conn
	s.clientsLock.Unlock()
	res, err := s.grpcClient.NotificationService().GetNotification(context.Background(), &notificationpb.GetNotificationReq{
		UserId: id,
	})
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(res.Messages); i++ {
		err := conn.WriteJSON(res.Messages[i])
		if err != nil {
			log.Println(err)
			return
		}
	}
	defer func() {
		s.clientsLock.Lock()
		delete(s.clients, id)
		s.clientsLock.Unlock()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("conn err:", err)
			break
		}
	}
}

func (s *Server) listenForNotifications() {
	pubsub := s.redisClient.Subscribe(context.Background(), "notifications")
	defer pubsub.Close()
	for {
		_, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println("conn err:", err)
			continue
		}

		s.clientsLock.Lock()
		for id, conn := range s.clients {
			res, err := s.grpcClient.NotificationService().GetNotification(context.Background(), &notificationpb.GetNotificationReq{
				UserId: id,
			})
			if err != nil {
				log.Println("conn err", id, ":", err)
				continue
			}

			if len(res.Messages) == 0 {
				continue
			}

			for _, msg := range res.Messages {
				err = conn.WriteJSON(msg)
				if err != nil {
					log.Println("conn err", id, ":", err)
					continue
				}
			}
		}
		s.clientsLock.Unlock()
	}
}
