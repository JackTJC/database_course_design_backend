package main

import (
	"context"
	"github.com/JackTJC/backend/pb_gen"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const address = "localhost:6789"

func TestCreateClient(t *testing.T) {
	conn,err:=grpc.Dial(address,grpc.WithBlock(),grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("dia error:%v",err)
	}
	defer conn.Close()
	c:=pb_gen.NewGoodsInfoAdminClient(conn)
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	res,err:=c.CreateClient(ctx,&pb_gen.CreateClientRequest{})
	if err!=nil{
		log.Fatalf("rpc call CreateClient error:%v",err)
	}
	log.Printf("%v",res)
}
