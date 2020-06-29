package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
	//	微信Accesstoken 1个小时时效
	WeChatAccessToken = "wechat:token"

//	首页数据缓存 1个小时
WXHOME="wechat:home"
)

// VideoViewKey 视频点击数的key
// view:course:1 -> 100
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:course:%s", strconv.Itoa(int(id)))
}

// VideoViewKey 视频点击数的key
// view:article:1 -> 100
func ArticleViewKey(id uint) string {
	return fmt.Sprintf("view:article:%s", strconv.Itoa(int(id)))
}

// VideoViewKey 视频点击数的key
// view:studio:1 -> 100
func StudioViewKey(id uint) string {
	return fmt.Sprintf("view:studio:%s", strconv.Itoa(int(id)))
}

// VideoViewKey 视频点击数的key
// view:studio:1 -> 100
func BannerViewKey(id uint) string {
	return fmt.Sprintf("view:banner:%s", strconv.Itoa(int(id)))
}

// VideoViewKey 视频点击数的key
// view:studio:1 -> 100
func StudioCourseViewKey(id uint) string {
	return fmt.Sprintf("view:studioCourse:%s", strconv.Itoa(int(id)))
}
