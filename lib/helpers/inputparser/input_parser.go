package inputparser

import "github.com/go-pg/pg/v10/orm"

type InputParser struct {
	query *orm.Query
}

func (p InputParser) GetQuery() *orm.Query {
	return p.query
}

func CreateParser(query *orm.Query) *InputParser {
	return &InputParser{query: query}
}
