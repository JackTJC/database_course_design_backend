package db

import (
	"github.com/JackTJC/backend/model"
	"github.com/jinzhu/gorm"
)

type orderConn struct {
	*gorm.DB
}

func NewOrderConn() (*orderConn, error) {
	return &orderConn{dbConn.Model(model.NewOrder())},nil
}

func (d *orderConn)CreateOrder(order *model.Order)error{
	return d.Create(order).Error
}

func (d *orderConn)QueryOrder(orderId,goodsId,clientId,offset,limit *int32)([]*model.Order,error){
	conn:=d.DB
	if orderId!=nil{
		conn=conn.Where("id = ?",*orderId)
	}
	if goodsId!=nil{
		conn=conn.Where("goods_id = ?",*goodsId)
	}
	if clientId!=nil{
		conn.Where("client_id = ?",*clientId)
	}
	if offset!=nil&&limit!=nil{
		conn.Offset(*offset).Limit(*limit)
	}
	orderList:=make([]*model.Order,0)
	if err:=conn.Find(&orderList).Error;err!=nil{
		return nil,err
	}
	return orderList,nil
}

func (d *orderConn) UpdateStatus(id int32, stat int32) error {
	updateMap:=map[string]interface{}{
		"status":stat,
	}
	return d.Where("id = ?",id).Update(updateMap).Error
}
