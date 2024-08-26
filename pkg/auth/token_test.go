package auth

import "testing"

func TestGenToken(t *testing.T) {
	token, err := GenToken(&UserInfo{
		UID:  "111",
		Name: "admin",
	})
	t.Log(token, err)
}
