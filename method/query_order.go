package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)

type QueryOrderHandler struct {
	Req *pb_gen.QueryOrderRequest
	Res *pb_gen.QueryOrderResponse
	Ctx context.Context
}

func NewQueryOrderHandler(ctx context.Context, req *pb_gen.QueryOrderRequest) *QueryOrderHandler {
	return &QueryOrderHandler{
		Req: req,
		Ctx: ctx,
		Res: &pb_gen.QueryOrderResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}

func (h *QueryOrderHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	h.checkParams()
	if h.Res.GetBaseResp().GetErrNo()!=pb_gen.ErrNo_ErrNo_Success{
		return
	}
	var orderId,goodsId,clientId,offset,limit *int32
	order:=h.Req.GetOrder()
	if order.GetOrderId()!=0{
		orderId=&order.OrderId
	}
	if order.GetGoodsId()!=0{
		goodsId=&order.GoodsId
	}
	if order.GetClientId()!=0{
		clientId=&order.ClientId
	}
	if h.Req.GetLimit()!=0{
		limit=&h.Req.Limit
	}
	offset=&h.Req.Offset
	orderList, err := dal.QueryOrder(orderId, goodsId, clientId,offset,limit)
	if err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
	goodsIDList:=make([]int32,0)
	res := make([]*pb_gen.Order,0)
	for _, order := range orderList {
		res = append(res, packer.PackOrderModelToPb(order))
		goodsIDList = append(goodsIDList, order.Id)
	}
	goodsMap,err:=dal.MGetGoods(goodsIDList)
	if err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
	for _, resOrder := range res {
		if goods,ok:=goodsMap[resOrder.GetGoodsId()];ok{
			resOrder.GoodsName=goods.Name
			resOrder.Price=goods.Price
		}
	}
	h.Res.OrderList=res
}

func (h *QueryOrderHandler) checkParams() {
	if h.Req.GetLimit()==0{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Req_Illegal,"illegal request")
		return
	}
}