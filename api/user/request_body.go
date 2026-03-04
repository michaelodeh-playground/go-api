package user

type CreateUserRequest struct {
	Name  string `json:"name" example:"John Doe" default:"John Doe"`
	Email string `json:"email" example:"john@example.com" default:"john@example.com"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" example:"John Updated" default:"John Updated"`
	Email string `json:"email" example:"updated@example.com" default:"updated@example.com"`
	Age   int    `json:"age" example:"30" default:"30"`
}
