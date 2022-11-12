package forms

type CategoryForm struct {
	Name           string `form:"name,required" json:"name,required"`
	ParentCategory int32  `form:"parent" json:"parent"`
	Level          int32  `form:"level,required" json:"level,required"`
	IsTab          *bool  `form:"is_tab,required" json:"is_tab,required"`
}

type UpdateCategoryForm struct {
	Name  string `form:"name,required" json:"name,required"`
	IsTab *bool  `form:"is_tab" json:"is_tab"`
}
