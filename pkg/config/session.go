package config

func init() {
	Add("session", StrMap{

		// 目前只支持 Cookie
		"default": Env("SESSION_DRIVER", "cookie"),

		// 会话的 Cookie 名称
		"session_name": Env("SESSION_NAME", "goblog-session"),
	})
}
