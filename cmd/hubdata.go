/*
Copyright © 2022 John

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// hubdataCmd represents the hubdata command
var hubdataCmd = &cobra.Command{
	Use:   "hub-data",
	Short: "hub-data returns data bout the hub that we want to connect to",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hubdata called")
	},
}

func init() {
	rootCmd.AddCommand(hubdataCmd)
	hubdataCmd.Flags().StringVarP(&Connection.HubId, "hub-id", "h", "", "the virtual hub id to connect the vnet to")
	err := connectCmd.MarkFlagRequired("hub-id")
	if err != nil {
		log.Fatal(err)
	}
	hubdataCmd.Flags().StringVarP(&Connection.HubType, "out-type", "o", "", "output type for hub data being returned(hcl,json,yml)")
	err = connectCmd.MarkFlagRequired("out-type")
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
