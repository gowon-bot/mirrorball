package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"

	"github.com/google/go-querystring/query"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/api"
	"github.com/joho/godotenv"
)

// API holds methods for interacting with the Last.fm API
type API struct {
	baseURL   string
	apiKey    string
	apiSecret string
}

func (lfm API) buildParams(method string, params interface{}) string {
	defaultParams := struct {
		APIKey string `url:"api_key"`
		Format string `url:"format"`
		Method string `url:"method"`
	}{
		APIKey: lfm.apiKey,
		Format: "json",
		Method: method,
	}

	defaultValues, _ := query.Values(defaultParams)
	paramValues, _ := query.Values(params)

	if paramValues.Get("api_sig") == ApiSigReplace {
		return lfm.authenticatedParams(defaultValues, paramValues).Encode()
	}

	if len(paramValues) > 0 {
		return fmt.Sprintf("%s&%s", defaultValues.Encode(), paramValues.Encode())
	}

	return defaultValues.Encode()
}

// MakeRequest calls the lastfm api with the given parameters
func (lfm API) MakeRequest(method string, params interface{}) *http.Response {
	queryparams := lfm.buildParams(method, params)

	log.Printf("Making call to last.fm with parameters: %s", queryparams)

	resp, err := http.Get(lfm.baseURL + "?" + queryparams)

	if err != nil {
		log.Println("Error! ", err)
	}

	return resp
}

// ParseResponse parses a JSON respone from the last.fm api
func (lfm API) ParseResponse(response *http.Response, output interface{}) *ErrorResponse {
	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	errorResponse := &ErrorResponse{}

	json.Unmarshal(responseBody, output)
	json.Unmarshal(responseBody, errorResponse)

	if errorResponse.Error != 0 {
		return errorResponse
	}

	return nil
}

// UserInfo fetches a user's info from the last.fm API
func (lfm API) UserInfo(requestable Requestable) (*ErrorResponse, *UserInfoResponse) {
	params := UserInfoParams{
		Username: requestable,
	}

	userInfo := &UserInfoResponse{}

	response := lfm.MakeRequest("user.getInfo", params)

	err := lfm.ParseResponse(response, userInfo)

	return err, userInfo
}

func (lfm API) ArtistInfo(artist string) (*ErrorResponse, *ArtistInfoResponse) {
	params := ArtistInfoParams{
		Artist: artist,
	}

	artistInfo := &ArtistInfoResponse{}

	response := lfm.MakeRequest("artist.getInfo", params)

	err := lfm.ParseResponse(response, artistInfo)

	return err, artistInfo
}

// RecentTracks fetches a user's recent tracks from the last.fm API
func (lfm API) RecentTracks(params RecentTracksParams) (*ErrorResponse, *RecentTracksResponse) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 1
	}
	if params.Period == "" {
		params.Period = "overall"
	}

	recentTracks := &RecentTracksResponse{}

	response := lfm.MakeRequest("user.getRecentTracks", params)

	err := lfm.ParseResponse(response, recentTracks)

	return err, recentTracks
}

// TopArtists fetches a user's top artists from the last.fm API
func (lfm API) TopArtists(params TopEntityParams) (*ErrorResponse, *TopArtistsResponse) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 1
	}
	if params.Period == "" {
		params.Period = "overall"
	}

	topArtists := &TopArtistsResponse{}

	response := lfm.MakeRequest("user.getTopArtists", params)

	err := lfm.ParseResponse(response, topArtists)

	return err, topArtists
}

// TopAlbums fetches a user's top albums from the last.fm API
func (lfm API) TopAlbums(params TopEntityParams) (*ErrorResponse, *TopAlbumsResponse) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 1
	}
	if params.Period == "" {
		params.Period = "overall"
	}

	topAlbums := &TopAlbumsResponse{}

	response := lfm.MakeRequest("user.getTopAlbums", params)

	err := lfm.ParseResponse(response, topAlbums)

	return err, topAlbums
}

// TopTracks fetches a user's top tracks from the last.fm API
func (lfm API) TopTracks(params TopEntityParams) (*ErrorResponse, *TopTracksResponse) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 1
	}
	if params.Period == "" {
		params.Period = "overall"
	}

	topTracks := &TopTracksResponse{}

	response := lfm.MakeRequest("user.getTopTracks", params)

	err := lfm.ParseResponse(response, topTracks)

	return err, topTracks
}

// ValidateUser validates that a given username exists in last.fm
func (lfm API) ValidateUser(username string) bool {
	err, _ := lfm.UserInfo(Requestable{Username: username})

	return err == nil
}

// LastScrobbledTimestamp returns the timestamp of the last scrobbled track
func (lfm API) LastScrobbledTimestamp(requestable Requestable) time.Time {
	err, recentTracks := lfm.RecentTracks(RecentTracksParams{
		Limit:    1,
		Username: requestable,
	})

	if err != nil || len(recentTracks.RecentTracks.Tracks) == 0 {
		return time.Now()
	}

	var lastTrack RecentTrack

	if (len(recentTracks.RecentTracks.Tracks)) == 2 {
		lastTrack = recentTracks.RecentTracks.Tracks[1]
	} else {
		lastTrack = recentTracks.RecentTracks.Tracks[0]
	}

	timestamp, parseErr := helpers.ParseUnix(lastTrack.Timestamp.UTS)

	if parseErr != nil {
		return time.Now()
	}

	return timestamp
}

func (api API) authenticatedParams(defaultParams, params url.Values) url.Values {
	valuesMap := make(url.Values)
	var values [][]string

	params.Del("api_sig")

	for key, value := range defaultParams {
		values = append(values, []string{key, value[0]})
		valuesMap.Add(key, value[0])
	}

	for key, value := range params {
		values = append(values, []string{key, value[0]})
		valuesMap.Add(key, value[0])
	}

	signature := ""

	sort.Sort(byAlphabeticalKey(values))

	for _, value := range values {
		if value[0] == "format" {
			continue
		}

		signature += value[0] + value[1]
	}

	signature += api.apiSecret

	hashedSignature := md5.Sum([]byte(signature))

	valuesMap.Add("api_sig", hex.EncodeToString(hashedSignature[:]))

	return valuesMap
}

type byAlphabeticalKey [][]string

func (s byAlphabeticalKey) Len() int {
	return len(s)
}
func (s byAlphabeticalKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAlphabeticalKey) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

// CreateAPIService creates an instance of the lastfm api service object
func CreateAPIService() *API {
	godotenv.Load()

	apiKey := os.Getenv("LAST_FM_API_KEY")
	apiSecret := os.Getenv("LAST_FM_API_SECRET")

	service := &API{
		baseURL:   "http://ws.audioscrobbler.com/2.0/",
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}

	return service
}
