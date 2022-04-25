package types

import (
	"encoding/json"
	"io"
)

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         int64  `json:"age"`
	Designation string `json:"designation"`
}

func (u *User) DecodeFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) EncodeToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)

}

var UsersList = []User{
	{
		Name:        "Eric",
		Email:       "eric@gmail.com",
		Age:         31,
		Designation: "Analyst",
	},
	{
		Name:        "Chaos",
		Email:       "chaos@gmail.com",
		Age:         30,
		Designation: "TeamLead",
	},
	{
		Name:        "Henry",
		Email:       "henry@gmail.com",
		Age:         42,
		Designation: "Manager",
	},
	{
		Name:        "Grey",
		Email:       "grey@gmail.com",
		Age:         25,
		Designation: "Developer",
	},
}
