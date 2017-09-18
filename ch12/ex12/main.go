// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("you can check if credit number format is valid.\n"))
		writer.Write([]byte("usage: http://localhost:12345/check?cred_num='credit number'\n"))
		writer.Write([]byte("ex: http://localhost:12345/check?cred_num=1234-5678-9012-3456\n"))
	})
	http.HandleFunc("/check", check)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func check(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		CreditNum string `http:"cred_num"`
	}
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}
	fmt.Fprintf(resp, "valid credit number format: %+v\n", data)
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	if len(req.Form) == 0 {
		return fmt.Errorf("no input for credit number format")
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		if name == "cred_num" && !validateCreditNumberFormat(values[0]) {
			return fmt.Errorf("illegal credit number format: %s", values[0])
		}
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func validateCreditNumberFormat(creditNum string) bool {
	r := regexp.MustCompile(`^\d\d\d\d-\d\d\d\d-\d\d\d\d-\d\d\d\d$`)
	return r.MatchString(creditNum)
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
