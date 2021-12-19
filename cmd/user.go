package cmd

import (
	"fmt"
	"github.com/pete911/auth0/cmd/flags"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gopkg.in/auth0.v5/management"
	"os"
)

var (
	listUserFlags = flags.ListUserFlags{}

	listUserCmd = &cobra.Command{
		Use:     "user",
		Short:   "user",
		Aliases: []string{"users"},
		Run:     listUserRun,
	}
)

func init() {
	listUserCmd.Flags().StringVar(&listUserFlags.Email, "email", "", "user email")
	listUserCmd.Flags().StringVar(&listUserFlags.Name, "name", "", "user name")
}

func listUserRun(_ *cobra.Command, _ []string) {
	m := NewManagement()
	list, err := m.User.List(listUserFlags.GetRequestOptions()...)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printUserList(list.Users)
}

func printUserList(users []*management.User) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Email \t Name \t Logins Count \t Created At \t Last Login \t Connections")
	for _, u := range users {
		fmt.Fprintf(w, "%s \t %s \t %s \t %s \t %s \t %v\n",
			StringValue(u.Email), StringValue(u.Name), Int64Value(u.LoginsCount), TimeValue(u.CreatedAt), TimeValue(u.LastLogin), getIdentityConnections(u))
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
