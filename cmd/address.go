/*
Copyright © 2022 John J. Hession
*/
package cmd

import (
	"fmt"
	net "github.com/johhess40/net_spaces/get_networking"
	"log"
	"strings"

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
	Short: "address returns data about the spoke you will be deploying",
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
		fmt.Printf("%s", strings.TrimSpace(entry))
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)
	addressCmd.PersistentFlags().StringVarP(&Switch.Size, "size", "s", "grande", "What size spoke shall we deploy?")
	err := addressCmd.MarkPersistentFlagRequired("size")
	if err != nil {
		log.Fatal(err)
	}
	addressCmd.PersistentFlags().StringVarP(&Switch.Space, "space", "z", "", "What is the overall space for the region?")
	err = addressCmd.MarkPersistentFlagRequired("space")
	if err != nil {
		log.Fatal(err)
	}

	addressCmd.PersistentFlags().StringVarP(&Switch.Region, "region", "r", "", "Where's your spoke at?")
	err = addressCmd.MarkPersistentFlagRequired("region")
	if err != nil {
		log.Fatal(err)
	}

	addressCmd.PersistentFlags().StringVarP(&Switch.Cidr, "cidr", "c", "", "Whats your spokes cidr?")
	err = addressCmd.MarkPersistentFlagRequired("cidr")
	if err != nil {
		log.Fatal(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
