package gutil

import (
	"net/url"
	"strings"
)

type SortQuery struct {
	Field string
	Desc  bool
}

func NewSortQuery(values url.Values, defaultDesc bool, fieldFn func(sortName string) string) *SortQuery {
	sortQuery := &SortQuery{}
	if sortFields, ok := values["sort"]; ok {
		if fieldFn != nil {
			sortQuery.Field = fieldFn(sortFields[0])
		} else {
			sortQuery.Field = sortFields[0]
		}
	}
	sortQuery.Desc = defaultDesc
	if orderFields, ok := values["order"]; ok {
		if strings.ToUpper(orderFields[0]) == "DESC" {
			sortQuery.Desc = true
		}
	}
	return sortQuery
}
