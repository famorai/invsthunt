package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/famorai/training/pb"
)

type TokenService struct {
	// if add a service that not exist, your servidor is not to be crash.. this is a protection
	pb.UnimplementedTokenServiceServer
}

// Constructor
func NewTokenService() *TokenService {
	return &TokenService{}
}

// type TokenServiceClient interface {
// 	AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
// 	AddTokenDetail(ctx context.Context, in *Token, opts ...grpc.CallOption) (TokenService_AddTokenDetailClient, error)
// }	AddTokens(ctx context.Context, opts ...grpc.CallOption) (TokenService_AddTokensClient, error)
// AddTokenStreamBi(ctx context.Context, opts ...grpc.CallOption) (TokenService_AddTokenStreamBiClient, error)

func (*TokenService) AddToken(ctx context.Context, reqToken *pb.Token) (*pb.Token, error) {
	//insert database exemple
	fmt.Println(reqToken.Name)

	return &pb.Token{
		Id:    "318798",
		Name:  reqToken.GetName(),
		Price: reqToken.GetPrice(),
	}, nil
}

func (*TokenService) AddTokenDetail(req *pb.Token, stream pb.TokenService_AddTokenDetailServer) error {
	stream.Send(&pb.TokenResultStream{ // iniciating send stream data
		Status: "Init",
		Token:  &pb.Token{},
	})
	time.Sleep(time.Second * 3) // go sleep for 3 seconds

	stream.Send(&pb.TokenResultStream{ // inserting send stream data
		Status: "Inserting DB",
		Token:  &pb.Token{},
	})
	time.Sleep(time.Second * 3) // go sleep for 3 seconds

	stream.Send(&pb.TokenResultStream{ // inserting send stream data
		Status: "Token has bee inserted",
		Token: &pb.Token{
			Id:    "318798",
			Name:  req.GetName(),
			Price: req.GetPrice(),
		},
	})
	time.Sleep(time.Second * 3) // go sleep for 3 seconds

	stream.Send(&pb.TokenResultStream{ // inserting send stream data
		Status: "Completed",
		Token: &pb.Token{
			Id:    "318798",
			Name:  req.GetName(),
			Price: req.GetPrice(),
		},
	})
	time.Sleep(time.Second * 3) // go sleep for 3 seconds

	return nil

}

func (*TokenService) AddTokens(stream pb.TokenService_AddTokensServer) error {
	tokens := []*pb.Token{} // slice tokens

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Tokens{
				Token: tokens,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving stream:%v", err)
		}

		tokens = append(tokens, &pb.Token{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Price: req.GetPrice(),
		})

		fmt.Println("Adding", req.GetName(), req.GetPrice())
	}
}

func (*TokenService) AddTokenStreamBi(stream pb.TokenService_AddTokenStreamBiServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.TokenResultStream{
			Status: "Added",
			Token:  req,
		})
		if err != nil {
			log.Fatalf("Error Sending stream to the client: %v", err)
		}
		err = stream.Send(&pb.TokenResultStream{
			Status: "Deleted",
			Token:  req,
		})
		if err != nil {
			log.Fatalf("Error Sending stream to the client: %v", err)
		}

	}
}
