package dbhelpers

import (
	"strconv"
	"strings"

	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
)

func InsertUniqueArtistTags(artistTags []db.ArtistTag) error {
	values := []string{}

	for _, artistTag := range artistTags {
		value := "(" + strconv.Itoa(int(artistTag.ArtistID)) + "," + strconv.Itoa(int(artistTag.TagID)) + ")"

		values = append(values, value)
	}

	_, err := db.Db.Model((*db.ArtistTag)(nil)).Exec("INSERT INTO ?TableName (artist_id, tag_id) VALUES " + strings.Join(values, ",") + " EXCEPT SELECT artist_id, tag_id FROM artist_tags")

	if err != nil {
		return customerrors.DatabaseUnknownError()
	}

	return nil
}
