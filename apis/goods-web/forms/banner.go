package forms

type BannerForm struct {
	Image string `form:"image" json:"image"`
	Index int    `form:"index,required" json:"index,required"`
	Url   string `form:"url" json:"url"`
}
