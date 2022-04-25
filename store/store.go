package store

import "userDetailsApi/types"

type DB struct {
}

func New() *DB {
	return &DB{}
}

func (d DB) InsertNewUser(u *types.User) error {

	types.UsersList = append(types.UsersList, *u)

	return nil
}

func (d DB) GetAllUsers() []types.User {

	return types.UsersList
}

func (d DB) GetUserByName(u *types.User, name string) types.User {

	var p int
	for p, v := range types.UsersList {
		if v.Name == name {
			return types.UsersList[p]
		}
	}

	return types.UsersList[p-1]
}
