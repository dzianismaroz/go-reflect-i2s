package main

import (
	"errors"
	"fmt"
	"reflect"
)

func validate(in, out interface{}) error {
	outType := reflect.TypeOf(out)
	if outType.Kind() != reflect.Pointer {
		return errors.New("target argument is unmodifiable")
	}

	if reflect.TypeOf(in).Kind() != reflect.Map {
		return errors.New("source argument is not a map")
	}
	return nil
}

// Function to fill struct fields with values of map.
// @params :
// data - map[string]interface{}
// out - struct{}
func i2s(data interface{}, out interface{}) error {
	// VALIDATE
	if validErr := validate(data, out); validErr != nil {
		return validErr
	}
	dataVal := reflect.ValueOf(data)
	outVal := reflect.ValueOf(out).Elem() // origignal strucy behing the pointer.

	rawHolder := map[string]interface{}{}

	for _, e := range dataVal.MapKeys() {
		key := e.String()
		rawHolder[key] = dataVal.MapIndex(e).Interface()

	}
	fmt.Println("raw Holder :", rawHolder)

	for i := 0; i < outVal.NumField(); i++ {
		f := outVal.FieldByIndex([]int{i})
		targetName := outVal.Type().Field(i).Name
		fmt.Print("target name:", targetName, " and val: ")
		if !f.IsValid() {
			return fmt.Errorf("field %s is invalid", targetName)
		}
		if !f.CanSet() {
			return fmt.Errorf("field %s is not modifiable", targetName)
		}
		val, exists := rawHolder[targetName]
		if !exists {
			continue
		}
		fmt.Println(val)
		switch f.Kind() {
		case reflect.String:
			switch val.(type) {
			case string: //is ok
			default:
				return errors.New("type mismatch")
			}
			f.SetString(val.(string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch val.(type) {
			case float64: //is ok
			default:
				return errors.New("type mismatch")
			}
			f.SetInt(int64(val.(float64)))
		case reflect.Float32, reflect.Float64:
			switch val.(type) {
			case float64: //is ok
			default:
				return errors.New("type mismatch")
			}
			f.SetFloat(val.(float64))
		case reflect.Bool:
			switch val.(type) {
			case bool: //is ok
			default:
				return errors.New("type mismatch")
			}
			f.SetBool(val.(bool))
		}
	}

	return nil
}

func extractFrom(holder map[string]interface{}) interface{} {

}
