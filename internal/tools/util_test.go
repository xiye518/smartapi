package tools

import "testing"

func TestGetGuid(t *testing.T) {
	s := GetGuid()
	t.Log(s)
}
