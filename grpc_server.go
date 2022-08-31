package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"ms-proto/service"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		log.Fatalf("证书错误", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		log.Fatalf("ca err", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		err = Auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))
	service.RegisterProdServiceServer(rpcServer, service.ProductService)
	listener, _ := net.Listen("tcp", ":8002")
	rpcServer.Serve(listener)
}

func Auth(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	var user string
	var password string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "admin" || password != "admin" {
		return status.Errorf(codes.Unauthenticated, "token err")
	}

	return nil
}
