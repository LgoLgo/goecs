package forms

type BrandForm struct {
	Name string `form:"name" json:"name"`
	Logo string `form:"logo" json:"logo"`
}

type CategoryBrandForm struct {
	CategoryId int `form:"category_id,required" json:"category_id,required"`
	BrandId    int `form:"brand_id,required" json:"brand_id,required"`
}
