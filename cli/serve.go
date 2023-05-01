package cli

import (
	"github.com/ahnsv/vectorman/router"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the vectorman server",
	Long:  `Start the vectorman server`,
	Run: func(cmd *cobra.Command, args []string) {
		vectormanRouter := router.CreateRouter()
		vectormanRouter.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
