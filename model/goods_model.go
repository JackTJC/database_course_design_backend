package model

func (Goods) TableName() string {
	return "goods"
}

type Goods struct {
	Id          int32   `gorm:"column:id" json:"id"`
	Type        int16   `gorm:"column:type" json:"type"`
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Price       float32 `gorm:"column:price" json:"price"`
	PictureUrl  string  `gorm:"column:picture_url" json:"picture_url"`
	Discount    int16   `gorm:"column:discount" json:"discount"`   // 折扣，代表几折
	Inventory   int32   `gorm:"column:inventory" json:"inventory"` // 库存
}

func NewGoods() *Goods {
	return &Goods{}
}
