package conve

import "github.com/cstockton/go-conv"

func StringDefault(i interface{}, v string) string {
	value, err := conv.String(i)
	if err != nil {
		return v
	}
	return value
}
