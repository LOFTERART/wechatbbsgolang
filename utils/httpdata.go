package utils

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type HTTPData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnHTTPSuccess(this *beego.Controller, val interface{}) {

	rtndata := HTTPData{
		Code:    200,
		Message: "",
		Data:    val,
	}

	data, err := json.Marshal(rtndata)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
}

func GetHTTPRtnJsonData(errno int, errmsg string) interface{} {

	rtndata := HTTPData{
		Code:    errno,
		Message: errmsg,
		Data:    nil,
	}
	data, _ := json.Marshal(rtndata)

	return json.RawMessage(string(data))

}
