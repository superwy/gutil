package gutil

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	var result []string
	var keys []interface{}
	keys = append(keys, "a")
	keys = append(keys, "b")
	keys = append(keys, []interface{}{"c", "d", "e"})
	keys = append(keys, []string{"f", "g"})
	m := map[string]interface{}{"field_a": "a", "field_c": "c", "field_d": []interface{}{"a", map[string]interface{}{"a": "c"}}}
	keys = append(keys, m)
	keys = append(keys, map[string]string{"aaaaa": ""})
	GenerateKey(&result, "", keys)
	fmt.Printf("keys %+v \n ", result)
	// [a b c d e f g field_a.a field_c.c field_d.a field_d.a.c aaaaa]
}
