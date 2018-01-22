package gutil

import (
	"fmt"
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
		} else if strings.ToUpper(orderFields[0]) == "ASC" {
			sortQuery.Desc = false
		}
	}
	if sortQuery.Field == "" {
		return nil
	}
	return sortQuery
}

func (sort *SortQuery) ToString() string {
	if sort.Desc {
		return fmt.Sprintf("%s %s", sort.Field, "DESC NULLS LAST")
	} else {
		return fmt.Sprintf("%s %s", sort.Field, "ASC NULLS FIRST")
	}
}
