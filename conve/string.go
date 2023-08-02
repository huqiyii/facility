package conve

import "github.com/cstockton/go-conv"

func StringDefault(i interface{}, v string) string {
	if i == nil {
		return v
	}
	value, err := conv.String(i)
	if err != nil {
		return v
	}
	return value
}
