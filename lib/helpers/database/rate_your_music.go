package dbhelpers

import (
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
)

type RatingsMap = map[int64]db.Rating

func UpdateOrCreateManyRatings(ratings []db.Rating, userID int64) (RatingsMap, error) {
	var rateYourMusicAlbumIDs []interface{}

	for _, rating := range ratings {
		rateYourMusicAlbumIDs = append(rateYourMusicAlbumIDs, rating.RateYourMusicAlbumID)
	}

	ratingsMap := make(RatingsMap)
	savedRatings, err := SelectRatingsWhereInMany(rateYourMusicAlbumIDs, userID, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, rating := range savedRatings {
		ratingsMap[rating.RateYourMusicAlbumID] = rating
	}

	var ratingsToCreate []db.Rating
	var ratingsToUpdate []db.Rating

	for _, rating := range ratings {
		if savedRating, ok := ratingsMap[rating.RateYourMusicAlbumID]; !ok {
			ratingsToCreate = append(ratingsToCreate, rating)
		} else {
			newRating := savedRating

			newRating.Rating = rating.Rating

			ratingsToUpdate = append(ratingsToUpdate, newRating)
		}
	}

	createdRatings, err := InsertManyRatings(ratingsToCreate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	updatedRatings, err := UpdateManyRatings(ratingsToUpdate, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, rating := range append(createdRatings, updatedRatings...) {
		ratingsMap[rating.RateYourMusicAlbumID] = rating
	}

	return ratingsMap, nil
}
