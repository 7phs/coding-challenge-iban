package cmd

import (
	"fmt"
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/spf13/cobra"
)

var (
	GitHash   string // should be uninitialized
	BuildTime string // should be uninitialized
)

var RootCmd = &cobra.Command{
	Use:   config.ApplicationName,
	Short: config.ApplicationDescription,
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ApplicationInfo())
	},
}

func ApplicationInfo() string {
	return config.ApplicationTitle + " " + config.Version + " [" + GitHash + "] " + BuildTime
}
