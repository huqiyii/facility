package conve

import "github.com/cstockton/go-conv"

func Int64Default(i interface{}, v int64) int64 {
	value, err := conv.Int64(i)
	if err != nil {
		return v
	}
	return value
}
