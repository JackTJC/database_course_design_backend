package packer

import (
	"github.com/JackTJC/backend/model"
	"github.com/JackTJC/backend/pb_gen"
)

func PackClientPbToModel(clientPb *pb_gen.Client) *model.Client {
	return &model.Client{
		Id:     clientPb.GetClientId(),
		Name:   clientPb.GetName(),
		Tel:    clientPb.GetTel(),
		Passwd: clientPb.GetPasswd(),
	}
}

func PackClientModelToPb(clientModel *model.Client) *pb_gen.Client {
	return &pb_gen.Client{
		ClientId: clientModel.Id,
		Name:     clientModel.Name,
		Tel:      clientModel.Tel,
		Passwd:   clientModel.Passwd,
	}
}
