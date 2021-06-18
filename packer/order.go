package packer

import (
	"github.com/JackTJC/backend/model"
	"github.com/JackTJC/backend/pb_gen"
)

func PackOrderPbToModel(orderPb *pb_gen.Order) *model.Order {
	return &model.Order{
		Id:       orderPb.GetOrderId(),
		GoodsId:  orderPb.GetGoodsId(),
		ClientId: orderPb.GetClientId(),
		Num:      orderPb.GetNum(),
		Status: int16(orderPb.GetOrderStatus()),
	}
}

func PackOrderModelToPb(orderModel *model.Order) *pb_gen.Order {
	return &pb_gen.Order{
		OrderId:     orderModel.Id,
		GoodsId:     orderModel.GoodsId,
		ClientId:    orderModel.ClientId,
		Num:         orderModel.Num,
		OrderStatus: pb_gen.OrderStatus(orderModel.Status),
	}
}
