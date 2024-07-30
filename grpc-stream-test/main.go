package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"grpc-demo/product"
	"log"
	"time"
)

const (
	address = "localhost:12121"
)

func main() {
	conn, err := grpc.Dial("localhost:12121", grpc.WithInsecure(), grpc.WithTimeout(time.Second*30))
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	//go func() {
	//	for{
	//		log.Println(conn.GetState())
	//		time.Sleep(time.Second)
	//	}
	//}()
	defer conn.Close()
	for {
		state := conn.GetState()
		for state != connectivity.Ready {
			log.Println("before:" + state.String())
			conn.Connect()
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			conn.WaitForStateChange(ctx, state)
			state = conn.GetState()
			log.Println("after:" + state.String())
		}
		time.Sleep(time.Second * 10)
		//stream,err:=conn.NewStream(context.Background(),&grpc.StreamDesc{ServerStreams: false, ClientStreams: false},"/product.ProductInfo/addProduct")
		//err=stream.SendMsg(&product.Product{Name: "1"})
		//if err != nil {
		//	log.Println("did not connect.", err)
		//	return
		//}
	}

	log.Println("success")
}

// 添加一个测试的商品
func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	aMac := &product.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success, id = ", productId.Value)
	return productId.Value
}

// 获取一个商品
func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	p, err := client.GerProduct(ctx, &product.ProductId{Value: id})
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Printf("get prodcut success : %+v\n", p)
}
