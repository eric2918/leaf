package util

import (
	"fmt"
	"strings"
)

func Addr(addr string) string {
	arr := strings.Split(addr, ":")
	if len(arr) > 1 {
		addr = fmt.Sprintf(":%s", arr[1])
	}
	return addr
}
