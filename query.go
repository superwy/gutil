package gutil

import (
	"net/url"
	"strings"
)

type QuerySort struct {
	Sort  string
	Order string
}

func NewQuerySort(values url.Values, allowSorts []string) *QuerySort {
	querySort := &QuerySort{}
	if sortFields, ok := values["sort"]; ok {
		for _, item := range allowSorts {
			if strings.ToLower(sortFields[0]) == strings.ToLower(item) {
				querySort.Sort = strings.ToLower(sortFields[0])
				break
			}
		}
	}
	querySort.Order = "ASC"
	if orderFields, ok := values["order"]; ok {
		if strings.ToUpper(orderFields[0]) == "DESC" {
			querySort.Order = "DESC"
		}
	}
	return querySort
}
