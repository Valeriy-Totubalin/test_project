package request

type SignUp struct {
	Login    string `json:"login" binding:"required,max=60"`
	Password string `json:"password" binding:"required,max=255,min=8"`
}

type SignIn struct {
	Login    string `json:"login" binding:"required,max=60"`
	Password string `json:"password" binding:"required,max=255,min=8"`
}
