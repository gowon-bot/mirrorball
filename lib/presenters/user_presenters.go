package presenters

import (
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// PresentUser converts a database user into a graphql user
func PresentUser(user *db.User) *model.User {
	return &model.User{
		ID:        int(user.ID),
		Username:  user.Username,
		DiscordID: user.DiscordID,
		UserType:  (*model.UserType)(&user.UserType),
		Privacy:   (*model.Privacy)(&user.Privacy),
	}
}
