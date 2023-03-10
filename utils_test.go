package arrayfuncs_test

import (
	"testing"

	arrayFuncs "github.com/izacgaldino23/array-funcs"
	"github.com/stretchr/testify/assert"
)

type Temp struct {
	msg   string
	order int
}

func (t *Temp) ToString() string {
	return t.msg
}

func TestAnyToString(t *testing.T) {
	assert.Equal(t, "", arrayFuncs.AnyToString(nil))
	assert.Equal(t, "10", arrayFuncs.AnyToString(10))
	assert.Equal(t, "true", arrayFuncs.AnyToString(true))
	assert.Equal(t, "10.5", arrayFuncs.AnyToString(10.5))
	assert.Equal(t, "temp", arrayFuncs.AnyToString("temp"))

	temp := &Temp{"test", 1}

	assert.Equal(t, temp.msg, arrayFuncs.AnyToString(temp))
}
