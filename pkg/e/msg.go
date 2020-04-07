package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	INVALID_PARAM: "invalid param",

	ERROR_AUTH:       "token验证失败",
	ERROR_AUTH_TOKEN: "token生成失败",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "token超时",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "验证token失败",
	ERROR_EXIST_USER:                "已经存在该用户",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "无效到图片格式",
	ERROR_UPLOAD_PREPARE_IMAGE_FAIL: "准备创建图片失败",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
}

func GetMsg(code int) string {
	if msg, exist := MsgFlags[code]; exist {
		return msg
	}
	return MsgFlags[ERROR]
}
