// Copyright 2017 Ken Miura
package ex06

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if checkIfValueIsDefault(v.Field(i)) {
				continue
			}
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if checkIfValueIsDefault(v.MapIndex(key)) {
				continue
			}
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "t")
		} else {
			fmt.Fprint(buf, "nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%g %g)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		if v.IsNil() {
			buf.WriteString("nil")
		} else {
			buf.WriteString("(")
			buf.WriteString(fmt.Sprintf("%q", v.Elem().Type().String()))
			buf.WriteByte(' ')
			if err := encode(buf, v.Elem()); err != nil {
				return err
			}
			buf.WriteByte(')')
		}

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func checkIfValueIsDefault(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() == 0 {
			return true
		} else {
			return false
		}
	case reflect.String:
		if v.String() == "" {
			return true
		} else {
			return false
		}
	case reflect.Bool:
		if v.Bool() {
			return false
		} else {
			return true
		}

	case reflect.Float32, reflect.Float64:
		if v.Float() == 0.0 {
			return true
		} else {
			return false
		}

	case reflect.Complex64, reflect.Complex128:
		if v.Complex() == 0+0i {
			return true
		} else {
			return false
		}

	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface, reflect.Chan, reflect.Func:
		if v.IsNil() {
			return true
		} else {
			return false
		}

	default:
		return false
	}
}
