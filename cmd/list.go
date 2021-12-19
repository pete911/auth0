package cmd

import (
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list auth0 resources",
	}
)

func init() {
	listCmd.AddCommand(listUserCmd)
	listCmd.AddCommand(listLogCmd)
}
