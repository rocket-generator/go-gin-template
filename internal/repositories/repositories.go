package repositories

import (
	"github.com/uptrace/bun"
	"time"
)

// Filter ...
type Filter struct {
	Column   string
	Operator string
	Value    interface{}
}

func ApplyFilters(tableName string, query *bun.SelectQuery, filters []Filter) *bun.SelectQuery {
	for _, filter := range filters {
		operator := filter.Operator
		if operator == "" {
			operator = "="
		}
		query = query.Where(tableName+"."+filter.Column+" "+filter.Operator+" ?", filter.Value)
	}
	return query
}

func ApplyDeleteFilters(tableName string, query *bun.DeleteQuery, filters []Filter) *bun.DeleteQuery {
	for _, filter := range filters {
		operator := filter.Operator
		if operator == "" {
			operator = "="
		}
		query = query.Where(tableName+"."+filter.Column+" "+filter.Operator+" ?", filter.Value)
	}
	return query
}

func WhereCurrentlyAvailable(tableName string, query *bun.SelectQuery) *bun.SelectQuery {
	return query.Where(tableName+".available_from < ?", time.Now().Unix()).
		Where(tableName+".available_to > ?", time.Now().Unix())
}
