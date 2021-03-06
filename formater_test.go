package golog

import (
	"testing"
)

func TestSimpleFormater(t *testing.T) {
	f := NewSimpleFormater()

	b := f.Format(LEVEL_EMERGENCY, []byte("abc"))
	t.Log(string(b))
}

func TestWebFormater(t *testing.T) {
	f := NewWebFormater([]byte("xyz"), []byte("10.0.0.1"))

	b := f.Format(LEVEL_EMERGENCY, []byte("abc"))
	t.Log(string(b))
}
