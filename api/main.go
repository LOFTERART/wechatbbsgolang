package api

import (
	"QUZHIYOU/conf"
	"QUZHIYOU/serializer"
	"encoding/json"
	"fmt"
	validator "gopkg.in/go-playground/validator.v9"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {

	fmt.Println(err, "999999")
	if ve, ok := err.(validator.ValidationErrors); ok {
		fmt.Println(ve, "------eeeee-----")

		for _, e := range ve {

			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))

			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))

			return serializer.Response{
				Code:  40002,
				Msg:   fmt.Sprintf("%s%s", field, tag),
				Error: fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code:  40001,
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Code:  40001,
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}
}
