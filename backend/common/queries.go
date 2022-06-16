package common

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	routing "github.com/go-ozzo/ozzo-routing"
)

type BaseQuery struct {
	Pager  Pager
	Sorts  []Sort
	Fields []string
}

type Pager struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type Sort struct {
	Field string `json:"field"`
	Desc  bool   `json:"desc"`
}

func (q *BaseQuery) OrderBy() string {
	res := ""
	if q.Sorts != nil {
		for i := range q.Sorts {
			if q.Sorts[i].Field != "" {
				if q.Sorts[i].Desc {
					res += `"` + q.Sorts[i].Field + `" desc`
				} else {
					res += `"` + q.Sorts[i].Field + `"`
				}
			}
			if i < len(q.Sorts) {
				res += ","
			}
		}
	}

	return strings.TrimSuffix(res, ",")
}

func (q *BaseQuery) ParseBase(ctx *routing.Context) error {
	q.parsePager(ctx)

	if err := q.parseSort(ctx); err != nil {
		return err
	}

	return nil
}

func (q *BaseQuery) parsePager(ctx *routing.Context) {
	limQ := ctx.Query("lim")
	offsetQ := ctx.Query("offset")

	lim, err := strconv.ParseInt(limQ, 10, 0)
	if err != nil {
		lim = 100
	}

	offset, err := strconv.ParseInt(offsetQ, 10, 0)
	if err != nil {
		offset = 0
	}

	q.Pager = Pager{
		Limit:  lim,
		Offset: offset,
	}
}

func (q *BaseQuery) parseSort(ctx *routing.Context) error {
	q.Sorts = make([]Sort, 0)

	s := ctx.Query("s")
	if s == "" {
		return nil
	}

	jn, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return NewBadRequest("incorrect query params", "", s)
	}

	if err := json.Unmarshal(jn, &q.Sorts); err != nil {
		return NewBadRequest("incorrect query params", "", jn)
	}

	return nil
}

func ParseQuery(ctx *routing.Context, q interface{}) error {
	qs := ctx.Query("q")
	if qs == "" {
		return nil
	}

	jn, err := base64.StdEncoding.DecodeString(qs)
	if err != nil {
		return NewBadRequest("incorrect query params", "", qs)
	}

	if err := json.Unmarshal(jn, &q); err != nil {
		return NewBadRequest("incorrect query params", "", jn)
	}

	return nil
}

type BoolSearch struct {
	HasValue bool `json:"hasValue"`
	Value    bool `json:"value"`
}
