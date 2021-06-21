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

func (qb *QueryBuilder) Insert(table string, params []string) *QueryBuilder {

	var questionsSymbols []string
	for range params {
		questionsSymbols = append(questionsSymbols, "?")
	}

	qb.Query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(params, ","), strings.Join(questionsSymbols, ","))
	return qb
}

func (qb *QueryBuilder) Update(table string, params []string) *QueryBuilder {
	qb.Query = fmt.Sprintf("UPDATE %s SET %s ", table, strings.Join(params, "= ?,"))
	qb.Query += "= ? "
	return qb
}

func (qb *QueryBuilder) Delete(table string) *QueryBuilder {
	qb.Query = fmt.Sprintf("DELETE FROM %s ", table)
	return qb
}

func (qb *QueryBuilder) InnerJoin(joinTable string, onCondition string) *QueryBuilder {
	qb.Query += fmt.Sprintf("INNER JOIN %s ON %s ", joinTable, onCondition)
	return qb
}

func (qb *QueryBuilder) Where(param string) *QueryBuilder {
	qb.Query += fmt.Sprintf("WHERE %s = ?", param)
	return qb
}

func (qb *QueryBuilder) AndWhere(param string) *QueryBuilder {
	qb.Query += fmt.Sprintf(" AND %s = ?", param)
	return qb
}
