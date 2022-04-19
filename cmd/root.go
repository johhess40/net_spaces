/*
Copyright © 2022 John J. Hession

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	NetStuff NetData
)

type NetData struct {
	NetSize   string
	NetName   string
	NetRegion string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "net_spaces",
	Short: "Returns the next available address space for your org!",
	Long: `
	
`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.net_spaces.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&NetStuff.NetSize, "netsize", "s", "smol", "What size spoke shall we deploy?")
	rootCmd.Flags().StringVarP(&NetStuff.NetName, "netname", "n", "", "What's you spokes name?")
	rootCmd.Flags().StringVarP(&NetStuff.NetRegion, "region", "r", "", "Where's your spoke at?")
}
