package net

import (
	"github.com/spf13/cobra"
)

// NetCmd represents the net command palette.
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net is a subcommand of snek",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
