package parsing

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	net "github.com/johhess40/net_spaces/get_networking"
	"github.com/zclconf/go-cty/cty"
	"log"
	"os"
)

type Net interface {
	Vnet()
	Vhub()
	Peer()
}

func (n *Vnet) Vnet() {

}

func (n *Vhub) Vhub() {

}

func (n *Vhub) Peer() {

}
func (n *Vnet) Peer() {

}

type Vnet struct{}

type VnetPeering struct{}

type VhubConnection struct{}

type Vhub struct{}

type Connection struct{}

func WriteHubConnection(s string, hub net.VirtualHub) {
	writer := hclwrite.NewFile()

	hclFile, err := os.OpenFile("connect.tf", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	seek, errseek := hclFile.Seek(0, 2)
	if errseek != nil {
		log.Fatal(errseek)
	}

	rootBod := writer.Body()

	b := rootBod.AppendNewBlock("resource", []string{"azurerm_virtual_hub_connection", "connect"})

	build := b.Body()

	u := uuid.New()

	build.SetAttributeValue("name", cty.StringVal(fmt.Sprintf("%s", u)))
	build.SetAttributeValue("virtual_hub_id", cty.StringVal(fmt.Sprintf("%s", hub.Id)))
	build.SetAttributeValue("remote_virtual_network_id", cty.StringVal(fmt.Sprintf("%s", s)))

	r := rootBod.AppendNewBlock("output", []string{"connection"})
	value := r.Body()

	value.SetAttributeRaw("value", hclwrite.Tokens{
		{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(`azurerm_virtual_hub_connection.connect.id`),
		},
	})

	_, err = hclFile.WriteAt(writer.Bytes(), seek)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteVnetConnection(s string, hub net.VirtualHub) {
	writer := hclwrite.NewFile()

	hclFile, err := os.OpenFile("connect.tf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	seek, errseek := hclFile.Seek(0, 2)
	if errseek != nil {
		log.Fatal(errseek)
	}

	rootBod := writer.Body()

	b := rootBod.AppendNewBlock("resource", []string{"azurerm_virtual_hub_connection", "connect"})

	build := b.Body()

	u := uuid.New()

	build.SetAttributeValue("name", cty.StringVal(fmt.Sprintf("%s", u)))
	build.SetAttributeValue("virtual_hub_id", cty.StringVal(fmt.Sprintf("%s", hub.Id)))
	build.SetAttributeValue("remote_virtual_network_id", cty.StringVal(fmt.Sprintf("%s", s)))

	//route := build.AppendNewBlock("routing",[]string{})
	//
	//r := route.Body()

	_, err = hclFile.WriteAt(writer.Bytes(), seek)
	if err != nil {
		log.Fatal(err)
	}
}

func Request(remoteId string, hub net.VirtualHub) {
	WriteHubConnection(remoteId, hub)
}
