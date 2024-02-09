package schemas

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
	Avatar   string `json:"avatar"`
}

type UserOmitPassword struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	Avatar   string `json:"avatar"`
}

type UserWithToken struct {
	User  UserOmitPassword `json:"user"`
	Token string           `json:"token"`
}
