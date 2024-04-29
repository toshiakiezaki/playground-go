package util

import (
	"time"
)

func Bool(v bool) *bool {
	return &v
}

func BoolSlice(vs []bool) []*bool {
	ps := make([]*bool, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func BoolMap(vs map[string]bool) map[string]*bool {
	ps := make(map[string]*bool, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Byte(v byte) *byte {
	return &v
}

func ByteSlice(vs []byte) []*byte {
	ps := make([]*byte, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func ByteMap(vs map[string]byte) map[string]*byte {
	ps := make(map[string]*byte, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func String(v string) *string {
	return &v
}

func StringSlice(vs []string) []*string {
	ps := make([]*string, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func StringMap(vs map[string]string) map[string]*string {
	ps := make(map[string]*string, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Int(v int) *int {
	return &v
}

func IntSlice(vs []int) []*int {
	ps := make([]*int, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func IntMap(vs map[string]int) map[string]*int {
	ps := make(map[string]*int, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Int8(v int8) *int8 {
	return &v
}

func Int8Slice(vs []int8) []*int8 {
	ps := make([]*int8, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Int8Map(vs map[string]int8) map[string]*int8 {
	ps := make(map[string]*int8, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Int16(v int16) *int16 {
	return &v
}

func Int16Slice(vs []int16) []*int16 {
	ps := make([]*int16, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Int16Map(vs map[string]int16) map[string]*int16 {
	ps := make(map[string]*int16, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Int32(v int32) *int32 {
	return &v
}

func Int32Slice(vs []int32) []*int32 {
	ps := make([]*int32, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Int32Map(vs map[string]int32) map[string]*int32 {
	ps := make(map[string]*int32, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Int64(v int64) *int64 {
	return &v
}

func Int64Slice(vs []int64) []*int64 {
	ps := make([]*int64, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Int64Map(vs map[string]int64) map[string]*int64 {
	ps := make(map[string]*int64, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Uint(v uint) *uint {
	return &v
}

func UintSlice(vs []uint) []*uint {
	ps := make([]*uint, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func UintMap(vs map[string]uint) map[string]*uint {
	ps := make(map[string]*uint, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Uint8(v uint8) *uint8 {
	return &v
}

func Uint8Slice(vs []uint8) []*uint8 {
	ps := make([]*uint8, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Uint8Map(vs map[string]uint8) map[string]*uint8 {
	ps := make(map[string]*uint8, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Uint16(v uint16) *uint16 {
	return &v
}

func Uint16Slice(vs []uint16) []*uint16 {
	ps := make([]*uint16, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Uint16Map(vs map[string]uint16) map[string]*uint16 {
	ps := make(map[string]*uint16, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Uint32(v uint32) *uint32 {
	return &v
}

func Uint32Slice(vs []uint32) []*uint32 {
	ps := make([]*uint32, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Uint32Map(vs map[string]uint32) map[string]*uint32 {
	ps := make(map[string]*uint32, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Uint64(v uint64) *uint64 {
	return &v
}

func Uint64Slice(vs []uint64) []*uint64 {
	ps := make([]*uint64, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Uint64Map(vs map[string]uint64) map[string]*uint64 {
	ps := make(map[string]*uint64, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Float32(v float32) *float32 {
	return &v
}

func Float32Slice(vs []float32) []*float32 {
	ps := make([]*float32, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Float32Map(vs map[string]float32) map[string]*float32 {
	ps := make(map[string]*float32, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Float64(v float64) *float64 {
	return &v
}

func Float64Slice(vs []float64) []*float64 {
	ps := make([]*float64, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func Float64Map(vs map[string]float64) map[string]*float64 {
	ps := make(map[string]*float64, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Time(v time.Time) *time.Time {
	return &v
}

func TimeSlice(vs []time.Time) []*time.Time {
	ps := make([]*time.Time, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func TimeMap(vs map[string]time.Time) map[string]*time.Time {
	ps := make(map[string]*time.Time, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}

func Duration(v time.Duration) *time.Duration {
	return &v
}

func DurationSlice(vs []time.Duration) []*time.Duration {
	ps := make([]*time.Duration, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}

	return ps
}

func DurationMap(vs map[string]time.Duration) map[string]*time.Duration {
	ps := make(map[string]*time.Duration, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}

	return ps
}
