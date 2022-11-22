package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

// MarshalJSON implements the Marshaller interface.
func (j JsonTime) MarshalJSON() ([]byte, error) {
	// custom format
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	NickName string   `json:"name"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}
