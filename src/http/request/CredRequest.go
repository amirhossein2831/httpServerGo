package request

type CredRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
