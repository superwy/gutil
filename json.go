package gutil

import (
	"encoding/json"
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

func MaskJsonWithKeys(inf interface{}, maskKeys []string) error {
	bytJson, err := json.Marshal(inf)
	if err != nil {
		return err
	}
	var jsonMap interface{}
	err = json.Unmarshal(bytJson, &jsonMap)
	if err != nil {
		return err
	}
	for _, item := range maskKeys {
		keys := strings.Split(item, ".")
		maskJsonItem(jsonMap, keys)
	}
	return nil
}