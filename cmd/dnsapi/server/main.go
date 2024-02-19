package server

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	projectId string
)

func init() {

	globalFlags := rootCmd.PersistentFlags()
	globalFlags.StringVar(&projectId, "project-id", os.Getenv("PROJECT_ID"), "Project ID of firestore project")

	rootCmd.AddCommand()
	rootCmd.AddCommand(run)
}

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Root command for the DNS API server",
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

var run = &cobra.Command{
	Use:   "run",
	Short: "command to start the DNS API server",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
