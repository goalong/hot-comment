package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400
	INVALID_PAGE = 420
	INVALID_ES_CLIENT = 430
)

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	INVALID_PAGE: "页码或数量参数不合法(范围应该在1~50之间)",
	INVALID_ES_CLIENT: "ES连接异常",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}