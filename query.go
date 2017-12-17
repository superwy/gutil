package gutil

import (
	"net/url"
	"strings"
)

type SortQuery struct {
	Sort  string
	Order string
}

func NewSortQuery(values url.Values, allowSorts []string) *SortQuery {
	sortQuery := &SortQuery{}
	if sortFields, ok := values["sort"]; ok {
		for _, item := range allowSorts {
			if strings.ToLower(sortFields[0]) == strings.ToLower(item) {
				sortQuery.Sort = strings.ToLower(sortFields[0])
				break
			}
		}
	}
	sortQuery.Order = "ASC"
	if orderFields, ok := values["order"]; ok {
		if strings.ToUpper(orderFields[0]) == "DESC" {
			sortQuery.Order = "DESC"
		}
	}
	return sortQuery
}
