package validator

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/toshiakiezaki/playground/optional"
	"github.com/toshiakiezaki/playground/util"
)

func TestCustomTypeOptionalString(t *testing.T) {
	var nilValue *string

	testCases := []struct {
		name  string
		value *optional.Optional[*string]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.String(""))},
		{name: "InputValue", value: optional.WithValue(util.String("value"))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalString(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalBool(t *testing.T) {
	var nilValue *bool

	testCases := []struct {
		name  string
		value *optional.Optional[*bool]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Bool(false))},
		{name: "InputValue", value: optional.WithValue(util.Bool(true))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalBool(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalFloat32(t *testing.T) {
	var nilValue *float32

	testCases := []struct {
		name  string
		value *optional.Optional[*float32]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Float32(0))},
		{name: "InputValue", value: optional.WithValue(util.Float32(100.53))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalFloat32(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalFloat64(t *testing.T) {
	var nilValue *float64

	testCases := []struct {
		name  string
		value *optional.Optional[*float64]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Float64(0))},
		{name: "InputValue", value: optional.WithValue(util.Float64(100.53))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalFloat64(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalInt8(t *testing.T) {
	var nilValue *int8

	testCases := []struct {
		name  string
		value *optional.Optional[*int8]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Int8(0))},
		{name: "InputValue", value: optional.WithValue(util.Int8(100))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalInt8(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalInt16(t *testing.T) {
	var nilValue *int16

	testCases := []struct {
		name  string
		value *optional.Optional[*int16]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Int16(0))},
		{name: "InputValue", value: optional.WithValue(util.Int16(100))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalInt16(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalInt32(t *testing.T) {
	var nilValue *int32

	testCases := []struct {
		name  string
		value *optional.Optional[*int32]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Int32(0))},
		{name: "InputValue", value: optional.WithValue(util.Int32(100))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalInt32(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalInt64(t *testing.T) {
	var nilValue *int64

	testCases := []struct {
		name  string
		value *optional.Optional[*int64]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Int64(0))},
		{name: "InputValue", value: optional.WithValue(util.Int64(100))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalInt64(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalTime(t *testing.T) {
	var nilValue *time.Time

	testCases := []struct {
		name  string
		value *optional.Optional[*time.Time]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(util.Time(time.Time{}))},
		{name: "InputValue", value: optional.WithValue(util.Time(time.Now()))},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalTime(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func TestCustomTypeOptionalUUID(t *testing.T) {
	var nilValue *uuid.UUID
	var blankValue uuid.UUID
	var inputValue uuid.UUID

	inputValue = uuid.New()
	testCases := []struct {
		name  string
		value *optional.Optional[*uuid.UUID]
	}{
		{name: "NullInstance", value: nil},
		{name: "NullValue", value: optional.WithValue(nilValue)},
		{name: "BlankValue", value: optional.WithValue(&blankValue)},
		{name: "InputValue", value: optional.WithValue(&inputValue)},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var reflectedValue reflect.Value
			if tC.value != nil {
				reflectedValue = reflect.ValueOf(*tC.value)
			} else {
				reflectedValue = reflect.New(reflect.TypeOf(tC.value))
			}

			actual := customTypeOptionalUUID(reflectedValue)
			if tC.value != nil {
				assert.Equal(t, tC.value.Get(), actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}
