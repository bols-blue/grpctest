package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"chat"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	client := chat.NewChatClient(conn)
	stream, err := client.SendRecv(context.Background())

	done := make(chan struct{})

	// recv
	go func() {
		for {
			if message, err := stream.Recv(); err != nil {
				log.Println(err)
			} else {
				fmt.Println(message.GetMessage())
			}
		}
	}()

	// send
	go func() {
		for {
			time.Sleep(time.Second * 1)
			if err := stream.Send(&chat.Message{Message: "message from client"}); err != nil {
				log.Println(err)
			} else {
				fmt.Println("message sent to server")
			}
		}
	}()

	<-done
	fmt.Println("done")
	os.Exit(0)
}
