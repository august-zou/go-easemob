package easemob

import "testing"

func Test_GetAccessToken(t *testing.T) {
	accessToken := GetAccessToken()
	t.Log(accessToken.Token)
	t.Log(accessToken.ExpiresIn)
	t.Log(accessToken.Application)
}
