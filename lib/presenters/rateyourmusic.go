package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

func PresentRateYourMusicAlbum(album db.RateYourMusicAlbum) model.RateYourMusicAlbum {
	return model.RateYourMusicAlbum{
		RateYourMusicID:  album.RateYourMusicID,
		ReleaseYear:      album.ReleaseYear,
		Title:            album.Title,
		ArtistName:       album.ArtistName,
		ArtistNativeName: album.ArtistNativeName,
	}
}

func PresentRating(rating db.Rating) model.Rating {
	builtRating := model.Rating{
		Rating: rating.Rating,
	}

	if rating.RateYourMusicAlbum != nil {
		builtAlbum := PresentRateYourMusicAlbum(*rating.RateYourMusicAlbum)

		builtRating.RateYourMusicAlbum = &builtAlbum
	}

	return builtRating
}

func PresentRatings(ratings []db.Rating) []*model.Rating {
	var builtRatings []*model.Rating

	for _, rating := range ratings {
		builtRating := PresentRating(rating)
		builtRatings = append(builtRatings, &builtRating)
	}

	return builtRatings
}

func PresentRateYourMusicArtist(album db.RateYourMusicAlbum) *model.RateYourMusicArtist {
	builtArtist := model.RateYourMusicArtist{}

	if album.ArtistName != "" {
		builtArtist.ArtistName = album.ArtistName
	}

	if album.ArtistNativeName != nil {
		builtArtist.ArtistNativeName = album.ArtistNativeName
	}

	return &builtArtist
}
