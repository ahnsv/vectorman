package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// write version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of vectorman",
	Long:  `All software has versions. This is vectorman's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("vectorman %s\n", Version)
		fmt.Printf("Build Time: %s\n", BuildTime)
		fmt.Printf("Git Commit: %s\n", GitCommit)
		fmt.Printf("Git Branch: %s\n", GitBranch)
		fmt.Printf("Go Version: %s\n", GoVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
