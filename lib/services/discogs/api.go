package discogs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (d Discogs) MakeRequest(target interface{}, toPath string, variables ...interface{}) error {
	req, _ := http.NewRequest("GET", d.baseURL+fmt.Sprintf(toPath, variables...), nil)

	req.Header.Set("Authorization", fmt.Sprintf("Discogs key=%s, secret=%s", d.apiKey, d.secret))

	resp, err := d.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)

	return err
}

func (d Discogs) GetUsersCollection(user string) (Collection, error) {
	var collection Collection

	err := d.MakeRequest(&collection, "/users/%s/collection/releases/%s?per_page=1", user, "0")

	return collection, err
}
