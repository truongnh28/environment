package dto

type LoginRequest struct {
	Username string `json:"user_name,omitempty"`
	Password string `json:"pass_word,omitempty"`
}

type LoginResponseData struct {
	Account     *Account `json:"account,omitempty"`
	AccessToken string   `json:"access_token,omitempty"`
}

type Account struct {
	Username string `json:"username,omitempty"`
	Code     string `json:"code,omitempty"`
	Role     string `json:"role,omitempty"`
	Status   string `json:"status,omitempty"`
}
