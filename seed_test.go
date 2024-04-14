package uuid

import (
	"testing"
)

func TestUUID(t *testing.T) {
	seed := New(10, 1000)

	uid, _ := seed.New()
	t.Logf("uid:%v\n", uid)
	sid, _ := uid.Parse()
	t.Logf("sid:%v \n", sid)

	seed, _ = Create(uid)
	iid := int32(10021)

	oid := seed.ObjectID(iid)
	t.Logf("oid:%v\n", oid)

	sid, iid, _ = oid.Parse()
	t.Logf("sid:%v ,iid:%v \n", sid, iid)

	//使用UID构建objectId

	oid, _ = uid.ObjectID(iid)
	t.Logf("oid:%v\n", oid)

	sid, iid, _ = oid.Parse()
	t.Logf("sid:%v ,iid:%v \n", sid, iid)

}
