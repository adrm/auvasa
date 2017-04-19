package auvasa

import (
	"testing"
)

func TestGet(t *testing.T) {
	ok, err := Get(812)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(ok)
}
