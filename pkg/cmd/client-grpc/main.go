package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	v1 "github.com/fboccacini/goauth/pkg/api/google.golang.org/protobuf/v1"
)

const (
	apiVersion = "v1"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewGoAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	login := &v1.Login{
		Username: "arale",
		Password: "ncha",
	}

	loginReq := v1.LoginRequest{
		Api:   apiVersion,
		Login: login,
	}

	res, err := c.Signup(ctx, &loginReq)
	if err != nil {
		log.Fatalf("Register failed: %v", err)
	}
	log.Printf("Register result: <%+v>\n\n", res)

	res, err = c.Login(ctx, &loginReq)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login result: <%+v>\n\n", res)

	service := "test service"

	authReq := v1.AuthenticateRequest{
		Api:      apiVersion,
		Username: login.Username,
		Service:  service,
		Token:    res.Token,
	}
	res, err = c.Authenticate(ctx, &authReq)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	log.Printf("Authentication result: <%+v>\n\n", res)

}
