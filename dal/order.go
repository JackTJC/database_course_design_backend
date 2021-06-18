package dal

import (
	"github.com/JackTJC/backend/dal/db"
	"github.com/JackTJC/backend/model"
)

func CreateOrder(order *model.Order) error {
	conn, err := db.NewOrderConn()
	if err!=nil{
		return err
	}
	return conn.CreateOrder(order)
}

func QueryOrder(orderId, goodsId, clientId,offset,limit *int32) ([]*model.Order, error) {
	conn, err := db.NewOrderConn()
	if err != nil {
		return nil,err
	}
	return conn.QueryOrder(orderId,goodsId,clientId,offset,limit)
}

func UpdateStat(id int32, status int32) error {
	conn,err:=db.NewOrderConn()
	if err!=nil{
		return err
	}
	return conn.UpdateStatus(id,status)
}
