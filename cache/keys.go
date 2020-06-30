package cache

import (
	"fmt"
	"strconv"
)

const (
	//	微信Accesstoken 1个小时时效
	WeChatAccessToken = "wechat:token"
	//	首页数据缓存 1个小时
	WXHOME = "wechat:home"
)

// VideoViewKey 视频点击数的key
// view:course:1 -> 100
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:course:%s", strconv.Itoa(int(id)))
}



