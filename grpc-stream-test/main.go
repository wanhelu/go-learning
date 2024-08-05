package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

const (
	address = "localhost:12121"
	method  = "/product.ProductInfo/addProduct"
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
	stream, err := conn.NewStream(context.Background(), &grpc.StreamDesc{ServerStreams: false, ClientStreams: false}, method)
	err = stream.SendMsg(&emptypb.Empty{})
	if err != nil {
		log.Println("send failed", err)
		return
	}
	res := &emptypb.Empty{}
	err = stream.RecvMsg(res)
	if err != nil {
		log.Println("recv failed", err)
		return
	}
	log.Println("res: %v", res)

	err = stream.SendMsg(&emptypb.Empty{})
	if err != nil {
		log.Println("send failed", err)
		return
	}
	err = stream.RecvMsg(res)
	if err != nil {
		log.Println("recv failed", err)
		return
	}
	log.Println("res: %v", res)
}
