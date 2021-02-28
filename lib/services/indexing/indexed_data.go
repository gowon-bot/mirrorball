package indexing

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// GetArtist gets and optionally creates an indexed artist
func (i Indexing) GetArtist(artistInput model.ArtistInput, create bool) (*db.Artist, error) {
	artist := new(db.Artist)

	query := db.Db.Model(artist)

	if artistInput.Name != nil {
		query = query.Where("name ILIKE ?", artistInput.Name)
	}

	err := query.Limit(1).Select()

	if err != nil && create == true && artistInput.Name != nil {
		artist = &db.Artist{
			Name: *artistInput.Name,
		}

		db.Db.Model(artist).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("artist")
	}

	return artist, nil
}

// GetArtistCount gets and optionally creates an artist count
func (i Indexing) GetArtistCount(artist *db.Artist, user *db.User, create bool) (*db.ArtistCount, error) {
	artistCount := new(db.ArtistCount)

	err := db.Db.Model(artistCount).Where("user_id=?", user.ID).Where("artist_id=?", artist.ID).Limit(1).Select()

	if err != nil && create == true {
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
