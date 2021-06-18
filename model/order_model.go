package model

func (Order) TableName() string {
	return "order"
}

type Order struct {
	Id       int32 `gorm:"column:id" json:"id"`
	GoodsId  int32 `gorm:"column:goods_id" json:"goods_id"`
	ClientId int32 `gorm:"column:client_id" json:"client_id"`
	Num      int32 `gorm:"column:num" json:"num"`
	Status   int16 `gorm:"column:status" json:"status"`
}

func NewOrder() *Order {
	return &Order{}
}
