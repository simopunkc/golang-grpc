package main

import (
	"context"
	"fmt"

	"golang-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:4040",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewAddServiceClient(conn)

	const (
		a = 10
		b = 10
	)

	req := &proto.Request{A: int64(a), B: int64(b)}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Result)
}
