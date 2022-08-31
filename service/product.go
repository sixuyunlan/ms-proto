package service

import (
	context "context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
)

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := p.GetStockById(request.ProdId)
	user := User{Username: "mszlu"}
	content := Content{
		Msg: "mszlu msg...",
	}
	any, _ := anypb.New(&content)

	return &ProductResponse{ProdStock: int32(stock), User: &user, Data: any}, nil
}
func (p *productService) UpdataProductStockClientStream(stream ProdService_UpdataProductStockClientStreamServer) error {
	count := 0
	for {
		recv, _ := stream.Recv()
		fmt.Println("stream is ", recv.ProdId, count)
		count++
		if count > 10 {
			rsp := &ProductResponse{ProdStock: recv.ProdId}
			stream.SendAndClose(rsp)
			return nil

		}
	}
}
func (p *productService) GetProductStockServerStream(request *ProductRequest, stream ProdService_GetProductStockServerStreamServer) error {
	count := 0
	for {
		rsp := &ProductResponse{ProdStock: request.ProdId}

		stream.Send(rsp)
		time.Sleep(time.Second)
		count++
		if count > 10 {
			return nil
		}
	}
}

func (p *productService) SayHelloStream(stream ProdService_SayHelloStreamServer) error {
	for {
		recv, err := stream.Recv()
		if err != nil {
			return nil
		}
		fmt.Println("服务端收到客户端的消息", recv.ProdId)
		time.Sleep(time.Second)
		rsp := &ProductResponse{ProdStock: recv.ProdId}
		stream.Send(rsp)

	}
}
func (p *productService) GetStockById(id int32) int32 {
	return id
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {}
