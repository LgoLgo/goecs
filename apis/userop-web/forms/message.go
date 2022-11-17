package forms

type MessageForm struct {
	MessageType int32  `form:"type,required" json:"type,required"`
	Subject     string `form:"subject,required" json:"subject,required" `
	Message     string `form:"message,required" json:"message,required" `
	File        string `form:"file,required" json:"file,required" `
}
