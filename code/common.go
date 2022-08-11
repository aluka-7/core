package code

import "github.com/aluka-7/metacode"

// common

var (
	RecordExistError = metacode.NewCode(1001)
)

func init() {
	Register(map[metacode.Code]string{
		RecordExistError: "该记录已存在",
	})
}
