package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)

type UpdateOrderStatHandler struct {
	ctx context.Context
	Req *pb_gen.UpdateOrderStatusRequest
	Res *pb_gen.UpdateOrderStatusResponse
}

func NewUpdateOrderStatHandler(ctx context.Context, req *pb_gen.UpdateOrderStatusRequest) *UpdateOrderStatHandler {
	return &UpdateOrderStatHandler{
		ctx: ctx,
		Req: req,
		Res: &pb_gen.UpdateOrderStatusResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}

func (h *UpdateOrderStatHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	err:=dal.UpdateStat(h.Req.GetOrderId(), int32(h.Req.GetTargetStatus()))
	if err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
}