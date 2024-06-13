package cmd

import (
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge the site's page cache",
	Long:  "Purge the site's page cache",
	Run: func(cmd *cobra.Command, args []string) {
		domain := CurrentDomain()

		PurgeCache(domain)
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
