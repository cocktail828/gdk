package errcode

import (
	"testing"
)

func TestIsNil(t *testing.T) {
	var k interface{} = 4
	t.Log(IsNil(k))
}
