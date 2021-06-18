package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)

type QueryGoodsHandler struct {
	Req *pb_gen.QueryGoodsRequest
	Res *pb_gen.QueryGoodsResponse
	Ctx context.Context
}

func NewQueryGoodsHandler(ctx context.Context,req *pb_gen.QueryGoodsRequest) *QueryGoodsHandler {
	return &QueryGoodsHandler{
		Req: req,
		Ctx: ctx,
		Res: &pb_gen.QueryGoodsResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}

func (h *QueryGoodsHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	h.checkParams()
	if h.Res.GetBaseResp().GetErrNo()!=pb_gen.ErrNo_ErrNo_Success{
		return
	}
	var goodsType *int
	var priceHigh,priceLow,offset,limit *int32
	var name,desc *string
	if h.Req.GetPriceHigh()!=0{
		priceHigh=&h.Req.PriceHigh
	}
	if h.Req.GetPriceLow()!=0{
		priceLow=&h.Req.PriceLow
	}
	if len(h.Req.GetName())!=0{
		name=&h.Req.Name
	}
	if len(h.Req.GetDescription())!=0{
		desc=&h.Req.Description
	}
	if t:=int(h.Req.GetGoodsType());t!=0{
		goodsType=&t
	}
	if h.Req.GetLimit()!=0{
		limit=&h.Req.Limit
	}
	offset=&h.Req.Offset
	goodsList, err := dal.QueryGoods(goodsType, name, desc, priceLow, priceHigh,offset,limit)
	if err != nil {
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
	res:=make([]*pb_gen.Goods,0)
	for _, goods := range goodsList {
		res = append(res, packer.PackGoodsModelToPb(goods))
	}
	h.Res.GoodsList=res
}

func (h *QueryGoodsHandler) checkParams() {
	if h.Req.GetLimit()==0{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Req_Illegal,"illegal request")
		return
	}
}