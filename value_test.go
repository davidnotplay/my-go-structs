
package structs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testValue struct {
	value int
}

func (tv testValue) Less(v Value) bool {
	intvalue, valid := v.(testValue)
	return valid &&  tv.value < intvalue.value
}

func (tv testValue) Eq(v Value) bool {
	intvalue, valid := v.(testValue)
	return valid &&  tv.value == intvalue.value
}

func (tv testValue) String() string {
	return fmt.Sprintf("%d", tv.value)
}

func (tv testValue) Key() Key {
	return tv.value
}

func (tv testValue) LessKey(k Key) bool {
	num, valid := k.(int)
	return valid && tv.value < num
}

func (tv testValue) EqKey(k Key) bool {
	num, valid := k.(int)
	return valid && tv.value == num
}

func (tv testValue) Number() int {
	return tv.value
}

//
// Start tests here
// ================
//
func Test_IntValue_Less_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.True(t, (IntValue{i}).Less(IntValue{i+1}))
		assert.False(t, (IntValue{i}).Less(IntValue{i}))
		assert.False(t, (IntValue{i}).Less(IntValue{i-1}))
	}

	// The parameter of the Less function isn't type IntValue
	assert.False(t, (IntValue{1}).Less(testValue{1}))
}

func Test_IntValue_Eq_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.True(t, (IntValue{i}).Eq(IntValue{i}))
		assert.False(t, (IntValue{i}).Eq(IntValue{i-1}))
		assert.False(t, (IntValue{i}).Eq(IntValue{i+1}))
	}

	// The parameter of the Eq function isn't type IntValue
	assert.False(t, (IntValue{1}).Eq(testValue{1}))
}

func Test_IntValue_String_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.Equal(t, (IntValue{i}).String(), fmt.Sprintf("%d", i))
	}
}

func Test_IntValue_Key_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.Equal(t, (IntValue{i}).Key(), i)
	}
}

func Test_IntValue_LessKey_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.True(t, (IntValue{i}).LessKey(i+1))
		assert.False(t, (IntValue{i}).LessKey(i))
		assert.False(t, (IntValue{i}).LessKey(i-1))
	}

	// The parameter of the LessKey function isn't type IntValue
	assert.False(t, (IntValue{1}).LessKey(3.22))
	assert.False(t, (IntValue{1}).LessKey(int32(2)))
}

func Test_IntValue_EqKey_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.True(t, (IntValue{i}).EqKey(i))
		assert.False(t, (IntValue{i}).EqKey(i - 1))
		assert.False(t, (IntValue{i}).EqKey(i + 1))
	}

	// The parameter of the LessKey function isn't type IntValue
	assert.False(t, (IntValue{1}).EqKey(IntValue{1}))
}

func Test_IntValue_Value_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.Equal(t, (IntValue{i}).Value(), i)
	}
}

func Test_Iv_fun(t *testing.T) {
	for i := -100; i <= 100; i++ {
		iv := Iv(i)
		assert.Equal(t, iv.value, i)
	}
}


