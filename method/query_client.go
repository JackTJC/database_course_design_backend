package method

import (
	"context"
	"log"

	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
)

type QueryClientHandler struct {
	Req *pb_gen.QueryClientRequest
	Res *pb_gen.QueryClientResponse
	Ctx context.Context
}

func NewQueryClientHandler(ctx context.Context, req *pb_gen.QueryClientRequest) *QueryClientHandler {
	return &QueryClientHandler{
		Req: req,
		Ctx: ctx,
		Res: &pb_gen.QueryClientResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
	}
}

func (h *QueryClientHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	h.checkParam()
	if h.Res.GetBaseResp().GetErrNo() != pb_gen.ErrNo_ErrNo_Success {
		return
	}
	var id, limit, offset *int32
	var tel *string
	if len(h.Req.GetTel()) != 0 {
		tel = &h.Req.Tel
	}
	if h.Req.GetId() != 0 {
		id = &h.Req.Id
	}
	if h.Req.GetLimit() != 0 {
		limit = &h.Req.Limit
	}
	offset = &h.Req.Offset
	clientList, err := dal.QueryClient(tel, id, offset, limit)
	if err != nil {
		h.Res.BaseResp = util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed, err.Error())
		return
	}
	res := make([]*pb_gen.Client, 0)
	for _, client := range clientList {
		res = append(res, packer.PackClientModelToPb(client))
	}
	h.Res.ClientList = res
}

func (h *QueryClientHandler) checkParam() {
	if h.Req.GetLimit() == 0 {
		h.Res.BaseResp = util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Req_Illegal, "illegal request")
		return
	}
}
