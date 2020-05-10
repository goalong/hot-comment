package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400
	INVALID_PAGE = 420
)

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	INVALID_PAGE: "页码或数量参数不合法",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}