package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/model"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)

type MCreateGoodsHandler struct {
	Ctx context.Context
	Req *pb_gen.MCreateGoodsRequest
	Res *pb_gen.MCreateGoodsResponse
}

func NewMCreateGodsHandler(ctx context.Context,req *pb_gen.MCreateGoodsRequest) *MCreateGoodsHandler {
	return &MCreateGoodsHandler{
		Ctx:  ctx,
		Req:  req,
		Res: &pb_gen.MCreateGoodsResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}
func (h *MCreateGoodsHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	h.checkParams()
	if h.Res.GetBaseResp().GetErrNo()!=pb_gen.ErrNo_ErrNo_Success{
		return
	}
	goodsModelList:=make([]*model.Goods,0,len(h.Req.GoodsList))
	for _, goods := range h.Req.GetGoodsList() {
		goodsModelList = append(goodsModelList, packer.PackGoodsPbToModel(goods))
	}
	if err:=dal.MCreateGoods(goodsModelList);err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,"MCreateGoodsHandler call MCreateGoods error")
		return
	}
	return
}

func (h *MCreateGoodsHandler) checkParams() {
	if len(h.Req.GetGoodsList())==0{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Req_Illegal,"illegal request")
		return
	}
}
