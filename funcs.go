package uuid

import (
	"errors"
	"strconv"
	"strings"
)

func Pack(id uint64, base int) string {
	arr := make([]string, 2)
	arr[1] = strconv.FormatUint(uint64(id), base)
	arr[0] = strconv.FormatUint(uint64(len(arr[1])), base)
	return strings.Join(arr, "")
}

func Split(s string, base int) (prefix, suffix string, err error) {
	var i int
	if i, err = Index(s, base); err != nil {
		return
	}
	prefix = s[1:i]
	suffix = s[i:]
	return
}

// Index 获取有效字符串长度
func Index(id string, base int) (r int, err error) {
	var v int64
	if v, err = strconv.ParseInt(id[0:1], base, 64); err != nil {
		return
	} else {
		r = int(v) + 1
	}
	if r > len(id) {
		err = errors.New("oid error")
	}
	return
}
