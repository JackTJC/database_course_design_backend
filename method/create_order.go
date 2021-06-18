package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)

type CreateOrderHandler struct {
	Req *pb_gen.CreateOrderRequest
	Res *pb_gen.CreateOrderResponse
	Ctx context.Context
}

func NewCreateOrderHandler(ctx context.Context,req *pb_gen.CreateOrderRequest)*CreateOrderHandler{
	return &CreateOrderHandler{
		Req: req,
		Ctx: ctx,
		Res: &pb_gen.CreateOrderResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}

func (h *CreateOrderHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	if err:=dal.CreateOrder(packer.PackOrderPbToModel(h.Req.GetOrder()));err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
}