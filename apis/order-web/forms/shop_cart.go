package forms

type ShopCartItemForm struct {
	GoodsId int32 `json:"goods,required"`
	Nums    int32 `json:"nums,required"`
}

type ShopCartItemUpdateForm struct {
	Nums    int32 `json:"nums,required"`
	Checked *bool `json:"checked"`
}
