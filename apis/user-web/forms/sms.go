package forms

type SendSmsForm struct {
	Mobile string `form:"mobile,required" json:"mobile,required" vd:"mobile($)"` //手机号码格式有规范可寻， 自定义validator
	Type   uint   `form:"type,required" json:"type,required" vd:"($)==1||($)==2;msg:'type must be 1 or 2'"`
}
