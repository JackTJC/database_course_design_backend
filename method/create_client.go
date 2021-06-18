package method

import (
	"context"
	"github.com/JackTJC/backend/dal"
	"github.com/JackTJC/backend/packer"
	"github.com/JackTJC/backend/pb_gen"
	"github.com/JackTJC/backend/util"
	"log"
)
type CreateClientHandler struct {
	Req *pb_gen.CreateClientRequest
	Res *pb_gen.CreateClientResponse
	Ctx context.Context
}

func NewCreateClientHandler(ctx context.Context,req *pb_gen.CreateClientRequest) *CreateClientHandler {
	return &CreateClientHandler{
		Req: req,
		Res: &pb_gen.CreateClientResponse{
			BaseResp: &pb_gen.BaseResp{
				ErrNo: pb_gen.ErrNo_ErrNo_Success,
			},
		},
		Ctx: ctx,
	}
}
func (h *CreateClientHandler) Run() {
	defer func() {
		log.Printf("request=%v\n", h.Req)
		log.Printf("response=%v\n", h.Res)
	}()
	clientList, err := dal.QueryClient(&h.Req.Client.Tel, nil, nil, nil)
	if len(clientList)>0{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_User_Has_Existed,"user already exist")
		return
	}
	err = dal.CreateClient(packer.PackClientPbToModel(h.Req.GetClient()))
	if err!=nil{
		h.Res.BaseResp=util.BuildBaseResp(pb_gen.ErrNo_ErrNo_Failed,err.Error())
		return
	}
}
