package parsing

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2/hclwrite"
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

func WriteHubConnection() {
	writer := hclwrite.NewFile()

	hclFile, err := os.Create("connect.tf")
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
	build.SetAttributeValue("virtual_hub_id", cty.StringVal(fmt.Sprintf("%s", u)))
	build.SetAttributeValue("remote_virtual_network_id", cty.StringVal(fmt.Sprintf("%s", u)))
	_, err = hclFile.WriteAt(writer.Bytes(), seek)
	if err != nil {
		log.Fatal(err)
	}
}

func Request() {}
