package auth

const (
	Authorization = "authorization"
	UserId        = "user-id"
	UserName      = "user-name"
)

type UserInfo struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}
