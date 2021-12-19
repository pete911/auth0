package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"gopkg.in/auth0.v5/management"
	"os"
)

var (
	listLogCmd = &cobra.Command{
		Use:     "log",
		Short:   "log",
		Aliases: []string{"logs"},
		Run:     listLogRun,
	}
)

func listLogRun(_ *cobra.Command, _ []string) {
	m := NewManagement()
	list, err := m.Log.List()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printLogList(list)
}

func printLogList(logs []*management.Log) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Date \t Client ID \t IP \t Description \t Details")
	for _, l := range logs {
		fmt.Fprintf(w, "%s \t %s \t %s \t %s \t %s\n",
			TimeValue(l.Date), StringValue(l.ClientID), StringValue(l.IP), StringValue(l.Description), MapValue(l.Details))
	}
	w.Flush()
}
