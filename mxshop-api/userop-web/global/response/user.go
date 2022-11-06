package response

import (
	"fmt"
	"time"
)

//JsonTime 时间格式转换
type JsonTime time.Time

//MarshalJSON 当数据调用c.JSON时，MarshalJSON会自动被调用
func (j JsonTime) MarshalJSON() ([]byte, error) {
	stdtime := fmt.Sprintf("\"%s\"", time.Time(j).Format("2022-01-01"))
	return []byte(stdtime), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	NickName string   `json:"name"`
	BirthDay JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}
