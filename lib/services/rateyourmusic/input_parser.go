package rateyourmusic

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

func ParseRatingsSettings(query *orm.Query, settings *model.RatingsSettings) *orm.Query {
	if settings == nil {
		return query
	}

	if settings.User != nil {
		query = ParseUserInput(query, *settings.User)
	}

	if settings.Album != nil {
		query = ParseAlbumInput(query, *settings.Album)
	}

	if settings.PageInput != nil {
		query = indexing.ParsePageInput(query, settings.PageInput)
	}

	return query
}

func ParseUserInput(query *orm.Query, userInput model.UserInput) *orm.Query {
	query = query.Relation("User")

	if userInput.DiscordID != nil {
		query.Where("discord_id = ?", userInput.DiscordID)
	}

	if userInput.LastFMUsername != nil {
		query.Where("user_type = 'Lastfm'").Where("username = ?", userInput.LastFMUsername)
	}

	if userInput.WavyUsername != nil {
		query.Where("user_type = 'Wavy'").Where("username = ?", userInput.WavyUsername)
	}

	return query
}

func ParseAlbumInput(query *orm.Query, albumInput model.AlbumInput) *orm.Query {
	subquery := db.Db.Model(&db.RateYourMusicAlbumAlbum{}).Column("rate_your_music_album_id").Relation("Album._")

	if albumInput.Name != nil {
		subquery = subquery.Where("album.name ILIKE ?", albumInput.Name)
	}

	if albumInput.Artist != nil {
		subquery = subquery.Relation("Album.Artist._").Where("album__artist.name ILIKE ?", albumInput.Artist.Name)
	}

	query = query.Where("?TableAlias.rate_your_music_album_id IN (?)", subquery)

	return query
}
