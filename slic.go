package gutil

import (
	"github.com/pkg/errors"
	"reflect"
)

func Slice2Interface(arg interface{}) []interface{} {
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
