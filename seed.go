package uuid

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type Seed struct {
	shard        uint16
	index        uint32 //UID递增ID
	prefix       string //len(shard) + shard
	objectSeed   uint64
	objectPrefix string
	objectSuffix string
}

// New 创建种子
//
//	shard 服务器分片ID
//	index 自增种子,如果不使用UUID可以为0
func New(shard uint16, index uint32) *Seed {
	if index >= math.MaxUint32 {
		panic("uuid index out of range")
	}
	u := &Seed{}
	u.shard = shard
	u.index = index
	u.prefix = Pack(uint64(shard), 10)

	u.objectPrefix = Pack(uint64(shard), BitSize)
	t, _ := time.Parse("2006-01-02 15:04:05-0700", "2024-04-11 12:00:00+0800")
	v := time.Now().Unix() - t.Unix()
	u.objectSuffix = Pack(uint64(v), BitSize)

	return u
}

func Create(uuid UUID) (*Seed, error) {
	prefix, suffix, err := Split(uuid.String(), 10)
	if err != nil {
		return nil, err
	}
	var v uint64
	var shard uint16
	var index uint32

	if v, err = strconv.ParseUint(prefix, 10, 64); err == nil {
		shard = uint16(v)
	} else {
		return nil, err
	}

	if v, err = strconv.ParseUint(suffix, 10, 64); err == nil {
		index = uint32(v)
	} else {
		return nil, err
	}

	return New(shard, index), nil
}

func (u *Seed) Shard() uint16 {
	return u.shard
}

func (u *Seed) Index() uint32 {
	return u.index
}

// New 生成UUID
func (u *Seed) New() (UUID, error) {
	if u.shard == 0 {
		return 0, errors.New("shard not init")
	}
	if u.index == 0 {
		return 0, errors.New("index not init")
	}
	i := atomic.AddUint32(&u.index, 1)
	var build strings.Builder
	build.WriteString(u.prefix)
	build.WriteString(fmt.Sprintf("%v", i))
	if v, err := strconv.ParseUint(build.String(), 10, 64); err == nil {
		return UUID(v), nil
	} else {
		return 0, err
	}
}

// ObjectID 完全随机的全服唯一字符串ID
func (u *Seed) ObjectID(iid int32) ObjectID {
	i := atomic.AddUint64(&u.objectSeed, 1)
	var build strings.Builder
	build.WriteString(u.objectPrefix)
	build.WriteString(Pack(uint64(iid), BitSize))
	build.WriteString(u.objectSuffix)
	build.WriteString(Pack(i, BitSize))
	return ObjectID(build.String())
}
