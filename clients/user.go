package clients

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/thisisrandom/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	Conn proto.TokenClient
}

type BasicAuth struct {
	Username string
	Password string
}

func (b BasicAuth) GetRequestMetaData(ctx context.Context, in ...string) (map[string]string, error) {
	auth := fmt.Sprintf("%s:%s", b.Username, b.Password)
	encrypted := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "Basic " + encrypted,
	}, nil
}


func InitTokenClient(url string) (*client, error) {
	log.Println(url)

	creds := grpc.WithTransportCredentials(
		insecure.NewCredentials(),
		grpc.WithPerRPCCredentials()
	)

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
