package user

import "time"

type User struct {
	UserId        string                 `json:"user_id"`
	Email         string                 `json:"email"`
	EmailVerified bool                   `json:"email_verified"`
	Username      string                 `json:"username"`
	PhoneNumber   string                 `json:"phone_number"`
	PhoneVerified bool                   `json:"phone_verified"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	Identities    []Identity             `json:"identities"`
	AppMetadata   map[string]interface{} `json:"user_metadata"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	Picture       string                 `json:"picture"`
	Name          string                 `json:"name"`
	Nickname      string                 `json:"nickname"`
	Multifactor   []string               `json:"multifactor"`
	LastIP        string                 `json:"last_ip"`
	LastLogin     time.Time              `json:"last_login"`
	LoginsCount   int                    `json:"logins_count"`
	Blocked       bool                   `json:"blocked"`
	GivenName     string                 `json:"given_name"`
	FamilyName    string                 `json:"family_name"`
}

func (u User) IdentityConnections() []string {
	var out []string
	for _, identity := range u.Identities {
		out = append(out, identity.Connection)
	}
	return out
}

type Identity struct {
	Connection string `json:"connection"`
	UserId     string `json:"user_id"`
	Provider   string `json:"provider"`
	IsSocial   bool   `json:"isSocial"`
}
