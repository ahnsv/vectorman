package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// write a cli using cobra to start server
var (
	// Version is the version of the application
	Version string
	// BuildTime is the time the application was built
	BuildTime string
	// GitCommit is the git commit hash of the application
	GitCommit string
	// GitBranch is the git branch of the application
	GitBranch string
	// GoVersion is the version of Go used to build the application
	GoVersion string
)

// root command
var rootCmd = &cobra.Command{
	Use:   "vectorman",
	Short: "vectorman is a service for managing on-call schedules",
	Long:  `vectorman is a service for managing on-call schedules. It is built using the CQRS pattern.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vectorman")
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
}
