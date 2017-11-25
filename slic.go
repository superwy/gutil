package gutil

func List2Interface(lst []interface{}) []interface{} {
	if len(lst) <= 0 {
		return []interface{}{}
	}
	result := make([]interface{}, 0, len(lst))
	for _, item := range lst {
		result = append(result, item)
	}
	return result
}
