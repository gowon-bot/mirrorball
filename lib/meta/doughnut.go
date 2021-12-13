package meta

import (
	"context"

	"github.com/jivison/gowon-indexer/lib/customerrors"
)

func CheckUserMatches(ctx context.Context, id string) error {
	doughnutDiscordID := ctx.Value(ContextDiscordIDKey).(string)

	if doughnutDiscordID != id {
		return customerrors.NotAuthorized()
	}

	return nil
}

func CheckNoUser(ctx context.Context) error {
	doughnutDiscordID := ctx.Value(ContextDiscordIDKey).(string)

	if doughnutDiscordID != "" {
		return customerrors.NotAuthorized()
	}

	return nil
}
