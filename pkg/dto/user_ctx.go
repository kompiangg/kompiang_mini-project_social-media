package dto

const (
	AccountCTXKey = "ACCOUNT-CTX-KEY"
)

type UserContext struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}
