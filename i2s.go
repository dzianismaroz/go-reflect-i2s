package main

import (
	"errors"
	"fmt"
	"reflect"
)

var UnmodifiableErr = errors.New("unmodifiable")

// if target is not a pointer / slice -> impossible to modify it.
func validate(out interface{}) error {
	kind := reflect.ValueOf(out).Kind()
	if kind != reflect.Pointer && kind != reflect.Slice {
		return errors.New("target argument is unmodifiable")
	}
	return nil
}

// Function to fill struct fields with values of map.
// @params :
// data - map[string]interface{}
// out - struct{}
func i2s(data interface{}, out interface{}) error {
	// VALIDATE : is assigneable
	if validErr := validate(out); validErr != nil {
		return validErr
	}
	outVal := reflect.ValueOf(out).Elem()
	if err := extract(data, outVal); err != nil {
		return err
	}
	return nil
}

func extract(from interface{}, to reflect.Value) (genErr error) {
	if from == nil {
		return nil // do nothing
	}
	defer func() {
		if r := recover(); r != nil {
			genErr = fmt.Errorf("%s", r)
		}
	}()

	fromVal, toVal := reflect.ValueOf(from), to

	if !toVal.CanSet() {
		return UnmodifiableErr
	}
	switch toVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		toVal.SetInt(int64(fromVal.Float()))
	case reflect.String:
		if toVal.Kind() != fromVal.Kind() {
			return errors.New("wrong type")
		}
		toVal.SetString(fromVal.String())
	case reflect.Bool:
		toVal.SetBool(fromVal.Bool())
	case reflect.Float32, reflect.Float64:
		toVal.SetFloat(fromVal.Float())
	case reflect.Slice:
		rawHolder := from.([]interface{})
		newSlice := reflect.MakeSlice(reflect.SliceOf(to.Type().Elem()), len(rawHolder), cap(rawHolder))
		for i := 0; i < len(rawHolder); i++ {
			if err := extract(rawHolder[i], newSlice.Index(i)); err != nil {
				return err
			}
		}
		to.Set(newSlice)
	case reflect.Struct:
		for i := 0; i < to.NumField(); i++ {
			structField := to.FieldByIndex([]int{i})
			targetName := to.Type().Field(i).Name // resolve struct field name
			rawHolder := from.(map[string]interface{})
			val, exists := rawHolder[targetName]
			if !exists {
				continue
			}
			if err := extract(val, structField); err != nil {
				return err
			}
		}
	case reflect.Pointer:
		return extract(fromVal, to.Elem())
	default:
		return fmt.Errorf("unsupported type: %s", toVal.Type().Name())
	}
	return nil
}
