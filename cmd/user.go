package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/auth0.v5/management"
	"os"
)

var (
	listUserCmd = &cobra.Command{
		Use:     "user",
		Short:   "user",
		Aliases: []string{"users"},
		Run:     listUserRun,
	}
)

func listUserRun(_ *cobra.Command, _ []string) {
	m, err := management.New(UserConfig.Domain, management.WithClientCredentials(UserConfig.ClientId, UserConfig.ClientSecret))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	list, err := m.User.List()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, user := range list.Users {
		fmt.Printf("%s\t%s\t%s\t %v\n", StringValue(user.Email), user.CreatedAt, user.LastLogin, getIdentityConnections(user))
	}
}

func getIdentityConnections(u *management.User) []string {
	if u == nil {
		return nil
	}

	var out []string
	for _, identity := range u.Identities {
		if identity != nil {
			out = append(out, StringValue(identity.Connection))
		}
	}
	return out
}
