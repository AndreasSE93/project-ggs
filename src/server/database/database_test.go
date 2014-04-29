package database

import (
	"testing"
	"server/connection"
)

type fakeAddr struct {
	Str string
}
func (a fakeAddr) String() string {
	return a.Str
}
func (a fakeAddr) Network() string {
	return ""
}

func TestDatabase(t *testing.T) {
	var con [4]connection.Connector
	str := []byte{'A'}
	for i := 0; i < len(con); i++ {
		con[i] = connection.Connector{ConnectorID: i, LocalAddr: fakeAddr{string(str)}}
		str[0]++
	}

	db := New()
	db.Add(con[0])
	assertEq(t, db.Get(0).LocalAddr.String(), "A")
	db.Delete(con[0])
	assertEq(t, db.Get(0).LocalAddr, nil)

	str = []byte{'A'}
	for i := 0; i < 3; i++ {
		db.Add(con[i])
		checkEq(t, db.Get(i).LocalAddr.String(), string(str))
		str[0]++
	}
	if t.Failed() {
		t.FailNow()
	}

	db.Update(1, con[3])
	assertEq(t, db.Get(1).LocalAddr.String(), "D")
}

func checkEq(t *testing.T, A interface{}, B interface{}) bool {
	if A != B {
		t.Error(A, "!=", B)
		return false
	} else if testing.Verbose() {
		t.Log(A, "==", B)
	}
	return true
}

func assertEq(t *testing.T, A interface{}, B interface{}) {
	if ! checkEq(t, A, B) {
		t.FailNow()
	}
}
