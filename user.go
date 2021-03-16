package todo

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
