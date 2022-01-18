package herrs

// 全局错误码定义
// 999**
var (
	SUCCESS = New("00000", "OK")
	FAIL    = New("ERROR", "FAIL")
)
