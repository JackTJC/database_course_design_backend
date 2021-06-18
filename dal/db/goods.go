package db

import (
	"fmt"
	"github.com/JackTJC/backend/model"
	"github.com/jinzhu/gorm"
)

type goodsConn struct {
	*gorm.DB
}

func NewGoodsConn()(*goodsConn,error){
	return &goodsConn{dbConn.Model(model.NewGoods())},nil
}

func (d *goodsConn)CreateGoods(goods *model.Goods) error {
	return d.Create(goods).Error
}

func (d *goodsConn)GetGoodsById(id int) (*model.Goods, error) {
	goods:=model.NewGoods()
	err:= d.Where("id = ?",id).Find(goods).Error
	if err!=nil{
		return nil, err
	}
	return goods,nil
}

func (d *goodsConn) MGetGoodsById(IDList []int32) ([]*model.Goods,error) {
	goodsList:=make([]*model.Goods,0)
	err:=d.Where("id in ?",IDList).Find(&goodsList).Error
	if err!=nil{
		return nil, err
	}
	return goodsList,nil
}

func (d *goodsConn)QueryGoods(goodsType *int,name,desc *string,priceLow,priceHigh,offset,limit *int32) ([]*model.Goods,error) {
	conn:= d.DB
	if goodsType!=nil{
		conn=conn.Where("type = ?",*goodsType)
	}
	if name != nil{
		conn=conn.Where(fmt.Sprintf("name like '%%%s%%'",*name)) // ignore_security_alert
	}
	if desc!=nil{
		conn=conn.Where(fmt.Sprintf("description like '%%%s%%'",*desc)) // ignore_security_alert
	}
	if priceLow!=nil{
		conn=conn.Where("price >= ?",priceLow)
	}
	if priceHigh!=nil{
		conn=conn.Where("price < ?",*priceHigh)
	}
	if limit!=nil&&offset!=nil{
		conn=conn.Offset(*offset).Limit(*limit)
	}
	res:=make([]*model.Goods,0)
	if err:=conn.Find(&res).Error;err!=nil{
		return nil,err
	}
	return res,nil
}