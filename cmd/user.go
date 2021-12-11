package cmd

import (
	"fmt"
	"github.com/pete911/auth0/internal/management"
	"github.com/pete911/auth0/internal/user"
	"github.com/spf13/cobra"
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
	manageClient, err := management.NewClient(UserConfig.Domain, UserConfig.ClientId, UserConfig.ClientSecret)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	users, err := user.NewClient(manageClient).ListUsers()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, user := range users {
		fmt.Printf("%s\t%s\t%s\t %v\n", user.Email, user.CreatedAt, user.LastLogin, user.IdentityConnections())
	}
}
