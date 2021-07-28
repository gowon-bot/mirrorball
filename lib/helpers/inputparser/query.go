package inputparser

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (p InputParser) ParsePageInput(pageInput *model.PageInput) *InputParser {
	if pageInput == nil {
		return &p
	}

	if pageInput.Limit != nil {
		p.query.Limit(*pageInput.Limit)
	}

	if pageInput.Offset != nil {
		p.query.Offset(*pageInput.Offset)
	}

	return &p
}

type SortSettings interface {
	getDefaultSort() string
}

func (p InputParser) ParseSort(sort *string, settings SortSettings) *InputParser {
	if sort != nil {
		p.query.Order(*sort)
	} else {
		defaultSort := settings.getDefaultSort()

		if defaultSort[0] >= '0' && defaultSort[0] <= '9' {
			p.query.OrderExpr(defaultSort)
		} else {
			p.query.Order(defaultSort)
		}
	}

	return &p
}

func (p InputParser) ParseTimerange(timerange model.Timerange, field string) {
	if timerange.From != nil {
		p.query.Where(field+" > to_timestamp(?)", timerange.From)
	}
	if timerange.To != nil {
		p.query.Where(field+" < to_timestamp(?)", timerange.To)
	}
}
