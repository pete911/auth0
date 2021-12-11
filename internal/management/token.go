package management

import "fmt"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AudienceUrl  string `json:"audience"`
}

func NewTokenRequest(domain, clientId, clientSecret string) TokenRequest {
	audienceUrl := fmt.Sprintf("https://%s/api/%s/", domain, apiVersion)
	return TokenRequest{
		GrantType:    "client_credentials",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		AudienceUrl:  audienceUrl,
	}
}
