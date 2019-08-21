
package structs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testItem struct {
	value int
}

func (ti testItem) Less(v Item) bool {
	tip, valid := v.(testItem)
	return valid && ti.value < tip.value
}

func (ti testItem) Eq(v Item) bool {
	tip, valid := v.(testItem)
	return valid && ti.value == tip.value
}

func (ti testItem) String() string {
	return fmt.Sprintf("%d", ti.value)
}

func (ti testItem) Number() int {
	return ti.value
}

//
// Start tests here
// ================
//
func Test_IntItem_Less_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.True(t, (It(i)).Less(It(i+1)))
		assert.False(t, (It(i)).Less(It(i)))
		assert.False(t, (It(i)).Less(It(i-1)))
	}

	// The parameter of the Less function isn't type IntItem
	assert.False(t, (It(1)).Less(testItem{1}))
}

func Test_IntItem_Eq_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.True(t, It(i).Eq(It(i)))
		assert.False(t, It(i).Eq(It(i-1)))
		assert.False(t, It(i).Eq(It(i+1)))
	}

	// The parameter of the Eq function isn't type IntItem
	assert.False(t, (It(1)).Eq(testItem{1}))
}

func Test_IntItem_String_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.Equal(t, (It(i)).String(), fmt.Sprintf("%d", i))
	}
}

func Test_IntItem_Value_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		assert.Equal(t, (It(i)).Value(), i)
	}
}

func Test_It_func(t *testing.T) {
	for i := -100; i <= 100; i++ {
		it := It(i)
		assert.Equal(t, it.value, i)
	}
}


