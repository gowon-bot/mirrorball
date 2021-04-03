package analysis

import (
	"github.com/go-pg/pg/v10"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// WhoKnowsArtist returns a list of who has listened to an artist
func (a Analysis) WhoKnowsArtist(artist *db.Artist, settings *model.WhoKnowsSettings) ([]db.ArtistCount, error) {
	var whoKnows []db.ArtistCount

	query := db.Db.Model(&whoKnows).
		Relation("User").
		Where("artist_id = ?", artist.ID).
		Order("playcount desc", "username desc")

	err := ParseWhoKnowsSettings(query, settings).Select()

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

	err := ParseWhoKnowsSettings(query, settings).Select()

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
	UserType  string
	DiscordID string
}

// WhoKnowsTrack returns a list of who has listened to an album
func (a Analysis) WhoKnowsTrack(tracks []db.Track, settings *model.WhoKnowsSettings) ([]*model.WhoKnowsRow, error) {
	var whoKnows []WhoKnowsTrackRow
	var trackIDs []int64
	var whoKnowsTracks []*model.WhoKnowsRow

	for _, track := range tracks {
		trackIDs = append(trackIDs, track.ID)
	}

	trackCounts := db.Db.Model((*db.TrackCount)(nil)).
		Relation("User._").
		Join("JOIN tracks ON tracks.id = track_id").
		Column("name", "artist_id", "tracks.name", "username", "discord_id", "user_type").
		ColumnExpr("sum(playcount) as total_playcount").
		ColumnExpr("track_count.user_id as user_id").
		Where("track_id IN (?)", pg.In(trackIDs)).
		Group("name", "artist_id", "track_count.user_id", "username", "discord_id", "user_type")

	trackCounts = ParseWhoKnowsSettings(trackCounts, settings)

	err := db.Db.Model().
		TableExpr("(?) as track_counts", trackCounts).
		ColumnExpr("total_playcount as playcount").
		ColumnExpr(`username as username`).
		ColumnExpr(`user_type as user_type`).
		ColumnExpr(`discord_id as discord_id`).
		Order("total_playcount desc", "username desc").
		Select(&whoKnows)

	if err != nil {
		return whoKnowsTracks, err
	}

	for _, row := range whoKnows {
		whoKnowsTracks = append(whoKnowsTracks, &model.WhoKnowsRow{
			Playcount: int(row.Playcount),
			User: &model.User{
				ID:        int(row.UserID),
				Username:  row.Username,
				UserType:  (*model.UserType)(&row.UserType),
				DiscordID: row.DiscordID,
			},
		})
	}

	return whoKnowsTracks, nil
}
