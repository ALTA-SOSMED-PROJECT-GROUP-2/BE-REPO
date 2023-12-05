package hendleruser

type UserRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
}

type UserResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Token       string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
