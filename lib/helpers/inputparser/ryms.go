package inputparser

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func (p InputParser) ParseRatingsSettings(settingsInput *model.RatingsSettings) *InputParser {
	if settingsInput == nil {
		return &p
	}

	if settingsInput.User != nil {
		p.ParseUserInput(*settingsInput.User, InputSettings{})
	}

	if settingsInput.Album != nil {
		p.ParseAlbumInputForRatings(*settingsInput.Album)
	}

	if settingsInput.PageInput != nil {
		p.ParsePageInput(settingsInput.PageInput)
	}

	return &p
}

func (p InputParser) ParseAlbumInputForRatings(albumInput model.AlbumInput) *InputParser {
	subquery := db.Db.Model(&db.RateYourMusicAlbumAlbum{}).
		Column("rate_your_music_album_id").
		Relation("Album._").
		Relation("Album.Artist._")

	subquery = CreateParser(subquery).
		ParseAlbumInput(albumInput, InputSettings{ArtistPath: "album__artist"}).
		GetQuery()

	p.query.Where("?TableAlias.rate_your_music_album_id IN (?)", subquery)

	return &p
}
