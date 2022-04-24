/*
Copyright © 2022 John
*/
package cmd

import (
	"fmt"
	net "github.com/johhess40/net_spaces/get_networking"
	"log"

	"github.com/spf13/cobra"
)

type HubBuilder struct {
	Id      string
	Type    string
	OutType string
}

var (
	Hub HubBuilder
)

// hubdataCmd represents the hubdata command
var hubdataCmd = &cobra.Command{
	Use:   "hub-data",
	Short: "hub-data returns data bout the hub that we want to connect to",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := net.ExecToken()
		if err != nil {
			return
		}
		switch Hub.Type {
		case "vhub":
			data, err := net.GetVirtualHubData(Hub.Id, token)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
		case "vnet":
			data, err := net.GetVirtualNetworkHubData(Hub.Id, token)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
		default:
			data, err := net.GetVirtualHubData(Hub.Id, token)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
		}
	},
}

func init() {
	rootCmd.AddCommand(hubdataCmd)
	hubdataCmd.Flags().StringVarP(&Hub.Id, "hub-id", "i", "null", "the virtual hub id to connect the vnet to")
	err := hubdataCmd.MarkFlagRequired("hub-id")
	if err != nil {
		log.Fatal(err)
	}
	hubdataCmd.Flags().StringVarP(&Hub.Type, "hub-type", "t", "vhub", "hub type for lookup")
	err = hubdataCmd.MarkFlagRequired("hub-type")
	if err != nil {
		log.Fatal(err)
	}
	hubdataCmd.Flags().StringVarP(&Hub.OutType, "out-type", "o", "json", "output type for hub data being returned(hcl,json,yml)")
	err = hubdataCmd.MarkFlagRequired("out-type")
	if err != nil {
		log.Fatal(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hubdataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hubdataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
