package utils

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrCodeMsgMap = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "内部服务错误",
	502: "用户名或者密码错误",
	504: "解析请求失败",
	505: "更新失败",
	506: "注册失败",
	507: "该手机号已注册",
	508: "输入的电话号码和登陆的电话号码不匹配",
	509: "添加失败",
	510: "删除数据失败",
	511: "获取数据失败",
	512: "获取参数失败",
	//继续添加其他的错误
}

func NewError(errCode int) ErrorResponse {
	msg, ok := ErrCodeMsgMap[errCode]
	if !ok {
		msg = "Unknown Error"
	}
	return ErrorResponse{
		Code:    errCode,
		Message: msg,
	}
}
