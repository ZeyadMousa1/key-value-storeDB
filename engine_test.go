package main

import "testing"

func Test_SetGetValue(t *testing.T){
	e := NewEngien()

	e.Set("foo", "bar")
	val, err := e.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if val != "bar" {
		t.Error("value should be bar")
	}
	_, err = e.Get("notfound")
	if err == nil {
		t.Error("should return error")
	}
}