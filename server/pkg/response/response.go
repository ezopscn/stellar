package response

// 请求结果状态码
const (
	RequestOK                  = 200
	RequestBadRequest          = 400
	RequestForbidden           = 403
	RequestNotFound            = 404
	RequestMethodNotAllowed    = 405
	RequestInternalServerError = 500
	RequestBadGateway          = 502
	RequestServiceUnavailable  = 503
	RequestGatewayTimeout      = 504
	RequestParameterError      = 1000
	RequestUnauthorized        = 1001
	RequestUnactived           = 1002
)

// 请求结果状态码对应的信息描述
const (
	RequestOKMessage                  = "请求成功"
	RequestBadRequestMessage          = "请求失败"
	RequestForbiddenMessage           = "无权限访问该资源"
	RequestNotFoundMessage            = "未找到相关资源"
	RequestMethodNotAllowedMessage    = "请求方法不允许"
	RequestInternalServerErrorMessage = "服务器错误"
	RequestBadGatewayMessage          = "网关错误"
	RequestServiceUnavailableMessage  = "服务不可达"
	RequestGatewayTimeoutMessage      = "网关超时"
	RequestParameterErrorMessage      = "参数错误"
	RequestUnauthorizedMessage        = "用户未登录"
	RequestUnactivedMessage           = "用户未激活"
)

// 请求结果状态码和描述信息做绑定
var ResponseMessage = map[int]string{
	RequestOK:                  RequestOKMessage,
	RequestBadRequest:          RequestBadRequestMessage,
	RequestForbidden:           RequestForbiddenMessage,
	RequestNotFound:            RequestNotFoundMessage,
	RequestMethodNotAllowed:    RequestMethodNotAllowedMessage,
	RequestInternalServerError: RequestInternalServerErrorMessage,
	RequestBadGateway:          RequestBadGatewayMessage,
	RequestServiceUnavailable:  RequestServiceUnavailableMessage,
	RequestGatewayTimeout:      RequestGatewayTimeoutMessage,
	RequestParameterError:      RequestParameterErrorMessage,
	RequestUnauthorized:        RequestUnauthorizedMessage,
	RequestUnactived:           RequestUnactivedMessage,
}

// 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 没有数据的时候响应空的数据
var empty = map[string]interface{}{}

// 响应方法
func Result(code int, message string, data interface{}) {
	panic(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// 成功响应
func Success() {
	Result(RequestOK, ResponseMessage[RequestOK], empty)
}

// 成功响应，带数据
func SuccessWithData(data interface{}) {
	Result(RequestOK, ResponseMessage[RequestOK], data)
}

// 失败响应
func Failed() {
	Result(RequestBadRequest, ResponseMessage[RequestBadRequest], empty)
}

// 失败响应，带状态码
func FailedWithCode(code int) {
	Result(code, ResponseMessage[code], empty)
}

// 失败响应，带消息
func FailedWithMessage(message string) {
	Result(RequestBadRequest, message, empty)
}

// 失败响应，带状态码和消息
func FailedWithCodeAndMessage(code int, message string) {
	Result(code, message, empty)
}
