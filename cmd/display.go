/*
Copyright © 2022 John J. Hession

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type DisplayData struct {
	DisplayAll bool
}

var (
	DataDisplay DisplayData
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "display returns data about the spoke you will be deploying",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating software defined network data!!!!!")
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)
	displayCmd.Flags().BoolVarP(&DataDisplay.DisplayAll, "display-all", "a", true, "Should we return all data by default? Defaults to true...")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
