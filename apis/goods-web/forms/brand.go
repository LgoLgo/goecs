package forms

type BrandForm struct {
	Name string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Logo string `form:"logo" json:"logo" binding:"url"`
}

type CategoryBrandForm struct {
	CategoryId int `form:"category_id" json:"category_id" binding:"required"`
	BrandId    int `form:"brand_id" json:"brand_id" binding:"required"`
}
