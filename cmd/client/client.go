package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/famorai/training/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to grpc server:%v", err)
	}
	defer connection.Close()
	client := pb.NewTokenServiceClient(connection) //instance
	// AddToken(client)
	// AddTokenDetail(client) // call application
	// AddTokens(client)
	AddTokenStreamBi(client)
}

func AddToken(client pb.TokenServiceClient) { //method
	req := &pb.Token{ //request
		Id:    "0",
		Name:  "Dent",
		Price: 1,
	}
	res, err := client.AddToken(context.Background(), req) //response
	if err != nil {                                        // check err
		log.Fatalf("Could not make grpc request:%v", err)
	}
	fmt.Println(res) //print result
}

func AddTokenDetail(client pb.TokenServiceClient) {
	req := &pb.Token{ //request
		Id:    "0",
		Name:  "Dent",
		Price: 1,
	}
	resultStream, err := client.AddTokenDetail(context.Background(), req) //response
	if err != nil {                                                       // check err
		log.Fatalf("Could not make grpc request:%v", err)
	}

	for { //loop stream data
		stream, err := resultStream.Recv()
		if err == io.EOF { // end of file
			break
		}

		if err != nil { // check err
			log.Fatalf("Could not receive the msg:%v", err)
		}
		fmt.Println("Stream:", stream.GetStatus(), "-", stream.GetToken())
	}
}

func AddTokens(client pb.TokenServiceClient) {
	reqs := []*pb.Token{
		// &pb.Token{
		{
			Id:    "tk1",
			Name:  "BTC",
			Price: 100000,
		},
		// &pb.Token
		{
			Id:    "tk2",
			Name:  "ETH",
			Price: 10000,
		},
		// &pb.Token
		{
			Id:    "tk3",
			Name:  "AXS",
			Price: 10000,
		},
		// &pb.Token
		{
			Id:    "tk4",
			Name:  "Vechain",
			Price: 1000,
		},
		// &pb.Token
		{
			Id:    "tk5",
			Name:  "MATIC",
			Price: 9000,
		},
	}

	stream, err := client.AddTokens(context.Background())
	if err != nil {
		log.Fatal("Error creating requests", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error receiving response", err)
	}
	fmt.Println(res)

}
func AddTokenStreamBi(client pb.TokenServiceClient) {
	stream, err := client.AddTokenStreamBi(context.Background())
	if err != nil {
		log.Fatalf("Error creating request stream: %v", err)
	}
	reqs := []*pb.Token{
		// &pb.Token{
		{
			Id:    "tk1",
			Name:  "BTC",
			Price: 100000,
		},
		// &pb.Token
		{
			Id:    "tk2",
			Name:  "ETH",
			Price: 10000,
		},
		// &pb.Token
		{
			Id:    "tk3",
			Name:  "AXS",
			Price: 10000,
		},
		// &pb.Token
		{
			Id:    "tk4",
			Name:  "Vechain",
			Price: 1000,
		},
		// &pb.Token
		{
			Id:    "tk5",
			Name:  "MATIC",
			Price: 9000,
		},
	}

	// var wait, for to create a channel that connect with go routines and control
	wait := make(chan int)

	// go routines implementation, anonymous func...background running. Sending data requests
	go func() {
		for _, req := range reqs {
			fmt.Println("Sending token: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	//receiving in paralello / concurrency data above and finish program, but we can create a channel above
	go func() {

		//infinit looping for receive data request
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error Receiving data: %v", err)
				break
			}
			fmt.Printf("Receiving Token: %v status: %v\n", res.GetToken().GetName(), res.GetStatus())
			fmt.Printf("Deleting Token: %v status: %v\n", res.GetToken().GetName(), res.GetStatus())
		}

		close(wait) //finish infinit loop above
	}()

	<-wait //finish channel

}
