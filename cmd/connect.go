/*
Copyright © 2022 John J. Hession
*/
package cmd

import (
	"fmt"
	net "github.com/johhess40/net_spaces/get_networking"
	"github.com/spf13/cobra"
	"log"
)

var (
	Connection net.Connect
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Tests data about an Azure virtual hub connection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := net.ExecToken()
		if err != nil {
			log.Fatal(err)
		}
		entry, errEntry := net.Entry(Switch, token)
		if errEntry != nil {
			return
		}
		generate, err := Connection.Generate(entry, Connection)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(generate)
	},
}

func init() {
	addressCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&Connection.HubId, "hub-id", "i", "", "the virtual hub id to connect the vnet to")
	err := connectCmd.MarkFlagRequired("hub-id")
	if err != nil {
		log.Fatal(err)
	}
	connectCmd.Flags().StringVarP(&Connection.HubType, "hub-type", "t", "", "the virtual hub type to connect the vnet to")
	err = connectCmd.MarkFlagRequired("hub-type")
	if err != nil {
		log.Fatal(err)
	}

	//ch := Connection.CheckValues()
	//if ch != nil {
	//	log.Fatal(ch)
	//}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
