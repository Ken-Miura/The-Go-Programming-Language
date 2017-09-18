// Copyright 2017 Ken Miura
package ex02

import (
	"reflect"
	"unsafe"
)

type data struct {
	x unsafe.Pointer
	t reflect.Type
}

func IsCircularDataStructure(v interface{}) bool {
	seen := make(map[data]bool)
	return isCircularDataStructure(reflect.ValueOf(v), seen)
}

func isCircularDataStructure(v reflect.Value, seen map[data]bool) bool {
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		d := data{vptr, v.Type()}
		if seen[d] {
			return true
		}
		seen[d] = true
	}
	switch v.Kind() {
	case reflect.Struct:
		result := false
		for i := 0; i < v.NumField(); i++ {
			r := isCircularDataStructure(v.Field(i), seen)
			result = result || r
		}
		return result
	case reflect.Ptr, reflect.Interface:
		return isCircularDataStructure(v.Elem(), seen)
	case reflect.Slice, reflect.Array:
		result := false
		for i := 0; i < v.Len(); i++ {
			r := isCircularDataStructure(v.Index(i), seen)
			result = result || r
		}
		return result
	case reflect.Map:
		result := false
		for _, k := range v.MapKeys() {
			r := isCircularDataStructure(v.MapIndex(k), seen)
			result = result || r
		}
		return result
	}
	return false
}
