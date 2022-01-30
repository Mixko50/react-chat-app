package event

type UserConnectType struct {
	Username string `json:"username"`
	Online   int    `json:"online"`
}
