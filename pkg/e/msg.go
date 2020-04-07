package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	INVALID_PARAM: "invalid param",

	ERROR_AUTH:       "token验证失败",
	ERROR_AUTH_TOKEN: "token生成失败",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token超时",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "验证token失败",
	ERROR_EXIST_USER:               "已经存在该用户",
}

func GetMsg(code int) string {
	if msg, exist := MsgFlags[code]; exist {
		return msg
	}
	return MsgFlags[ERROR]
}
