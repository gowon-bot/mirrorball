package discogs

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Discogs struct {
	baseURL string
	apiKey  string
	secret  string
	client  *http.Client
}

func CreateService() *Discogs {
	godotenv.Load()

	apiKey := os.Getenv("DISCOGS_KEY")
	secret := os.Getenv("DISCOGS_SECRET")

	service := &Discogs{
		apiKey:  apiKey,
		baseURL: `https://api.discogs.com/`,
		client:  &http.Client{},
		secret:  secret,
	}

	return service
}
