package rateyourmusic

import (
	"encoding/csv"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	dbhelpers "github.com/jivison/gowon-indexer/lib/helpers/database"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/generic"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
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
	AllAlbums        []indexing.AlbumToConvert
}

// var asianCharacters = `[\p{Hangul}\p{Han}\p{Katakana}\p{Hiragana}]`

// var containsAsianCharacters = regexp.MustCompile(asianCharacters + "+")
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
		if strings.HasPrefix(record[0], "RYM") || record[Rating] == "0" {
			continue
		}

		artistName := unescape(combineNames(record[FirstName], record[LastName]))
		artistNameLocalized := unescape(combineNames(record[FirstNameLocalized], record[LastNameLocalized]))

		title := unescape(record[Title])

		row := RawRateYourMusicRating{}
		if len(artistNameLocalized) > 0 {
			row.ArtistName = artistNameLocalized
			row.ArtistNativeName = &artistName
		} else {
			nativeArtistNames := strings.TrimSpace(removeLocalizedArtistNames(artistName))
			localizedArtistNames := strings.TrimSpace(removeNativeArtistNames(artistName))

			row.ArtistName = localizedArtistNames
			row.ArtistNativeName = &nativeArtistNames
		}

		_, err = rym.indexingService.GetAlbum(model.AlbumInput{
			Artist: &model.ArtistInput{Name: &row.ArtistName},
			Name:   &title,
		}, true)

		if err != nil {
			return nil, err
		}

		if row.ArtistNativeName != nil {
			_, err = rym.indexingService.GetAlbum(model.AlbumInput{
				Artist: &model.ArtistInput{Name: row.ArtistNativeName},
				Name:   &title,
			}, true)

			if err != nil {
				return nil, err
			}
		}

		row.RYMID = record[RYMID]
		row.Title = title
		row.Rating, _ = strconv.Atoi(record[Rating])
		row.ReleaseYear, _ = strconv.Atoi(record[ReleaseYear])

		albums, _ := rym.generateRawAlbumCombinations(record, row)

		row.AllAlbums = albums

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

func (rym RateYourMusic) generateRawAlbumCombinations(record []string, row RawRateYourMusicRating) ([]indexing.AlbumToConvert, error) {
	releaseTitle := unescape(record[Title])
	artistName := unescape(record[LastName])

	artistsToCheck := []indexing.AlbumToConvert{{ArtistName: row.ArtistName, AlbumName: row.Title}}

	if row.ArtistNativeName != nil {
		artistsToCheck = append(artistsToCheck, indexing.AlbumToConvert{ArtistName: *row.ArtistNativeName, AlbumName: row.Title})
	}

	var individualArtistNames []string

	splitOnAnds := regexp.MustCompile(`( & | ,)`).Split(artistName, -1)

	for _, split := range splitOnAnds {
		trimmedSplit := strings.TrimSpace(split)

		if localization, ok := parseLocalization(trimmedSplit); ok {
			individualArtistNames = append(individualArtistNames, localization.Localized, localization.Native)
		} else {
			individualArtistNames = append(individualArtistNames, trimmedSplit)
		}
	}

	for _, permutation := range helpers.Combinations(individualArtistNames) {
		artistsToCheck = append(artistsToCheck, indexing.AlbumToConvert{ArtistName: joinArtists(permutation), AlbumName: releaseTitle})
	}

	filteredArtists, err := rym.filterOutNonExistantCombinations(artistsToCheck)

	if err != nil {
		return nil, err
	}

	return filteredArtists, nil
}

func (rym RateYourMusic) filterOutNonExistantCombinations(combinationsToCheck []indexing.AlbumToConvert) ([]indexing.AlbumToConvert, error) {
	var combos []indexing.AlbumToConvert

	searchableAlbums := rym.indexingService.GenerateAlbumsToSearch(combinationsToCheck)

	databaseAlbums, err := dbhelpers.SelectAlbumsWhereInMany(searchableAlbums, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, album := range databaseAlbums {
		combos = append(combos, indexing.AlbumToConvert{ArtistName: album.Artist.Name, AlbumName: album.Name})
	}

	return combos, nil
}

func joinArtists(artists []string) string {
	if len(artists) == 1 {
		return artists[0]
	} else if len(artists) == 2 {
		return artists[0] + " & " + artists[1]
	} else {
		secondLastIndex := len(artists) - 1
		return strings.Join(artists[1:secondLastIndex], ", ") + artists[len(artists)-1]
	}
}

func unescape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&amp;", "&"), "&#34;", `"`)
}

// regex functions
var nativeArtistNamesRegex = regexp.MustCompile(`\[([^\]]+)\]`)
var localizedArtistNamesRegex = regexp.MustCompile(`(^|[&\,] ?)[^&,\[]+ \[([^\]]+)\]`)

type Localization = struct {
	Localized string
	Native    string
}

func parseLocalization(name string) (Localization, bool) {
	localization := artistLocalization.FindAllStringSubmatch(name, 1)

	if len(localization) > 0 && len(localization[0]) == 3 {
		return Localization{
			Localized: localization[0][2],
			Native:    localization[0][1],
		}, true
	}

	return Localization{}, false
}

func removeLocalizedArtistNames(artistName string) string {
	return nativeArtistNamesRegex.ReplaceAllString(artistName, "")
}

func removeNativeArtistNames(artistName string) string {
	return localizedArtistNamesRegex.ReplaceAllString(artistName, "${1}${2}")
}
