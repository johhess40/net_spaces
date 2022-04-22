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

type Connect net.Connect

var (
	Connection Connect
)

func (s Connect) Address(sw net.SwitchData) (string, error) {
	token, err := net.ExecToken()
	if err != nil {
		return "", err
	}
	entry, errEntry := net.Entry(Switch, token)
	if errEntry != nil {
		return "", err
	}
	return fmt.Sprintf("%s", strings.TrimSpace(entry)), nil
}

func (s Connect) CheckLength() error {
	if len(s.HubId) == 0 || len(s.HubType) == 0 {
		return fmt.Errorf("all flags must have a value")
	} else {
		return nil
	}
}

func (s Connect) CheckValues() error {
	if len(strings.Split(s.HubId, "/"))%8 != 0 {
		return fmt.Errorf("hub id must be divisible by 8 your hub id is wrong see video here on how to properly enter resource id's => https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	} else if s.HubType != "vhub" || s.HubType != "vnet" {
		return fmt.Errorf("hub type must be vhub or vnet")
	} else {
		return nil
	}
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Tests data about an Azure virtual hub connection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		con, err := Connection.Address(Switch)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	addressCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&Connection.HubId, "hub-id", "h", "", "the virtual hub id to connect the vnet to")
	err := connectCmd.MarkFlagRequired("hub-id")
	c := Connection.CheckLength()
	if err != nil || c != nil {
		if err == nil {
			err = Connection.CheckLength()
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}
	connectCmd.Flags().StringVarP(&Connection.HubType, "hub-type", "t", "virtual-hub", "the virtual hub id to connect the vnet to")
	err = connectCmd.MarkFlagRequired("hub-type")
	if err != nil || c != nil {
		if err == nil {
			err = Connection.CheckLength()
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}

	ch := Connection.CheckValues()
	if ch != nil {
		log.Fatal(ch)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
