package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ProfileConfig Config

	profileFlag string

	RootCmd = &cobra.Command{
		Use:              "auth0",
		Short:            "auth0 management client",
		PersistentPreRun: loadConfigRun,
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&profileFlag, "profile", "p", "default", "auth0 profile")
	RootCmd.AddCommand(listCmd)
}

func loadConfigRun(_ *cobra.Command, _ []string) {
	config, err := LoadConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	profileConfig, ok := config[profileFlag]
	if !ok {
		fmt.Printf("profile %s does not exist in config file\n", profileFlag)
		os.Exit(1)
	}
	ProfileConfig = profileConfig
}
