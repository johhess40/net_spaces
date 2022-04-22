/*
Copyright © 2022 John J. Hession
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Connect struct {
	HubId string
}

var (
	Connection Connect
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Tests data about an Azure virtual hub connection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect called")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.PersistentFlags().StringVarP(&Connection.HubId, "hub-id", "h", "", "the virtual hub id to connect the vnet to")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
