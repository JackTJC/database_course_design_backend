package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)

type CreateGoodsHandler struct {
	Req *pb_gen.CreateGoodsRequest
	Res *pb_gen.CreateGoodsResponse
	Ctx context.Context
}

func NewCreateGoodsHandler(ctx context.Context,req *pb_gen.CreateGoodsRequest)*CreateGoodsHandler{
	return &CreateGoodsHandler{
		Req: req,
		Ctx: ctx,
		Res: &pb_gen.CreateGoodsResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}

func (h *CreateGoodsHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	if err:=dal.CreateGoods(packer.PackGoodsPbToModel(h.Req.GetGoods()));err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
}