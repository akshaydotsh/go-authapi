package models

type HttpResponse struct {
	Message               string      `json:"message"`
	Error                 interface{} `json:"error,omitempty"`
	AccessToken           string      `json:"access_token,omitempty"`
	AccessTokenExpiresAt  int64       `json:"access_token_expires_at,omitempty"`
	RefreshToken          string      `json:"refresh_token,omitempty"`
	RefreshTokenExpiresAt int64       `json:"refresh_token_expires_at,omitempty"`
	Users                 []User      `json:"users,omitempty"`
}
