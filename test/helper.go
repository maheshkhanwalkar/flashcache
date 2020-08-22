package test

import "testing"

func AssertTrue(b bool, t *testing.T) {
	t.Helper()
	if b {
		return
	}

	t.Fatalf("assert failed -- ")
}

func AssertFalse(b bool, t *testing.T) {
	t.Helper()
	AssertTrue(!b, t)
}

func AssertEqual(first interface{}, second interface{}, t *testing.T) {
	t.Helper()
	AssertTrue(first == second, t)
}

func AssertNotEqual(first interface{}, second interface{}, t *testing.T) {
	t.Helper()
	AssertFalse(first == second, t)
}
