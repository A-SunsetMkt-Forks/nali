package cmd

import (
	"github.com/zu1k/nali/internal/app"
	"github.com/zu1k/nali/internal/ipdb"

	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Query IP information",
	Long:  `Query IP information.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitIPDB(ipdb.GetIPDBType())
		app.ParseIPs(args)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
