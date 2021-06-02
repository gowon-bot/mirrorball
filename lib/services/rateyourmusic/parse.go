package rateyourmusic

import (
	"encoding/csv"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/jivison/gowon-indexer/lib/customerrors"
)

// RYM Album, First Name,Last Name,First Name localized, Last Name localized,Title,Release_Date,Rating,Ownership,Purchase Date,Media Type,Review
const (
	RYMID              = 0
	FirstName          = 1
	LastName           = 2
	FirstNameLocalized = 3
	LastNameLocalized  = 4
	Title              = 5
	ReleaseYear        = 6
	Rating             = 7
)

type RawRateYourMusicRating = struct {
	RYMID            string
	Title            string
	ArtistName       string
	ArtistNativeName *string
	Rating           int
	ReleaseYear      int
}

var asianCharacters = `[\p{Hangul}\p{Han}\p{Katakana}\p{Hiragana}]`
var containsAsianCharacters = regexp.MustCompile(asianCharacters + "+")
var artistLocalization = regexp.MustCompile(`([^\[]+) \[([^\]]+)\]`)

func (rym RateYourMusic) ParseRYMSExport(csvString string) ([]RawRateYourMusicRating, error) {
	r := csv.NewReader(strings.NewReader(csvString))
	r.LazyQuotes = true

	var albumRatings []RawRateYourMusicRating

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, customerrors.CSVParseError()
		}
		// Header
		if strings.HasPrefix(record[0], "RYM") {
			continue
		}

		artistName := combineNames(record[FirstName], record[LastName])
		artistNameLocalized := combineNames(record[FirstNameLocalized], record[LastNameLocalized])
		localization := artistLocalization.FindAllStringSubmatch(artistName, 1)

		row := RawRateYourMusicRating{}
		row.RYMID = record[RYMID]
		row.Title = record[Title]
		row.Rating, _ = strconv.Atoi(record[Rating])
		row.ReleaseYear, _ = strconv.Atoi(record[ReleaseYear])

		if len(localization) > 0 && len(localization[0]) == 3 {
			row.ArtistName = localization[0][2]
			row.ArtistNativeName = &localization[0][1]
		} else {
			if containsAsianCharacters.FindStringIndex(artistName) != nil && artistNameLocalized != "" {
				row.ArtistName = artistNameLocalized
				row.ArtistNativeName = &artistName
			} else {
				row.ArtistName = artistName
			}
		}

		albumRatings = append(albumRatings, row)
	}

	return albumRatings, nil
}

func combineNames(firstName string, lastName string) string {
	name := firstName

	if name != "" && lastName != "" {
		name += " " + lastName
	} else if lastName != "" {
		name = lastName
	}

	return name
}
