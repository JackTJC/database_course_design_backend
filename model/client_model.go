package model

func (Client) TableName() string {
	return "client"
}

type Client struct {
	Id     int32  `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Tel    string `gorm:"column:tel" json:"tel"`
	Passwd string `gorm:"column:passwd" json:"passwd"`
}

func NewClient() *Client {
	return &Client{}
}
