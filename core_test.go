package core

import (
	"fmt"
	"github.com/aluka-7/core/code"
	"github.com/aluka-7/metacode"
	"testing"
)

func TestCore(t *testing.T) {
	code.Register("zh", map[metacode.Code]string{
		metacode.OK:                "好",
		metacode.Success:           "成功",
		metacode.NotModified:       "木有改动",
		metacode.TemporaryRedirect: "撞车跳转",
		metacode.RequestErr:        "请求错误",
	})

	code.Register("en", map[metacode.Code]string{
		metacode.OK:                "OK",
		metacode.Success:           "Success",
		metacode.NotModified:       "Not Modified",
		metacode.TemporaryRedirect: "Temporary Redirect",
		metacode.RequestErr:        "Request Error",
	})

	code.Register("ar", map[metacode.Code]string{
		metacode.OK:                "OK!!!!!!!!!!!!!!!!",
		metacode.Success:           "Success!!!!!!!!!!!!!!!!",
		metacode.NotModified:       "Not Modified!!!!!!!!!!!!!!!!",
		metacode.TemporaryRedirect: "Temporary Redirect!!!!!!!!!!!!!!!!",
		metacode.RequestErr:        "Request Error!!!!!!!!!!!!!!!!",
	})

	cause := metacode.Cause(metacode.NotModified)
	fmt.Println(cause.Code(), cause.Message("ar"))
}
