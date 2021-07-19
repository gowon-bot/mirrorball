package inputparser

import (
	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (p InputParser) ParseTagsSettings(tagSettings *model.TagsSettings) *InputParser {
	if tagSettings == nil {
		return &p
	}

	if tagSettings.Artists != nil {
		var artistNames []string

		for _, artist := range tagSettings.Artists {
			if artist.Name == nil {
				continue
			}

			artistNames = append(artistNames, *artist.Name)
		}

		p.query.
			Join(`JOIN artist_tags "artist_tag" ON "tag".id = "artist_tag".tag_id`).
			Join(`JOIN artists "artist" ON "artist".id = "artist_tag".artist_id AND "artist".name IN (?)`, pg.In(artistNames))
	}

	if tagSettings.Keyword != nil {
		p.query.Where("name ILIKE ?", "%"+*tagSettings.Keyword+"%")
	}

	if tagSettings.PageInput != nil {
		p.ParsePageInput(tagSettings.PageInput)
	}

	return &p
}
