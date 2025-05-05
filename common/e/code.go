package e

const (
	SUCCESS         = 1
	UNKNOW_IDENTITY = 403
	MysqlERR        = 1001
)

var ErrMsg = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "内部错误",
	UNKNOW_IDENTITY: "未知身份",
}
