package forms

type UserFavForm struct {
	GoodsId int32 `form:"goods,required" json:"goods,required" `
}
