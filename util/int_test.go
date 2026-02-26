package util

import "testing"

func TestInt64String(t *testing.T) {
	if Int64String(0) != "0" {
		t.Errorf("Int64String(0) = %s", Int64String(0))
	}
	if Int64String(123) != "123" {
		t.Errorf("Int64String(123) = %s", Int64String(123))
	}
}

func TestIntString(t *testing.T) {
	if IntString(0) != "0" {
		t.Errorf("IntString(0) = %s", IntString(0))
	}
	if IntString(42) != "42" {
		t.Errorf("IntString(42) = %s", IntString(42))
	}
}
