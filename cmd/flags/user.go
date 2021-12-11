package flags

import (
	"fmt"
	"gopkg.in/auth0.v5/management"
)

type ListUserFlags struct {
	Email string
	Name  string
}

func (l ListUserFlags) GetRequestOptions() []management.RequestOption {
	var requestOptions []management.RequestOption
	if l.Email != "" {
		requestOptions = append(requestOptions, management.Query(fmt.Sprintf(`email:"%s"`, l.Email)))
	}
	if l.Name != "" {
		requestOptions = append(requestOptions, management.Query(fmt.Sprintf(`name:"%s"`, l.Name)))
	}
	return requestOptions
}
