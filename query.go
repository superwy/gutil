package gutil

import (
	"fmt"
	"github.com/go-pg/pg/orm"
	"net/url"
	"strconv"
	"strings"
)

//region Sort
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

//endregion

//region PagerQuery
type PagerQuery struct {
	Limit     int
	Offset    int
	MaxLimit  int
	MaxOffset int
}

func NewPagerQuery(values url.Values) (*PagerQuery, error) {
	p := new(PagerQuery)
	err := p.SetURLValues(values)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *PagerQuery) SetURLValues(values url.Values) error {
	limit, err := intParam(values, "limit")
	if err != nil {
		return err
	}
	p.Limit = limit

	page, err := intParam(values, "page")
	if err != nil {
		return err
	}
	if page > 0 {
		p.SetPage(page)
	}
	return nil
}

func (p *PagerQuery) maxLimit() int {
	if p.MaxLimit > 0 {
		return p.MaxLimit
	}
	return 1000
}

func (p *PagerQuery) maxOffset() int {
	if p.MaxOffset > 0 {
		return p.MaxOffset
	}
	return 1000000
}

func (p *PagerQuery) GetLimit() int {
	const defaultLimit = 100

	if p.Limit <= 0 {
		return defaultLimit
	}
	if p.Limit > p.maxLimit() {
		return p.maxLimit()
	}
	return p.Limit
}

func (p *PagerQuery) GetOffset() int {
	if p.Offset > p.maxOffset() {
		return p.maxOffset()
	}
	if p.Offset > 0 {
		return p.Offset
	}
	return 0
}

func (p *PagerQuery) SetPage(page int) {
	p.Offset = (page - 1) * p.GetLimit()
}

func (p *PagerQuery) GetPage() int {
	return (p.GetOffset() / p.GetLimit()) + 1
}

func (p *PagerQuery) Paginate(q *orm.Query) (*orm.Query, error) {
	q = q.Limit(p.GetLimit()).
		Offset(p.GetOffset())
	return q, nil
}

func intParam(urlValues url.Values, paramName string) (int, error) {
	values, ok := urlValues[paramName]
	if !ok {
		return 0, nil
	}

	value, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, fmt.Errorf("param=%s value=%s is invalid: %s", paramName, values[0], err)
	}

	return value, nil
}

//endregion
