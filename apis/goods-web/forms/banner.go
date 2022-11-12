package forms

type BannerForm struct {
	Image string `form:"image" json:"image" binding:"url"`
	Index int    `form:"index" json:"index" binding:"required"`
	Url   string `form:"url" json:"url" binding:"url"`
}
