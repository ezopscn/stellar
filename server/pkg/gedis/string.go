package gedis

// String 类型结果封装
type StringResult struct {
	Result string
	Error  error
}

// 构造函数
func NewStringResult(result string, error error) *StringResult {
	return &StringResult{Result: result, Error: error}
}

// 解析结果
func (r *StringResult) Unwrap() string {
	// if r.Error != nil {
	// 	common.SystemLog.Debug("Redis Cache Query Error: ", r.Error.Error())
	// }
	return r.Result
}

// 查询失败返回默认值
func (r *StringResult) UnwrapWithDefaultValue(v string) string {
	if r.Error != nil {
		// common.SystemLog.Debug("Redis Cache Query Error, Return Default Value: ", r.Error.Error())
		return v
	}
	return r.Result
}

// 查询失败执行函数
func (r *StringResult) UnwrapWithFunc(f func() string) string {
	if r.Error != nil {
		// common.SystemLog.Debug("Redis Cache Query Error, Return Function Value: ", r.Error.Error())
		return f()
	}
	return r.Result
}
