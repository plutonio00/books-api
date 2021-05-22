package mysql

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	Query string
}

func newQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Query: "",
	}
}

func (qb *QueryBuilder) Select(selected []string, table string) *QueryBuilder {
	qb.Query = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(selected, ","), table)
	return qb
}

func (qb *QueryBuilder) Insert() {

}

func (qb *QueryBuilder) Update() {

}

func (qb *QueryBuilder) Delete() {

}

func (qb *QueryBuilder) InnerJoin(joinTable string, onCondition string) *QueryBuilder {
	qb.Query += fmt.Sprintf("INNER JOIN %s ON %s ", joinTable, onCondition)
	return qb
}

func (qb *QueryBuilder) Where(param string) *QueryBuilder {
	qb.Query += fmt.Sprintf("WHERE %s = ?", param)
	return qb
}
