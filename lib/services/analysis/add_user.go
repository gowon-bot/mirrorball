package analysis

import (
	"github.com/jinzhu/copier"
	"github.com/jivison/gowon-indexer/lib/constants"
	"github.com/jivison/gowon-indexer/lib/db"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/database"
)

func (a Analysis) AddUserToPlays(plays []db.Play) ([]db.Play, error) {
	var userList []int64

	for _, play := range plays {
		userList = append(userList, play.UserID)
	}

	usersMap, err := a.generateUsersMap(userList)

	if err != nil {
		return plays, err
	}

	var newPlays []db.Play

	for _, play := range plays {
		user := usersMap[play.UserID]

		copiedUser := db.User{}

		copier.Copy(&copiedUser, &user)

		play.User = &copiedUser

		newPlays = append(newPlays, play)
	}

	return newPlays, nil
}

type UsersMap = map[int64]db.User

func (a Analysis) generateUsersMap(userList []int64) (UsersMap, error) {
	usersMap := make(UsersMap)

	users, err := helpers.SelectUsersWhereInMany(userList, constants.ChunkSize)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		usersMap[user.ID] = user
	}

	return usersMap, nil
}
