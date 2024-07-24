package main

import (
	"fmt"
	"reflect"
)

// Function to fill struct fields with values of map.
// @params :
// data - map[string]interface{}
// out - struct{}
func i2s(data interface{}, out interface{}) error {
	dataType := reflect.TypeOf(data)
	dataVal := reflect.ValueOf(data)
	outType := reflect.TypeOf(out).Elem()
	outVal := reflect.ValueOf(out).Elem()

	// tempMap := map[string]interface{}{}

	if dataType.Kind() != reflect.Map || outType.Kind() != reflect.Struct {
		return fmt.Errorf("invalid types provided: %v, %v", dataType.Kind(), outType.Kind())
	}

	fmt.Println("we got: ")
	fmt.Println("data:", dataType.Kind(), "out:", outType.Kind(), "name:", outType.Name())
	for _, e := range dataVal.MapKeys() {
		f := outVal.FieldByName(e.String())
		if !f.IsValid() || !f.CanSet() {
			return fmt.Errorf("unable to set field  %s", e.String())
		}
		val := dataVal.MapIndex(e)
		switch f.Kind() {
		case reflect.String:
			f.SetString(val.String())
		case reflect.Int, reflect.Float64:
			f.SetInt(val.Int())
		case reflect.Bool:
			f.SetBool(val.Bool())
		}

		// tempMap[e.String()] = dataVal.MapIndex(e).Interface()
	}

	return nil
}
