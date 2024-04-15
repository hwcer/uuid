package uuid

import (
	"strconv"
	"strings"
)

type UUID uint64

func (u UUID) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u UUID) Parse() (sid int32, err error) {
	var v uint64
	var p string
	if p, _, err = Split(u.String(), 10); err != nil {
		return
	} else if v, err = strconv.ParseUint(p, 10, 64); err != nil {
		return
	} else {
		sid = int32(v)
	}
	return
}

// ObjectID 使用uuid和物品ID生成全服唯一并且uuid下唯一(每个用户的每个IID唯一)的ObjectID
func (u UUID) ObjectID(iid int32) (oid ObjectID, err error) {
	var p, s string
	var v uint64
	p, s, err = Split(u.String(), 10)
	if err != nil {
		return
	}
	var build strings.Builder
	if v, err = strconv.ParseUint(p, 10, 64); err != nil {
		return
	} else {
		build.WriteString(Pack(v, BitSize))
	}
	build.WriteString(Pack(uint64(iid), BitSize))
	if v, err = strconv.ParseUint(s, 10, 64); err != nil {
		return
	} else {
		build.WriteString(Pack(v, BitSize))
	}
	oid = ObjectID(build.String())
	return
}
