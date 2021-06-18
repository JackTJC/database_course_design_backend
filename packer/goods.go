package packer

import (
	"github.com/JackTJC/backend/model"
	"github.com/JackTJC/backend/pb_gen"
)

func PackGoodsPbToModel(goodsPb *pb_gen.Goods) *model.Goods {
	return &model.Goods{
		Id:          goodsPb.GetGoodsId(),
		Type:        int16(goodsPb.GetGoodsType()),
		Name:        goodsPb.GetGoodsName(),
		Description: goodsPb.GetGoodsDescription(),
		Price:       goodsPb.GetPrice(),
		PictureUrl:  goodsPb.GetPictureUrl(),
		Discount: int16(goodsPb.GetDiscount()),
		Inventory:   goodsPb.GetInventory(),
	}
}

func PackGoodsModelToPb(goodsModel *model.Goods)*pb_gen.Goods {
	return &pb_gen.Goods{
		GoodsId:          goodsModel.Id,
		GoodsType:        pb_gen.GoodsType(goodsModel.Type),
		GoodsName:        goodsModel.Name,
		GoodsDescription: goodsModel.Description,
		Price:            goodsModel.Price,
		PictureUrl:       goodsModel.PictureUrl,
		Discount:         int32(goodsModel.Discount),
		Inventory: int32(goodsModel.Discount),
	}
}