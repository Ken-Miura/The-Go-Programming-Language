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

func IsCyclic(v interface{}) bool {
	seen := make(map[data]bool)
	return isCyclic(reflect.ValueOf(v), seen)
}

func isCyclic(v reflect.Value, seen map[data]bool) bool {
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
			r := isCyclic(v.Field(i), seen)
			result = result || r
		}
		return result
	case reflect.Ptr, reflect.Interface:
		return isCyclic(v.Elem(), seen)
	case reflect.Slice, reflect.Array:
		result := false
		for i := 0; i < v.Len(); i++ {
			r := isCyclic(v.Index(i), seen)
			result = result || r
		}
		return result
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if isCyclic(v.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	return false
}
