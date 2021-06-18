package cache

import (
	"context"
	"encoding/json"
	"github.com/JackTJC/backend/model"
)

const goodsKeyFmt = "goods:%v"

func SetGoods(ctx context.Context, goods *model.Goods) error {
	data, _ := json.Marshal(goods)
	cmd := cacheConn.Set(ctx, constructKey(goodsKeyFmt, int(goods.Id)), string(data), 0)
	return cmd.Err()
}

func GetGoodsByID(ctx context.Context,id int)(*model.Goods,error){
	cmd := cacheConn.Get(ctx, constructKey(goodsKeyFmt, id))
	if cmd.Err()!=nil{
		return nil,cmd.Err()
	}
	if cmd.Val()=="nil"{
		return nil,nil
	}
	retModel:=&model.Goods{}
	err := json.Unmarshal([]byte(cmd.Val()), retModel)
	if err!=nil{
		return nil, err
	}
	return retModel,nil
}
