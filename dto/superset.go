package dto

type SupersetAuthRequest struct {
	Password string `json:"password"`
	Provider string `json:"provider"`
	Username string `json:"username"`
	Refresh  bool   `json:"refresh"`
}

type SupersetUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type SupersetResourceRequest struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}
type SupersetRLSRequest struct {
	Clause  string `json:"clause"`
	Dataset int64  `json:"dataset"`
}
type SupersetGuestTokenRequest struct {
	User      SupersetUserRequest       `json:"user"`
	Resources []SupersetResourceRequest `json:"resources"`
	Rls       []SupersetRLSRequest      `json:"rls"`
}

type SupersetTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SupersetCsrfResponse struct {
	Result string `json:"result"`
}

type SupersetGuestTokenResponse struct {
	Token string `json:"token"`
}
