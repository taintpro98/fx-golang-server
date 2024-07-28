package dto

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type UserPayload struct {
	UserID string `json:"user_id"`
}

func (uc UserPayload) Valid() error {
	return nil
}
