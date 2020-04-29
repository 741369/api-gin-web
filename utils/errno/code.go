package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "成功", EnMessage: "Success", ThMessage: "สำเร็จ"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务出错", EnMessage: "Internal server error", ThMessage: "ข้อผิดพลาดของบริการภายใน"}
	ErrParam            = &Errno{Code: 10002, Message: "获取接口参数出错", EnMessage: "Param error, see doc for more info"}
	SignInvalid         = &Errno{Code: 10003, Message: "签名错误", EnMessage: "sign invalid"}
	ErrTokenInvalid     = &Errno{Code: 10004, Message: "校验token失败", EnMessage: "The token was invalid"}
	ErrDatabase         = &Errno{Code: 10005, Message: "数据库内部错误", EnMessage: "Database error"}
	UploadFileErr       = &Errno{Code: 10006, Message: "上传文件失败", EnMessage: "Upload file error"}
	NotWhiteList        = &Errno{Code: 10007, Message: "没有权限", EnMessage: "you do not have permission to request this port"}
	SessionInvalid      = &Errno{Code: 10008, Message: "登录态失效，请重新登录", EnMessage: "session invalid, login again please"}
)
