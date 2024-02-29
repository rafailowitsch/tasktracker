package domain

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
