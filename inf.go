package gutil

import (
	"github.com/pkg/errors"
	"reflect"
)

func SliceToInf(arg interface{}) []interface{} {
	slice := reflect.ValueOf(arg)
	if slice.Kind() != reflect.Slice {
		panic(errors.New("bad slice"))
	}
	count := slice.Len()
	result := make([]interface{}, count)
	for i := 0; i < count; i++ {
		result[i] = slice.Index(i).Interface()
	}
	return result
}

func MapToInf(arg interface{}) map[interface{}]interface{} {
	m := reflect.ValueOf(arg)
	if m.Kind() != reflect.Map {
		panic(errors.New("bad map"))
	}
	result := make(map[interface{}]interface{})
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k).Interface()
		result[k.Interface()] = v
	}
	return result
}
