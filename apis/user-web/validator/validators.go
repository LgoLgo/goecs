package validator

import (
	"fmt"
	"regexp"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
)

func ValidateMobile() {
	binding.MustRegValidateFunc("mobile", func(args ...interface{}) error {
		ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, args[0].(string))
		if !ok {
			return fmt.Errorf("invalid mobile number")
		}
		return nil
	})
}
