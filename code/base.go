package code

import (
	"github.com/aluka-7/metacode"
)

func Register(l string, m map[metacode.Code]string) {
	cm := make(map[int]string, len(m))
	for k, v := range m {
		cm[k.Code()] = v
	}
	metacode.Register(l, cm)
}
