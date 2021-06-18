package db

import (
	"github.com/JackTJC/backend/model"
	"github.com/jinzhu/gorm"
)

type ClientDB struct {
	*gorm.DB
}

func NewClientConn()(*ClientDB,error){
	return &ClientDB{dbConn.Model(model.NewClient())},nil
}

func (d *ClientDB)CreateClient(client *model.Client) error {
	return d.Create(client).Error
}

func (d *ClientDB)QueryClient(tel *string,id ,offset,limit *int32)([]*model.Client,error){
	clientList:=make([]*model.Client,0)
	conn:=d.DB
	if tel!=nil{
		conn=d.Where("tel = ?",*tel)
	}
	if id!=nil{
		conn=conn.Where("id = ?",*id)
	}
	if offset!=nil&&limit!=nil{
		conn=conn.Offset(*offset).Limit(*limit)
	}
	err:=conn.Find(&clientList).Error
	return clientList,err

}

func (d *ClientDB) QueryClientID(tel *string,id ,offset,limit *int32)([]int32,error) {
	clientIDList:=make([]int32,0)
	conn := d.DB
	conn = conn.Select("id")
	if tel!=nil{
		conn=conn.Where("tel = ?",*tel)
	}
	if id!=nil{
		conn=conn.Where("id = ?",*id)
	}
	if offset!=nil&&limit!=nil{
		conn=conn.Offset(*offset).Limit(*limit)
	}
	err:=conn.Find(&clientIDList).Error
	return clientIDList,err
}

func (d *ClientDB) GetClientByTel(tel string) (*model.Client, error) {
	client:=model.NewClient()
	if err:= d.Where("tel = ?",tel).Find(client).Error;err!=nil{
		return nil,err
	}
	return client,nil
}

func (d *ClientDB)GetClientByIds(idList []int32)([]*model.Client,error){
	clientList:=make([]*model.Client,0)
	if err:= d.Where("id in ?",idList).Find(&clientList).Error;err!=nil{
		return nil,err
	}
	return clientList,nil
}
