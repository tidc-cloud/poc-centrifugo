package main

//https://centrifugal.dev/docs/server/server_api#grpc-example-for-go
import (
	"context"
	"fmt"
	"log"
	"time"

	"centrifugo_example/apiproto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := apiproto.NewCentrifugoApiClient(conn)
	subRes, err := client.Subscribe(context.Background(), &apiproto.SubscribeRequest{
		User:    "1234",
		Channel: "public:1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("subRes", subRes)
	for {
		resp, err := client.Publish(context.Background(), &apiproto.PublishRequest{
			Channel: "public:1",
			Data:    []byte(`{"value": "hello from GRPC"}`),
		})
		if err != nil {
			log.Printf("Transport level error: %v", err)
		} else {
			if resp.GetError() != nil {
				respError := resp.GetError()
				log.Printf("Error %d (%s)", respError.Code, respError.Message)
			} else {
				log.Println("Successfully published")
			}
		}
		time.Sleep(time.Second)
	}
}
