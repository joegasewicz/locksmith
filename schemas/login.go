package schemas

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
