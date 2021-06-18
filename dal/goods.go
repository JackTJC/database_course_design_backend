package dal

import (
	"github.com/JackTJC/backend/dal/db"
	"github.com/JackTJC/backend/model"
	"github.com/jinzhu/gorm"
)

func CreateGoods(goods *model.Goods) error {
	conn, err := db.NewGoodsConn()
	if err!=nil{
		return err
	}
	return conn.CreateGoods(goods)
}

func QueryGoods(goodsType *int, name, desc *string, priceLow, priceHigh,offset,limit *int32) ([]*model.Goods, error) {
	conn, err := db.NewGoodsConn()
	if err!=nil{
		return nil,err
	}
	return conn.QueryGoods(goodsType,name,desc,priceLow,priceHigh,offset,limit)
}

func MCreateGoods(goodsList []*model.Goods) error {
	conn,err:=db.NewGoodsConn()
	if err!=nil{
		return err
	}
	err=conn.Transaction(func(db *gorm.DB) error {
		for _, goods := range goodsList {
			if conn.CreateGoods(goods)!=nil{
				return err
			}
		}
		return nil
	})
	if err!=nil{
		return err
	}
	return nil
}

func MGetGoods(IDList []int32) (map[int32]*model.Goods, error) {
	conn,err:=db.NewGoodsConn()
	if err!=nil{
		return nil,err
	}
	ret:=make(map[int32]*model.Goods)
	goodList,err:=conn.MGetGoodsById(IDList)
	if err!=nil{
		return nil,err
	}
	for _, goods := range goodList {
		ret[goods.Id]=goods
	}
	return ret,nil
}


