package proto

type LoginRequest struct {
	Name     string
	Password string
}

type LoginResponse struct {
	Code      int
	AuthToken string
}

type RegisterRequest struct {
	Name     string
	Password string
}

type RegisterResponse struct {
	Code      int
	AuthToken string
}

type GetUserInfoResponse struct {
	Code   int
	UserId int
}
