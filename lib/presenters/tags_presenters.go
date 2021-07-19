package presenters

import (
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/analysis"
)

func PresentTagsResponse(tags []analysis.TagResponse, count int) *model.TagsResponse {
	response := &model.TagsResponse{
		PageInfo: PresentPageInfo(count),
	}

	for _, tag := range tags {
		response.Tags = append(response.Tags, PresentTag(tag))
	}

	return response
}

func PresentTag(tag analysis.TagResponse) *model.Tag {
	return &model.Tag{
		Name:        tag.Name,
		Occurrences: tag.Occurrences,
	}
}
