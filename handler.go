package main

import (
	"context"
	"github.com/JackTJC/backend/method"
	"github.com/JackTJC/backend/pb_gen"
)

type GoodsInfoAdminServer struct {
	pb_gen.UnimplementedGoodsInfoAdminServer
}

func (s *GoodsInfoAdminServer) CreateGoods(ctx context.Context, request *pb_gen.CreateGoodsRequest) (*pb_gen.CreateGoodsResponse, error) {
	h:=method.NewCreateGoodsHandler(ctx,request)
	h.Run()
	return h.Res,nil
}

func (s *GoodsInfoAdminServer) QueryGoods(ctx context.Context, request *pb_gen.QueryGoodsRequest) (*pb_gen.QueryGoodsResponse, error) {
	h:=method.NewQueryGoodsHandler(ctx,request)
	h.Run()
	return h.Res,nil
}

func (s *GoodsInfoAdminServer) CreateClient(ctx context.Context, request *pb_gen.CreateClientRequest) (*pb_gen.CreateClientResponse, error) {
	h:=method.NewCreateClientHandler(ctx,request)
	h.Run()
	return h.Res,nil
}

func (s *GoodsInfoAdminServer) QueryClient(ctx context.Context, request *pb_gen.QueryClientRequest) (*pb_gen.QueryClientResponse, error) {
	h:=method.NewQueryClientHandler(ctx,request)
	h.Run()
	return h.Res,nil
}

func (s *GoodsInfoAdminServer) CreateOrder(ctx context.Context, request *pb_gen.CreateOrderRequest) (*pb_gen.CreateOrderResponse, error) {
	h:=method.NewCreateOrderHandler(ctx,request)
	h.Run()
	return h.Res,nil
}

func (s *GoodsInfoAdminServer) QueryOrder(ctx context.Context, request *pb_gen.QueryOrderRequest) (*pb_gen.QueryOrderResponse, error) {
	h:=method.NewQueryOrderHandler(ctx,request)
	h.Run()
	return h.Res,nil
}

