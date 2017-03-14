package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/pankona/grpctest/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "50052"
)

type chatserver struct {
	chat.ChatServer
}

func (server *chatserver) SendRecv(stream chat.Chat_SendRecvServer) error {
	log.Println("server SendRecv() function started")

	done := make(chan struct{})

	// recv
	go func() {
		for {
			message, err := stream.Recv()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(message.GetMessage())
			}
		}
	}()

	// send
	go func() {
		for {
			time.Sleep(time.Second * 1)
			err := stream.Send(&chat.Message{Message: "message from server"})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("message sent to client")
			}
		}
	}()

	<-done
	fmt.Println("server SendRecv() function done")
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	chat.RegisterChatServer(server, &chatserver{})
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
