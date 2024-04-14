package uuid

import (
	"strconv"
)

const BitSize = 32

type ObjectID string

// Parse 通过道具ID解析出服务器ID和配置ID
func (id ObjectID) Parse() (sid, iid int32, err error) {
	var v uint64
	var p string
	var s string
	if p, s, err = Split(string(id), BitSize); err != nil {
		return
	} else if v, err = strconv.ParseUint(p, BitSize, 64); err != nil {
		return
	} else {
		sid = int32(v)
	}

	if p, _, err = Split(s, BitSize); err != nil {
		return
	} else if v, err = strconv.ParseUint(p, BitSize, 64); err == nil {
		iid = int32(v)
	}
	return
}
