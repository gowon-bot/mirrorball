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
		Privacy:   PresentPrivacy(user.Privacy),
	}
}

func PresentUsers(users []*db.User) []*model.User {
	var builtUsers []*model.User

	for _, user := range users {
		builtUsers = append(builtUsers, PresentUser(user))
	}

	return builtUsers
}

func PresentPrivacy(privacy int64) *model.Privacy {
	var privacyString string

	switch privacy {
	case 1:
		privacyString = "PRIVATE"
	case 2:
		privacyString = "DISCORD"
	case 3:
		privacyString = "FMUSERNAME"
	case 4:
		privacyString = "BOTH"
	case 5:
		privacyString = "UNSET"

	}

	return (*model.Privacy)(&privacyString)
}
