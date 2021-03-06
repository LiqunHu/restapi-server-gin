package e

var MsgFlags = map[string]string{
	"SUCCESS": "方法不支持",
	"ERROR":   "未指定错误",
	"INPUT":   "请求接口错误",
	"auth_01": "用户不存在或者密码错误",
}

// GetMsg get error information based on Code
func GetMsg(code string) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags["ERROR"]
}
