package meta

import (
	"context"
	"net/http"
	"os"
)

type ContextKey string

const ContextDiscordIDKey ContextKey = "doughnutDiscordID"

func EnforcePassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")
		passsword := os.Getenv("MIRRORBALL_PASSWORD")

		if os.Getenv("ENVIRONMENT") != "development" && authorization != passsword {
			http.Error(w, `{ "message": "Incorrect password" }`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func DoughnutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		discordID := r.Header.Get("Doughnut-Discord-Id")

		ctx := context.WithValue(r.Context(), ContextDiscordIDKey, discordID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
