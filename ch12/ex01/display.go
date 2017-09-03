// Copyright 2017 Ken Miura
package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
// 練習問題12の1により、配列と構造体の場合のケースを別途追加
// 配列は、丸括弧で囲まれ、その中に要素が列挙される。要素間の区切り文字は空白が利用される。形式としては (要素1 要素2 ...) のようになる。
// 構造体は、丸括弧で囲まれ、その中にエントリ（フィールドと値の組）が列挙される。エントリ間の区切り文字として空白が利用される。
// エントリは、丸括弧で囲まれ、フィールド名と値の組を空白を区切りとして表される。形式としては ((フィールド名1 値1) (フィールド名2 値2) ...) のようになる。
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Array:
		var buf bytes.Buffer
		buf.WriteString("(")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(formatAtom(v.Index(i)))
		}
		buf.WriteString(")")
		return buf.String()
	case reflect.Struct:
		var buf bytes.Buffer
		buf.WriteString("(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(fmt.Sprintf("(%s ", v.Type().Field(i).Name))
			buf.WriteString(formatAtom(v.Field(i)))
			buf.WriteString(")")
		}
		buf.WriteString(")")
		return buf.String()
	default: // reflect.Interface
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
