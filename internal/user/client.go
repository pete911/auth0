package user

import (
	"github.com/pete911/auth0/internal/management"
	"net/http"
)

type Client struct {
	manageClient management.Client
}

func NewClient(manageClient management.Client) Client {
	return Client{manageClient: manageClient}
}

func (c Client) ListUsers() ([]User, error) {
	var out []User
	err := c.manageClient.Do(http.MethodGet, "users", nil, nil, &out)
	return out, err
}
