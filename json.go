package gutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func maskJsonItem(val interface{}, keys []string) {
	if len(keys) == 0 {
		return
	}
	lastIndex := len(keys) == 1
	firstKey := keys[0]
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Map:
		mapVal, ok := val.(map[string]interface{})
		if ok {
			if lastIndex {
				delete(mapVal, firstKey)
			} else {
				maskJsonItem(mapVal[firstKey], keys[1:])
			}
		}
	case reflect.Slice:
		if sliceVal, ok := val.([]interface{}); ok {
			for _, itemSliceVal := range sliceVal {
				if mapVal, ok := itemSliceVal.(map[string]interface{}); ok {
					if lastIndex {
						delete(mapVal, firstKey)
					} else {
						maskJsonItem(mapVal[firstKey], keys[1:])
					}
				}
			}

		}
	}
}

func MaskJsonWithKeys(struc interface{}, maskKeys []string) (interface{}, error) {
	bytJson, err := json.Marshal(struc)
	if err != nil {
		return nil, err
	}
	var result interface{}
	err = json.Unmarshal(bytJson, &result)
	if err != nil {
		return nil, err
	}
	for _, item := range maskKeys {
		keys := strings.Split(item, ".")
		maskJsonItem(result, keys)

	}
	return result, nil
}

func GenerateKey(allKeys *[]string, key string, val interface{}) {
	processStr := func(keyItem string) {
		if key == "" {
			if keyItem != "" {
				*allKeys = append(*allKeys, keyItem)
			}
		} else {
			if keyItem != "" {
				*allKeys = append(*allKeys, fmt.Sprintf("%s.%s", key, keyItem))
			} else {
				*allKeys = append(*allKeys, key)
			}
		}
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String:
		processStr(val.(string))
	case reflect.Slice:
		valSlice := SliceToInf(val)
		if len(valSlice) == 0 {
			if key != "" {
				*allKeys = append(*allKeys, key)
			}
		} else {
			for _, valSliceItem := range valSlice {
				vSlice := reflect.ValueOf(valSliceItem)
				switch vSlice.Kind() {
				case reflect.String:
					processStr(valSliceItem.(string))
				case reflect.Slice:
					GenerateKey(allKeys, key, valSliceItem)
				case reflect.Map:
					GenerateKey(allKeys, key, valSliceItem)
				}
			}
		}
	case reflect.Map:
		valMap := MapToInf(val)
		for k, v := range valMap {
			kRef := reflect.ValueOf(k)
			if kRef.Kind() == reflect.String {
				kStr := k.(string)
				if kStr != "" {
					newKey := kStr
					if key != "" {
						newKey = fmt.Sprintf("%s.%s", key, newKey)
					}
					GenerateKey(allKeys, newKey, v)
				}
			}
		}
	}
}

func MaskJson(struc interface{}, key interface{}) (interface{}, error) {
	var maskKeys []string
	GenerateKey(&maskKeys, "", key)
	bytJson, err := json.Marshal(struc)
	if err != nil {
		return nil, err
	}
	var result interface{}
	err = json.Unmarshal(bytJson, &result)
	if err != nil {
		return nil, err
	}
	for _, item := range maskKeys {
		keys := strings.Split(item, ".")
		maskJsonItem(result, keys)

	}
	return result, nil
}
