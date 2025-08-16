package main

import (
	"encoding/json"
	"fmt"
)

// Run with: go run types/010_json_newtype.go

type UserID int

type UID struct{ UserID }

func (u UID) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ ID int `json:"id"` }{ID: int(u.UserID)})
}

func (u *UID) UnmarshalJSON(b []byte) error {
	var tmp struct{ ID int `json:"id"` }
	if err := json.Unmarshal(b, &tmp); err != nil { return err }
	u.UserID = UserID(tmp.ID)
	return nil
}

func main() {
	u := UID{UserID: 42}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	var v UID
	_ = json.Unmarshal(b, &v)
	fmt.Println(v.UserID)
}

