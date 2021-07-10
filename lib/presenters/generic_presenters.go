package presenters

import "github.com/jivison/gowon-indexer/lib/graph/model"

func PresentPageInfo(recordCount int) *model.PageInfo {
	return &model.PageInfo{RecordCount: recordCount}
}
