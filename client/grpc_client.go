package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"ms-proto/client/auth"
	"ms-proto/service"
	"time"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

func main() {
	//creds, err := credentials.NewClientTLSFromFile("cert/server.pem", "*.mszlu.com")
	//if err != nil {
	//	log.Fatalf("证书错误", err)
	//}
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
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
		ServerName:   "*.mszlu.com",
		RootCAs:      certPool,
	})

	token := &auth.Authentication{
		User:     "admin",
		Password: "admin",
	}

	conn, _ := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(token))
	defer conn.Close() //退出时关闭链接，defer=延时，使推迟

	prodClient := service.NewProdServiceClient(conn)

	//stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("success", stockResponse.ProdStock, stockResponse.User, stockResponse.Data)
	//	stream, err := prodClient.UpdataProductStockClientStream(context.Background())
	//	if err != nil {
	//		log.Fatalf("stream is err,err =", err)
	//	}
	//	rsp := make(chan struct{}, 1)
	//	go prodRequest(stream, rsp)
	//	select {
	//	case <-rsp:
	//		recv, _ := stream.CloseAndRecv()
	//		stock := recv.ProdStock
	//		fmt.Println("收到响应", stock)
	//	}
	//

	//request := &service.ProductRequest{
	//	ProdId: 123,
	//}
	//stream, _ := prodClient.GetProductStockServerStream(context.Background(), request)
	//for {
	//	recv, err := stream.Recv()
	//	if err != nil {
	//		if err == io.EOF {
	//			fmt.Println("客户端数据接收完成")
	//			err := stream.CloseSend()
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//		}
	//		log.Fatal(err)
	//	}
	//	fmt.Println("客户端收到的流", recv.ProdStock)
	//
	//}
	stream, _ := prodClient.SayHelloStream(context.Background())
	for {
		request := &service.ProductRequest{
			ProdId: 123,
		}
		stream.Send(request)
		recv, err := stream.Recv()
		if err != nil {

			log.Fatal(err)
		}
		fmt.Println("客户端收到的流信息", recv.ProdStock)
		time.Sleep(time.Second)

	}
}

//func prodRequest(stream service.ProdService_UpdataProductStockClientStreamClient, rsp chan struct{}) {
//	count := 0
//tatile1:
//	for {
//		request := &service.ProductRequest{
//			ProdId: 123,
//		}
//		stream.Send(request)
//		time.Sleep(time.Second)
//		count++
//		if count > 10 {
//			rsp <- struct{}{}
//			break tatile1
//		}
//	}
//}
