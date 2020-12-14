package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/go-querystring/query"
	"github.com/joho/godotenv"
)

// LastFM holds methods for interacting with LastFM
type LastFM struct {
	baseURL string
	apiKey  string
}

func (lfm LastFM) buildParams(method string, params interface{}) string {
	defaultParams := struct {
		APIKey string `url:"api_key"`
		Format string `url:"format"`
		Method string `url:"method"`
	}{
		APIKey: lfm.apiKey,
		Format: "json",
		Method: method,
	}

	fmt.Println(defaultParams)

	defaultValues, _ := query.Values(defaultParams)
	paramValues, _ := query.Values(params)

	if len(paramValues) > 0 {
		return fmt.Sprintf("%s&%s", defaultValues.Encode(), paramValues.Encode())
	}

	return defaultValues.Encode()
}

// MakeRequest calls the lastfm api with the given parameters
func (lfm LastFM) MakeRequest(method string, params interface{}) *http.Response {
	queryparams := lfm.buildParams(method, params)

	log.Print(queryparams)

	resp, err := http.Get(lfm.baseURL + "?" + queryparams)

	if err != nil {
		log.Println("Error! ", err)
	}

	return resp
}

// ParseResponse parses a JSON respone from the last.fm api
func (lfm LastFM) ParseResponse(response *http.Response, output interface{}) *ErrorResponse {
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
func (lfm LastFM) UserInfo(username string) (*ErrorResponse, *UserInfoResponse) {
	params := UserInfoParams{
		Username: username,
	}

	userInfo := &UserInfoResponse{}

	response := lfm.MakeRequest("user.getInfo", params)

	err := lfm.ParseResponse(response, userInfo)

	return err, userInfo
}

// CreateService creates an instance of the lastfm service object
func CreateService() *LastFM {
	godotenv.Load()

	apiKey := os.Getenv("LAST_FM_API_KEY")

	service := &LastFM{
		baseURL: "http://ws.audioscrobbler.com/2.0/",
		apiKey:  apiKey,
	}

	return service
}
