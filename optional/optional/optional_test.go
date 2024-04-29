package optional

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/toshiakiezaki/playground/util"
)

func TestOptionalGet(t *testing.T) {
	ptrTestCases := []struct {
		name     string
		value    *string
		expected *string
	}{
		{name: "NullValue", value: nil, expected: nil},
		{name: "BlankValue", value: util.String(""), expected: util.String("")},
		{name: "InputValue", value: util.String("value"), expected: util.String("value")},
	}
	for _, tC := range ptrTestCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := Optional[*string]{
				value: tC.value,
			}
			assert.Equal(t, tC.expected, actual.Get())
		})
	}
}

func TestOptionalIsPresent(t *testing.T) {
	ptrTestCases := []struct {
		name     string
		value    *string
		expected bool
	}{
		{name: "NullValue", value: nil, expected: false},
		{name: "BlankValue", value: util.String(""), expected: true},
		{name: "InputValue", value: util.String("value"), expected: true},
	}
	for _, tC := range ptrTestCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := Optional[*string]{
				value: tC.value,
			}
			assert.Equal(t, tC.expected, actual.IsPresent())
		})
	}
}

func TestOptionalMarshalJSON(t *testing.T) {
	opt1 := Optional[*string]{value: nil, set: true}
	actual1, err1 := opt1.MarshalJSON()
	expected1 := []byte(`null`)
	opt2 := Optional[*string]{value: util.String(""), set: true}
	actual2, err2 := opt2.MarshalJSON()
	expected2 := []byte(`""`)
	opt3 := Optional[*string]{value: util.String("value"), set: true}
	actual3, err3 := opt3.MarshalJSON()
	expected3 := []byte(`"value"`)

	assert.NoError(t, err1)
	assert.Equal(t, string(expected1), string(actual1))
	assert.NoError(t, err2)
	assert.Equal(t, string(expected2), string(actual2))
	assert.NoError(t, err3)
	assert.Equal(t, string(expected3), string(actual3))
}

func TestOptionalUnmarshalJSON(t *testing.T) {
	value1 := []byte(`null`)
	actual1 := Optional[*string]{}
	err1 := actual1.UnmarshalJSON(value1)
	expected1 := Optional[*string]{value: nil, set: true}
	value2 := []byte(`""`)
	actual2 := Optional[*string]{}
	err2 := actual2.UnmarshalJSON(value2)
	expected2 := Optional[*string]{value: util.String(""), set: true}
	value3 := []byte(`"value"`)
	actual3 := Optional[*string]{}
	err3 := actual3.UnmarshalJSON(value3)
	expected3 := Optional[*string]{value: util.String("value"), set: true}

	assert.NoError(t, err1)
	assert.Equal(t, expected1, actual1)
	assert.NoError(t, err2)
	assert.Equal(t, expected2, actual2)
	assert.NoError(t, err3)
	assert.Equal(t, expected3, actual3)
}

func TestWithValue(t *testing.T) {
	actual := WithValue(util.Time(time.Now()))
	assert.NotNil(t, actual)
	assert.True(t, actual.IsPresent())
	assert.NotNil(t, actual.Get())
}
