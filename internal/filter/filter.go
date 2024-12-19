package filter

import (
	"github.com/garrettladley/prods/internal/model/category"
)

type SortBy string

const (
	Price SortBy = "price"
	Name  SortBy = "name"
	Stars SortBy = "stars"
)

var SortByValues = []SortBy{Price, Name, Stars}

type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

var OrderValues = []Order{Asc, Desc}

type Filter struct {
	p Params
}

type Option func(*Filter)

func New(opts ...Option) *Params {
	f := &Filter{}
	for _, opt := range opts {
		opt(f)
	}
	return &f.p
}

func Sort(by SortBy, order Order) Option {
	return func(f *Filter) {
		f.p.Sort = by
		f.p.Order = order
	}
}

func Categories(cats []category.Category) Option {
	return func(f *Filter) {
		f.p.Categories = cats
	}
}

func Paginate(offset uint, limit uint) Option {
	return func(f *Filter) {
		f.p.Offset = offset
		f.p.Limit = limit
	}
}
