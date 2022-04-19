package get_networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//AzureNetworks contains data about the virtual networks per subscription in Azure
type AzureNetworks struct {
	SubscriptionId string
	Value          []struct {
		Name     string `json:"name"`
		ID       string `json:"id"`
		Etag     string `json:"etag"`
		Type     string `json:"type"`
		Location string `json:"location"`
		Tags     struct {
		} `json:"tags"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
			ResourceGUID      string `json:"resourceGuid"`
			AddressSpace      struct {
				AddressPrefixes []string `json:"addressPrefixes"`
			} `json:"addressSpace"`
			Subnets []struct {
				Name       string `json:"name"`
				ID         string `json:"id"`
				Etag       string `json:"etag"`
				Properties struct {
					ProvisioningState string `json:"provisioningState"`
					AddressPrefix     string `json:"addressPrefix"`
					IPConfigurations  []struct {
						ID string `json:"id"`
					} `json:"ipConfigurations"`
					Delegations                       []interface{} `json:"delegations"`
					PrivateEndpointNetworkPolicies    string        `json:"privateEndpointNetworkPolicies"`
					PrivateLinkServiceNetworkPolicies string        `json:"privateLinkServiceNetworkPolicies"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"subnets"`
			VirtualNetworkPeerings []interface{} `json:"virtualNetworkPeerings"`
			EnableDdosProtection   bool          `json:"enableDdosProtection"`
		} `json:"properties"`
	} `json:"value"`
}

type AzureNetData []AzureNetworks

//GetNetworks reads all the virtual networks in a given subscription for Azure
func GetNetworks(s string, t TokenBuilder) (AzureNetworks, error) {

	u := fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Network/virtualNetworks?api-version=2021-04-01", s)

	token := t.BearerToken.AccessToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var a AzureNetworks
	a.SubscriptionId = s
	err = json.Unmarshal(response, &a)
	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func BuildCompositeNetworkData(t TokenBuilder, s Subscriptions) (AzureNetData, error) {
	var d AzureNetData
	for _, x := range s {
		networks, err := GetNetworks(x.SubscriptionId, t)
		if err != nil {
			return d, fmt.Errorf(err.Error())
		}
		d = append(d, networks)
	}
	return d, nil
}
