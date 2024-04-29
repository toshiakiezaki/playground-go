package optional

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type (
	Types interface {
		*string | *bool | *float32 | *float64 | *int8 | *int16 | *int32 | *int64 | *time.Time | *uuid.UUID
	}

	Optional[E Types] struct {
		value E
		set   bool
	}
)

func (o *Optional[E]) Get() E {
	return o.value
}

func (o *Optional[E]) IsPresent() bool {
	var empty E
	return o.value != empty
}

func (o *Optional[E]) IsSet() bool {
	return o.set
}

func (o *Optional[E]) MarshalJSON() ([]byte, error) {
	if o.IsSet() && o.IsPresent() {
		return json.Marshal(o.value)
	}
	return []byte("null"), nil
}

func (o *Optional[E]) UnmarshalJSON(data []byte) error {
	o.set = true
	return json.Unmarshal(data, &o.value)
}

func WithValue[E Types](data E) *Optional[E] {
	return &Optional[E]{
		value: data,
		set:   true,
	}
}
