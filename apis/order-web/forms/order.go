package forms

type CreateOrderForm struct {
	Address string `json:"address,required"`
	Name    string `json:"name,required"`
	Mobile  string `json:"mobile,required"`
	Post    string `json:"post,required"`
}
