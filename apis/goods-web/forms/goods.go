package forms

type GoodsForm struct {
	Name        string   `form:"name,required" json:"name,required" vd:"len($)>2 && len($)<100"`
	GoodsSn     string   `form:"goods_sn,required" json:"goods_sn,required"`
	Stocks      int32    `form:"stocks,required" json:"stocks,required"`
	CategoryId  int32    `form:"category,required" json:"category,required"`
	MarketPrice float32  `form:"market_price,required" json:"market_price,required"`
	ShopPrice   float32  `form:"shop_price,required" json:"shop_price,required"`
	GoodsBrief  string   `form:"goods_brief,required" json:"goods_brief,required"`
	Images      []string `form:"images,required" json:"images,required"`
	DescImages  []string `form:"desc_images,required" json:"desc_images,required"`
	ShipFree    *bool    `form:"ship_free,required" json:"ship_free,required"`
	FrontImage  string   `form:"front_image,required" json:"front_image,required"`
	Brand       int32    `form:"brand,required" json:"brand,required"`
}

type GoodsStatusForm struct {
	IsNew  *bool `form:"new,required" json:"new,required"`
	IsHot  *bool `form:"hot,required" json:"hot,required"`
	OnSale *bool `form:"sale,required" json:"sale,required"`
}
