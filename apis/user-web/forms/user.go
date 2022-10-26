package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile,required" json:"mobile,required" vd:"mobile($)"`
	PassWord  string `form:"password,required" json:"password,required" vd:"len($)>3 && len($)<20; msg:'password length should be 4 - 19'"`
	Captcha   string `form:"captcha,required" json:"captcha,required"`
	CaptchaId string `form:"captcha_id,required" json:"captcha_id,required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile,required" json:"mobile,required" vd:"mobile($)"`
	PassWord string `form:"password,required" json:"password,required" vd:"len($)>3 && len($)<20; msg:'password length should be 4 - 19'"`
	Code     string `form:"code,required" json:"code,required" vd:"len($)==6; msg:'code length should be 6'"`
}

type UpdateUserForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Gender   string `form:"gender" json:"gender" binding:"required,oneof=female male"`
	Birthday string `form:"birthday" json:"birthday" binding:"required,datetime=2006-01-02"`
}
