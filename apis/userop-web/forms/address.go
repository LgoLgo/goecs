package forms

type AddressForm struct {
	Province     string `form:"province,required" json:"province,required"`
	City         string `form:"city,required" json:"city,required"`
	District     string `form:"district,required" json:"district,required"`
	Address      string `form:"address,required" json:"address,required"`
	SignerName   string `form:"signer_name,required" json:"signer_name,required"`
	SignerMobile string `form:"signer_mobile,required" json:"signer_mobile,required"`
}
