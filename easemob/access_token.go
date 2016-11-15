package easemob

import (
	"encoding/json"
	"time"
)

//AccessToken accessToken
type AccessToken struct {
	Token       string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Application string `json:"application"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (a *AccessToken) setExpiresAt(t time.Time) {
	a.ExpiresAt = t.Unix() + a.ExpiresIn
}

func (a AccessToken) hasExpires() bool {
	return a.ExpiresAt > time.Now().Unix()
}

var accessToken *AccessToken

//GetAccessToken get valid accessToken
func GetAccessToken() *AccessToken {
	if accessToken != nil && accessToken.hasExpires() {
		return accessToken
	}
	url := BaseURL + "/token"
	values := map[string]string{"grant_type": "client_credentials", "client_id": "YXA6raOlIC6mEeW-GjVdRj8JRw", "client_secret": "YXA63WSkGzqwiHd0ytFudEcUKCzVQb0"}
	jsonValue, _ := json.Marshal(values)

	now := time.Now()

	data := HTTPPost(url, jsonValue)
	json.Unmarshal(data, &accessToken)

	accessToken.setExpiresAt(now)
	return accessToken
}
