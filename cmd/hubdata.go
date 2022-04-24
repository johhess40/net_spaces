/*
Copyright © 2022 John
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type HubBuilder struct {
	Id      string
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
		fmt.Println("hub-data called")
	},
}

func init() {
	rootCmd.AddCommand(hubdataCmd)
	hubdataCmd.Flags().StringVarP(&Hub.Id, "hub-id", "h", "", "the virtual hub id to connect the vnet to")
	err := hubdataCmd.MarkFlagRequired("hub-id")
	if err != nil {
		log.Fatal(err)
	}
	hubdataCmd.Flags().StringVarP(&Hub.OutType, "out-type", "o", "", "output type for hub data being returned(hcl,json,yml)")
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
