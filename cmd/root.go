package cmd

import (
	"github.com/palantir/pkg/cobracli"
	"github.com/spf13/cobra"
)

var Version = "none"

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Random CLI!",
}

func Execute() int {
	return cobracli.ExecuteWithDefaultParams(rootCmd, cobracli.VersionFlagParam(Version))
}
