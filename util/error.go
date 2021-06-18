package util

import "github.com/JackTJC/backend/pb_gen"

func BuildBaseResp(errCode pb_gen.ErrNo, msg string) *pb_gen.BaseResp {
	return &pb_gen.BaseResp{
		ErrNo: errCode,
		Msg:msg,
	}
}
