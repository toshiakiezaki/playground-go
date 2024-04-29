package validator

import (
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/toshiakiezaki/playground/optional"
)

var (
	instance *validator.Validate
	mu       sync.Mutex
)

func NewValidator() validator.Validate {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = validator.New()
		instance.RegisterCustomTypeFunc(customTypeOptionalString, optional.Optional[*string]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalBool, optional.Optional[*bool]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalFloat32, optional.Optional[*float32]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalFloat64, optional.Optional[*float64]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalInt8, optional.Optional[*int8]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalInt16, optional.Optional[*int16]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalInt32, optional.Optional[*int32]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalInt64, optional.Optional[*int64]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalTime, optional.Optional[*time.Time]{})
		instance.RegisterCustomTypeFunc(customTypeOptionalUUID, optional.Optional[*uuid.UUID]{})
		instance.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return *instance
}

func customTypeOptionalString(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*string])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalBool(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*bool])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalFloat32(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*float32])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalFloat64(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*float64])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalInt8(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*int8])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalInt16(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*int16])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalInt32(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*int32])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalInt64(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*int64])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalTime(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*time.Time])
	if !ok {
		return nil
	}
	return o.Get()
}

func customTypeOptionalUUID(v reflect.Value) interface{} {
	o, ok := v.Interface().(optional.Optional[*uuid.UUID])
	if !ok {
		return nil
	}
	return o.Get()
}
