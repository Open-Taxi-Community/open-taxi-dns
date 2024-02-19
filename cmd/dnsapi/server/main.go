package server

import "github.com/spf13/cobra"

func init() {

}

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the DNS server",
	Run: func(cmd *cobra.Command, args []string) {

		return nil
	},
}
