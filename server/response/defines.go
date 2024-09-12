package response

const (
	CodeSuccess      = "00000"
	CodeNotFound     = "B0001"
	CodeParamsError  = "B0002"
	CodeNotLogin     = "B1001"
	CodeNoPermission = "B1002"
	CodeBanned       = "B1003"
	CodeServerError  = "A0001"

	MessageSuccess      = "success"
	MessageNotFound     = "not found"
	MessageParamsError  = "params error"
	MessageNotLogin     = "not login"
	MessageNoPermission = "no permission"
	MessageServerError  = "server error"
	MessageBanned       = "your account has been banned"

	MessageUnknown = "unknown error"
)

var msgMap = map[string]string{
	CodeSuccess:      MessageSuccess,
	CodeNotFound:     MessageNotFound,
	CodeParamsError:  MessageParamsError,
	CodeServerError:  MessageServerError,
	CodeNotLogin:     MessageNotLogin,
	CodeNoPermission: MessageNoPermission,
	CodeBanned:       MessageBanned,
}

func message(code string) string {
	if msg, ok := msgMap[code]; ok {
		return msg
	}
	return MessageUnknown
}
