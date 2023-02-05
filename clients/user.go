package clients

import (
	"context"
	"log"

	"github.com/thisisrandom/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	Conn proto.TokenClient
}

func InitTokenClient(url string) (*client, error) {
	log.Println(url)

	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return &client{
		Conn: proto.NewTokenClient(cc),
	}, nil
}

func (c *client) GenerateToken(username string, password string) (*proto.CreateTokenResponse, error) {
	req := &proto.CreateTokenRequest{
		Name:     username,
		Password: password,
	}

	return c.Conn.CreateToken(context.Background(), req)
}
