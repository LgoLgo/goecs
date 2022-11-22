package forms

type SendSmsForm struct {
	Mobile string `form:"mobile,required" json:"mobile,required" vd:"mobile($)"`
	Type   uint   `form:"type,required" json:"type,required" vd:"($)==1||($)==2;msg:'type must be 1 or 2'"`
}
