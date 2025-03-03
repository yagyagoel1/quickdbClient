package cmd

import (
	"os"

	"github.com/spf13/cobra"
	hashmapcmd "github.com/yagyagoel1/quickdbClient/internal/HashmapCmd"
	basecmd "github.com/yagyagoel1/quickdbClient/internal/baseCmd"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quickdb",
	Short: "It is the client side for the quickdb server",
	Long: `quickdb is a CLI library helps you to interact with the quickdb server by using 
	resp protocol.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(basecmd.SetCmd)
	rootCmd.AddCommand(basecmd.GetCmd)
	rootCmd.AddCommand(basecmd.PingCmd)
	rootCmd.AddCommand(hashmapcmd.HsetCmd)
	rootCmd.AddCommand(hashmapcmd.HgetCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
