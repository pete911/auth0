package cmd

import (
	"fmt"
	"text/tabwriter"

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
	printUserList(list.Users)
}

func printUserList(users []*management.User) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Email \t Logins Count \t Created At \t Last Login \t Connections")
	for _, u := range users {
		fmt.Fprintf(w, "%s \t %s \t %s \t %s \t %v\n",
			StringValue(u.Email), Int64Value(u.LoginsCount), TimeValue(u.CreatedAt), TimeValue(u.LastLogin), getIdentityConnections(u))
	}
	w.Flush()
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
