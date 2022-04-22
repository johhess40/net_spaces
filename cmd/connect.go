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

func (c get_networkingConnect) Generate() (struct {
	Token net.TokenBuilder
	Conn  string
}, error) {
	token, err := net.ExecToken()
	if err != nil {
		return struct{	Token net.TokenBuilder
			Conn  string}{}, err
	}

	con, err := Connection.Address(Switch, token)
	if err != nil {
		return struct{	Token net.TokenBuilder
			Conn  string}{}, err
	}

	return struct{	Token net.TokenBuilder
		Conn  string}{
		Token: token,
		Conn: con,
	}, nil
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Tests data about an Azure virtual hub connection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		gen, error := Connection.

		switches, err := net.MakeConnectionSwitches(con, Connection, token)
		if err != nil {
			return
		}

		connectable, err := net.ReturnConnectable(switches)
		if err != nil {
			return
		}

		fmt.Printf("%s", connectable)

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
