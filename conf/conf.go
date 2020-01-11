package conf

// Init 初始化配置项
func init() {
	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}
}
