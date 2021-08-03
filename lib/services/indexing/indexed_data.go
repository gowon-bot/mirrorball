package indexing

import (
	"context"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/helpers/inputparser"
	"github.com/jivison/gowon-indexer/lib/meta"
)

// GetArtist gets and optionally creates an indexed artist
func (i Indexing) GetArtist(artistInput model.ArtistInput, create bool) (*db.Artist, error) {
	artist := new(db.Artist)

	parser := inputparser.CreateParser(db.Db.Model(artist))

	query := parser.ParseArtistInput(artistInput, inputparser.InputSettings{}).GetQuery()

	err := query.Limit(1).Select()

	if err != nil && create && artistInput.Name != nil {
		artist = &db.Artist{
			Name: *artistInput.Name,
		}

		db.Db.Model(artist).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("artist")
	}

	return artist, nil
}

func (i Indexing) GetArtists(artistInputs []*model.ArtistInput, tagInput *model.TagInput, ctx context.Context) ([]db.Artist, error) {
	var artists []db.Artist

	parser := inputparser.CreateParser(db.Db.Model(&artists))

	if len(artistInputs) != 0 {
		parser.ParseArtistInputs(artistInputs, inputparser.InputSettings{})
	}

	if tagInput != nil {
		parser.ParseTagInputForArtist(tagInput, inputparser.InputSettings{})
	}

	query := parser.GetQuery()

	if meta.HasRequestedField(ctx, "tags") {
		query.Relation("Tags")
	}

	err := query.Select()

	if err != nil {
		return nil, customerrors.DatabaseUnknownError()
	}

	return artists, nil
}

// GetAlbum returns (and optionally creates) an album from the database
func (i Indexing) GetAlbum(albumInput model.AlbumInput, create bool) (*db.Album, error) {
	album := new(db.Album)

	parser := inputparser.CreateParser(db.Db.Model(album).Relation("Artist"))

	query := parser.ParseAlbumInput(albumInput, inputparser.InputSettings{}).GetQuery()

	err := query.Limit(1).Select()

	if err != nil && create && albumInput.Name != nil && albumInput.SafeGetArtistName() != nil {
		artist, _ := i.GetArtist(*albumInput.Artist, true)

		album = &db.Album{
			Name: *albumInput.Name,

			ArtistID: artist.ID,
			Artist:   artist,
		}

		db.Db.Model(album).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("album")
	}

	return album, nil
}

// GetTrack returns (and optionally creates) a track from the database
func (i Indexing) GetTrack(trackInput model.TrackInput, create bool) (*db.Track, error) {
	track := new(db.Track)

	parser := inputparser.CreateParser(db.Db.Model(track).Relation("Artist").Relation("Album"))

	query := parser.ParseTrackInput(trackInput, inputparser.InputSettings{}).GetQuery()

	err := query.Limit(1).Select()

	if err != nil && create && trackInput.Name != nil && trackInput.SafeGetArtistName() != nil {
		track = i.SaveTrack(*trackInput.Name, *trackInput.SafeGetArtistName(), trackInput.SafeGetAlbumName())
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("track")
	}

	return track, nil
}

// GetTracks returns a list of tracks from the database
func (i Indexing) GetTracks(trackInput model.TrackInput, create bool) ([]db.Track, error) {
	var tracks []db.Track

	parser := inputparser.CreateParser(db.Db.Model(&tracks).Relation("Artist").Relation("Album"))

	err := parser.ParseTrackInput(trackInput, &inputparser.InputSettings{}).GetQuery().Select()

	if err != nil {
		return nil, customerrors.EntityDoesntExistError("track")
	}

	return tracks, nil
}

// GetArtistCount gets and optionally creates an artist count
func (i Indexing) GetArtistCount(artist *db.Artist, user *db.User, create bool) (*db.ArtistCount, error) {
	artistCount := new(db.ArtistCount)

	err := db.Db.Model(artistCount).Where("user_id=?", user.ID).Where("artist_id=?", artist.ID).Limit(1).Select()

	if err != nil && create {
		artistCount = &db.ArtistCount{
			UserID: user.ID,
			User:   user,

			ArtistID: artist.ID,
			Artist:   artist,
		}

		db.Db.Model(artistCount).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("artist count")
	}

	return artistCount, nil
}

// GetAlbumCount gets and optionally creates an album count
func (i Indexing) GetAlbumCount(album *db.Album, user *db.User, create bool) (*db.AlbumCount, error) {
	albumCount := new(db.AlbumCount)

	err := db.Db.Model(albumCount).Where("user_id=?", user.ID).Where("album_id=?", album.ID).Limit(1).Select()

	if err != nil && create {
		albumCount = &db.AlbumCount{
			UserID: user.ID,
			User:   user,

			AlbumID: album.ID,
			Album:   album,
		}

		db.Db.Model(albumCount).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("album count")
	}

	return albumCount, nil
}

// GetTrackCount gets and optionally creates an track count
func (i Indexing) GetTrackCount(track *db.Track, user *db.User, create bool) (*db.TrackCount, error) {
	trackCount := new(db.TrackCount)

	err := db.Db.Model(trackCount).Where("user_id=?", user.ID).Where("track_id=?", track.ID).Limit(1).Select()

	if err != nil && create {
		trackCount = &db.TrackCount{
			UserID: user.ID,
			User:   user,

			TrackID: track.ID,
			Track:   track,
		}

		db.Db.Model(trackCount).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("track count")
	}

	return trackCount, nil
}

// SaveTrack saves a track in the database
func (i Indexing) SaveTrack(trackName, artistName string, albumName *string) *db.Track {

	artist, _ := i.GetArtist(model.ArtistInput{Name: &artistName}, true)
	var album *db.Album = nil

	if albumName != nil {
		album, _ = i.GetAlbum(model.AlbumInput{
			Name:   albumName,
			Artist: &model.ArtistInput{Name: &artistName},
		}, true)
	}

	track := &db.Track{
		Name: trackName,
	}

	if album != nil {
		track.Album = album
		track.AlbumID = &album.ID
	}

	if artist != nil {
		track.Artist = artist
		track.ArtistID = artist.ID
	}

	db.Db.Model(track).Insert()

	return track
}
