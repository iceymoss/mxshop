package forms

type MessageForm struct {
	MessageType int32  `form:"type" json:"type" binding:"required,oneof=1 2 3 4 5"`
	Subject     string `form:"subject" json:"subject" binding:"required"`
	Message     string `form:"message" json:"message" binding:"required"`
	File        string `form:"file" json:"file" binding:"required"`
}
