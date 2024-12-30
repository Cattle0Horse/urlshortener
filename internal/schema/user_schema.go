package schema

type LoginRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8,max=20"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
	UserID      uint   `json:"user_id"`
}

type RegisterReqeust struct {
	LoginRequest
	EmailCode string `json:"email_code" validate:"required,len=6"`
}

type ResetPasswordReqeust struct {
	LoginRequest
	EmailCode string `json:"email_code" validate:"required,len=6"`
}

type SendCodeRequest struct {
	Email string `validate:"required,email"`
}
