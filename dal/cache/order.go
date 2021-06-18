package cache

import (
	"context"
	"encoding/json"
	"github.com/JackTJC/backend/model"
)

const orderKeyFmt = "order:%v"


func SetOrder(ctx context.Context, order *model.Order) error {
	data, _ := json.Marshal(order)
	cmd := cacheConn.Set(ctx, constructKey(orderKeyFmt, int(order.Id)), string(data), 0)
	return cmd.Err()
}

func GetOrderByID(ctx context.Context,id int)(*model.Order,error){
	cmd := cacheConn.Get(ctx, constructKey(orderKeyFmt, id))
	if cmd.Err()!=nil{
		return nil,cmd.Err()
	}
	if cmd.Val()=="nil"{
		return nil,nil
	}
	retModel:=&model.Order{}
	err := json.Unmarshal([]byte(cmd.Val()), retModel)
	if err!=nil{
		return nil, err
	}
	return retModel,nil
}
