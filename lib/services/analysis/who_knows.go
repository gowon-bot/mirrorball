package analysis

import (
	"github.com/go-pg/pg/v10"
	"github.com/jinzhu/copier"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
)

// WhoKnowsArtist returns a list of who has listened to an artist
func (a Analysis) WhoKnowsArtist(artist *db.Artist, settings *model.WhoKnowsSettings) ([]db.ArtistCount, error) {
	if artist == nil {
		return []db.ArtistCount{}, nil
	}

	var whoKnows []db.ArtistCount

	query := db.Db.Model(&whoKnows).
		Relation("User").
		Where("artist_id = ?", artist.ID).
		Order("playcount desc", "username desc")

	err := inputparser.CreateParser(query).ParseWhoKnowsSettings(settings, inputparser.InputSettings{
		UserIDPath: `user"."id`,
	}).GetQuery().Select()

	if err != nil {
		return whoKnows, customerrors.DatabaseUnknownError()
	}

	return whoKnows, nil
}

// WhoKnowsAlbum returns a list of who has listened to an album
func (a Analysis) WhoKnowsAlbum(album *db.Album, settings *model.WhoKnowsSettings) ([]db.AlbumCount, error) {
	var whoKnows []db.AlbumCount

	query := db.Db.Model(&whoKnows).
		Relation("User").
		Where("album_id = ?", album.ID).
		Order("playcount desc", "username desc")

	err := inputparser.CreateParser(query).ParseWhoKnowsSettings(settings, inputparser.InputSettings{
		UserIDPath: `user"."id`,
	}).GetQuery().Select()

	if err != nil {
		return whoKnows, customerrors.DatabaseUnknownError()
	}

	return whoKnows, nil
}

// WhoKnowsTrackRow represents a raw row of a who knows track query
type WhoKnowsTrackRow struct {
	Playcount int64

	UserID    int64
	Username  string
	Privacy   int64
	DiscordID string
}

// WhoKnowsTrack returns a list of who has listened to an album
func (a Analysis) WhoKnowsTrack(tracks []db.Track, settings *model.WhoKnowsSettings) ([]*model.WhoKnowsRow, error) {
	if len(tracks) < 1 {
		return []*model.WhoKnowsRow{}, nil
	}

	limit := -1

	if settings.Limit != nil {
		limit = *settings.Limit
	}

	settings.Limit = nil

	var whoKnows []WhoKnowsTrackRow
	var trackIDs []int64
	var whoKnowsTracks []*model.WhoKnowsRow

	for _, track := range tracks {
		trackIDs = append(trackIDs, track.ID)
	}

	trackCounts := db.Db.Model((*db.TrackCount)(nil)).
		Relation("User._").
		Join("JOIN tracks ON tracks.id = track_id").
		Column("name", "artist_id", "tracks.name", "username", "discord_id", "privacy").
		ColumnExpr("sum(playcount) as total_playcount").
		ColumnExpr("track_count.user_id as user_id").
		Where("track_id IN (?)", pg.In(trackIDs)).
		Group("name", "artist_id", "track_count.user_id", "username", "discord_id", "privacy")

	trackCounts = inputparser.CreateParser(trackCounts).ParseWhoKnowsSettings(settings, inputparser.InputSettings{
		UserIDPath: `user"."id`,
	}).GetQuery()

	query := db.Db.Model().
		TableExpr("(?) as track_counts", trackCounts).
		ColumnExpr("total_playcount as playcount").
		ColumnExpr(`username as username`).
		ColumnExpr(`privacy as privacy`).
		ColumnExpr(`discord_id as discord_id`).
		Order("total_playcount desc", "username desc")

	if limit != -1 {
		query.Limit(limit)
	}

	err := query.Select(&whoKnows)

	if err != nil {
		return whoKnowsTracks, err
	}

	for _, row := range whoKnows {
		copiedRow := WhoKnowsTrackRow{}

		copier.Copy(&copiedRow, row)

		privacyString := db.ConvertPrivacyToString(copiedRow.Privacy)

		whoKnowsTracks = append(whoKnowsTracks, &model.WhoKnowsRow{
			Playcount: int(copiedRow.Playcount),
			User: &model.User{
				ID:       int(copiedRow.UserID),
				Username: copiedRow.Username,
				Privacy:  (*model.Privacy)(&privacyString),

				DiscordID: row.DiscordID,
			},
		})
	}

	return whoKnowsTracks, nil
}
