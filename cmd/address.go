/*
Copyright © 2022 John J. Hession

*/
package cmd

import (
	"fmt"
	net "github.com/johhess40/net_spaces/get_networking"
	"log"

	"github.com/spf13/cobra"
)

type DisplayData struct {
	DisplayAll bool
}

var (
	Switch net.SwitchData
)

// displayCmd represents the display command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "display returns data about the spoke you will be deploying",
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
		fmt.Printf("%s", entry)
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)
	addressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addressCmd.Flags().StringVarP(&Switch.Size, "size", "s", "grande", "What size spoke shall we deploy?")
	err := addressCmd.MarkFlagRequired("size")
	if err != nil {
		return
	}
	addressCmd.Flags().StringVarP(&Switch.Space, "space", "z", "", "What is the overall space for the region?")
	err = addressCmd.MarkFlagRequired("space")
	if err != nil {
		return
	}

	addressCmd.Flags().StringVarP(&Switch.Region, "region", "r", "", "Where's your spoke at?")
	err = addressCmd.MarkFlagRequired("region")
	if err != nil {
		return
	}

	addressCmd.Flags().StringVarP(&Switch.Cidr, "cidr", "c", "", "Whats your spokes cidr?")
	err = addressCmd.MarkFlagRequired("region")
	if err != nil {
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
