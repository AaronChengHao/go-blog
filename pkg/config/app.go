package config

func init() {
	Add("app", StrMap{

		// 应用名称，暂时没有使用到
		"name": Env("APP_NAME", "GoBlog"),

		// 当前环境，用以区分多环境
		"env": Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": Env("APP_DEBUG", false),

		// 应用服务端口
		"port": Env("APP_PORT", "3000"),

		// gorilla/sessions 在 Cookie 中加密数据时使用
		"key": Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
	})
}
