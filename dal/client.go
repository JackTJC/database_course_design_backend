package dal

import (
	"context"
	"github.com/JackTJC/backend/dal/cache"
	"github.com/JackTJC/backend/dal/db"
	"github.com/JackTJC/backend/model"
)

func CreateClient(client *model.Client) error {
	conn,err:=db.NewClientConn()
	if err!=nil{
		return err
	}
	return conn.CreateClient(client)
}

func QueryClient(tel *string, id ,offset,limit *int32) ([]*model.Client,error) {
	conn, err := db.NewClientConn()
	if err!=nil{
		return nil,err
	}
	return conn.QueryClient(tel,id,offset,limit)
}

func QueryClientWithCache(ctx context.Context, tel *string, id, offset, limit *int32) ([]*model.Client,error){
	conn, err := db.NewClientConn()
	if err!=nil{
		return nil, err
	}
	idList,err:=conn.QueryClientID(tel,id,offset,limit)
	hits:=make([]*model.Client,0)
	missIDs:=make([]int32,0)
	for _, id := range idList {
		client, err := cache.GetClientByID(ctx, int(id))
		if err!=nil{
			return nil,err
		}
		if client==nil{
			missIDs = append(missIDs, id)
			continue
		}
		hits = append(hits, client)
	}
	miss,err:=conn.GetClientByIds(missIDs)
	if err!=nil{
		return nil,err
	}
	hits = append(hits, miss...)
	if err:=mSetClientInCache(ctx,miss);err!=nil{
		return nil,err
	}
	return hits,nil
}

func mSetClientInCache(ctx context.Context, clientList []*model.Client) error {
	for _, client := range clientList {
		if err := cache.SetClient(ctx, client);err!=nil{
			return err
		}
	}
	return nil
}