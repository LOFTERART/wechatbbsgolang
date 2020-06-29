package service

import (
	"QUZHIYOU/utils"
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
)

func getControllerAndAction(rawvalue string) (controller, action string) {
	vals := strings.Split(rawvalue, "/")
	return vals[2], vals[2] + "/" + vals[3]
}

func FilterFunc(ctx *context.Context) {

	controller, action := getControllerAndAction(ctx.Request.RequestURI)

	fmt.Println("-----S----")
	fmt.Println(controller, action)
	fmt.Println("-----E----")

	token := ctx.Input.Header("userId")

	fmt.Println(token)

	if token == "" {
		data := utils.GetHTTPRtnJsonData(401, "请重新登录 缺少token")
		ctx.Output.JSON(data, true, false)
		ctx.Redirect(200, "/")
		return
	}

}
