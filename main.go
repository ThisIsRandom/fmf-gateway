package main

import (
	"log"

	"github.com/thisisrandom/fmf-gateway/clients"
)

func main() {
	userClient, err := clients.InitTokenClient(":8080")

	if err != nil {
		log.Fatalf("err: %s", err)
	}

	res, err := userClient.GenerateToken("test", "test")

	if err != nil {
		log.Fatalf("err: %s", err)
	}

	log.Printf("token: %s", res.Token)
}

/* grpc.NewServer(
    grpc.StreamInterceptor(streamInterceptor),
    grpc.UnaryInterceptor(unaryInterceptor)
)

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
    if err := authorize(stream.Context()); err != nil {
        return err
    }

    return handler(srv, stream)
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    if err := authorize(ctx); err != nil {
        return err
    }

    return handler(ctx, req)
}

func authorize(ctx context.Context) error {
    if md, ok := metadata.FromContext(ctx); ok {
        if len(md["username"]) > 0 && md["username"][0] == "admin" &&
            len(md["password"]) > 0 && md["password"][0] == "admin123" {
            return nil
        }

        return AccessDeniedErr
    }

    return EmptyMetadataErr
} */
