package cmd

import (
	"os"

	"github.com/jodisfields/snek/cmd/net"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snek",
	Short: "A collection of useful tools",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPaettes() {
	rootCmd.AddCommand(net.NetCmd)
}
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPaettes()
}
