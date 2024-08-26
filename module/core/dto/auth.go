package dto

type LoginRequest struct {
	Phone string `json:"phone"`
}

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type UserPayload struct {
	Sub string `json:"sub"`
}

func (uc UserPayload) Valid() error {
	return nil
}
