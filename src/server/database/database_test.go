package database

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	db := New()
	db.Add(interface{}(0), interface{}("Test"))
	assertEq(t, db.Get(0).(string), "Test")
	db.Delete(interface{}(0))
	assertEq(t, db.Get(0), nil)

	str := []byte{'A'}
	for i := 0; i < 3; i++ {
		db.Add(interface{}(i), interface{}(string(str)))
		checkEq(t, db.Get(i).(string), string(str))
		str[0]++
	}
	if t.Failed() {
		t.FailNow()
	}

	db.Add(interface{}(1), interface{}("D"))
	assertEq(t, db.Get(1).(string), "D")
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
