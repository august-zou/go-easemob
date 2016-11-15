package easemob

import (
	"encoding/json"
	"time"
)

//User easemob user
type User struct {
	UUID      string    `json:"uuid"`
	Type      string    `json:"type"`
	Created   time.Time `json:"crreated"`
	Modified  time.Time `json:"modified"`
	UserName  string    `json:"username"`
	Activated bool      `json:"activated"`
}

//GetSingleUser get single user
func GetSingleUser(id string) interface{} {
	url := BaseURL + "/users/" + id
	header := map[string]string{"Authorization": "Bearer " + GetAccessToken().Token}
	data := HTTPGet(url, header)

	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {

	}

	return f
}
