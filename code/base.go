package code

import (
	"github.com/aluka-7/metacode"
)

var (
	cm map[int]string
)

func init() {
	Register(map[metacode.Code]string{
		metacode.Success:            "Success",
		metacode.NotModified:        "Not Modified",
		metacode.TemporaryRedirect:  "Redirect",
		metacode.RequestErr:         "Request Error",
		metacode.Unauthorized:       "Unauthorized",
		metacode.AccessDenied:       "Access Denied",
		metacode.NothingFound:       "Nothing Found",
		metacode.MethodNotAllowed:   "不支持该方法",
		metacode.Conflict:           "冲突",
		metacode.Canceled:           "客户端取消请求",
		metacode.ServerErr:          "服务器错误",
		metacode.ServiceUnavailable: "过载保护,服务暂不可用",
		metacode.Deadline:           "服务调用超时",
		metacode.LimitExceed:        "超出限制",
		metacode.ValidateErr:        "服务器请求参数校验出错",
	})
}

func Register(m map[metacode.Code]string) {
	if cm == nil {
		cm = make(map[int]string, len(m))
	}

	for k, v := range m {
		cm[k.Code()] = v
	}
	metacode.Register(cm)
}
