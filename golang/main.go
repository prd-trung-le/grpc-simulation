package main

import (
	"fmt"
	"log"
	"net"
	"crypto/tls"

	"auth/auth"
	"auth/config"
	"auth/repository"
	"auth/usecase"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	server_cert, err := tls.LoadX509KeyPair("ssl_cert/server-cert.pem", "ssl_cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{server_cert},
		ClientAuth: tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	db := config.ConnectDB()
	userRepository := repository.InitUserRepository(db)
	userUsecase := usecase.InitUserUsecase(userRepository)

	s := auth.InitServer(userUsecase)

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9009))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
	}
	fmt.Println("Listen to port 9009")

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
	}
}
