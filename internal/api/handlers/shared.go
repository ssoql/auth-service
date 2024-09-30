package handlers

type UserInput struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenInput struct {
	Token string `json:"token" binding:"required"`
}
