package cache

import (
	"context"
	"encoding/json"
	"github.com/JackTJC/backend/model"
)

const clientKeyFmt = "client:%v"

func SetClient(ctx context.Context, client *model.Client) error {
	data, _ := json.Marshal(client)
	cmd := cacheConn.Set(ctx, constructKey(clientKeyFmt, int(client.Id)), string(data), 0)
	return cmd.Err()
}

func GetClientByID(ctx context.Context,id int)(*model.Client,error){
	cmd := cacheConn.Get(ctx, constructKey(clientKeyFmt, id))
	if cmd.Err()!=nil{
		return nil,cmd.Err()
	}
	if cmd.Val()=="nil"{
		return nil,nil
	}
	retModel:=&model.Client{}
	err := json.Unmarshal([]byte(cmd.Val()), retModel)
	if err!=nil{
		return nil, err
	}
	return retModel,nil
}