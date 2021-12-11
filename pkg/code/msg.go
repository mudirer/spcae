package code

var MsgFlags = map[int]string {
	OK : "ok",
	ERROR : "fail",
	ERROR_NOT_EXIST_USER:   "用户不存在",
	INVALID_PARAMS:         "参数错误",
	ERROR_EXIST_ACCOUNT:    "账户已存在",
	SUCCESS:                "成功",
	PARAMETER_ERROR:        "参数错误",
	ERROR_CHECK_EXIST_FAIL: "查询ID存在错误",
	ERROR_NOT_EXIST:        "查询不存在",
	ERROR_GET_FAIL:         "查询失败",
	ERROR_DELETE_FAIL:      "删除失败",
	ERROR_POST_FAIL:        "添加失败",
	ERROR_PUT_FAIL:         "更新失败",
	ERROR_ID_FAIL:          "ID不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}