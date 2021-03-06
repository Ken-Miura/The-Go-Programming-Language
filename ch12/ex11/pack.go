// Copyright 2017 Ken Miura
package ex11

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// 入力された構造体のフィールドをクエリパラメータとして返す。
// httpタグを含んでいた場合、そのタグ名がクエリに適用される。
func Pack(v interface{}) string {
	rv := reflect.ValueOf(v).Elem()
	fields := make(map[string][]reflect.Value)
	for i := 0; i < rv.NumField(); i++ {
		fieldInfo := rv.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = append(fields[name], rv.Field(i))
	}

	var buf bytes.Buffer
	i := 0
	for key, values := range fields {
		if checkIfAllValuesAreDefault(values) {
			continue
		}
		if i > 0 {
			buf.WriteByte('&')
		}
		for j, v := range values {
			if j > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(key)
			buf.WriteByte('=')
			buf.WriteString(fmt.Sprintf("%v", v))
		}
		i++
	}
	return buf.String()
}

func checkIfAllValuesAreDefault(values []reflect.Value) bool {
	for _, v := range values {
		if !checkIfValueIsDefault(v) {
			return false
		}
	}
	return true
}

func checkIfValueIsDefault(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0+0i
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}
