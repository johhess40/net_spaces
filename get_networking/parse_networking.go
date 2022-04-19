package get_networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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

func ReturnNetworks(s SwitchData) ([]string, error) {
	var n []string
	var err error

	switch s.Region {
	case "westus2":
		n, err = MakeTheSwitch(s)
		if err != nil {
			return n, fmt.Errorf(err.Error())
		}

		return n, nil
	case "eastus2":
		n, err = MakeTheSwitch(s)
		if err != nil {
			return n, fmt.Errorf(err.Error())
		}

		return n, nil
	}
	return n, nil
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

func EvaluateAvailableNetworks(data SwitchData, a AzureNetworks) ([]string, error) {
	var available []string
	ret, err := ReturnNetworks(data)
	if err != nil {
		return available, err
	}

	for _, v := range ret {
		for _, b := range a.Value {
			for _, z := range b.Properties.AddressSpace.AddressPrefixes {
				if z != v {
					available = append(available, v)
				}
			}
		}
	}
	return available, nil
}

func RunAll(data SwitchData, t TokenBuilder, n Subscriptions) (string, error) {
	var grail []string
	x, err := BuildCompositeNetworkData(t, n)
	if err != nil {
		return "", err
	}
	for _, a := range x {
		eval, err := EvaluateAvailableNetworks(data, a)
		if err != nil {
			return "", err
		}
		for _, m := range eval {
			grail = append(grail, m)
		}
	}
	randomIndex := rand.Intn(len(grail))

	return grail[randomIndex], nil
}

func JsonReturn(data SwitchData, t TokenBuilder) (struct {
	Region       string
	ClientId     string
	TenantId     string
	AddressSpace string
}, error) {
	var jsonData struct {
		Region       string
		ClientId     string
		TenantId     string
		AddressSpace string
	}

	n, err := GetSubscriptions(t)
	if err != nil {
		return jsonData, err
	}

	all, err := RunAll(data, t, n)
	if err != nil {
		return jsonData, err
	}
	jsonData = struct {
		Region       string
		ClientId     string
		TenantId     string
		AddressSpace string
	}{
		Region:       data.Region,
		ClientId:     t.ClientId,
		TenantId:     t.TenantId,
		AddressSpace: all,
	}

	return jsonData, nil
}

func CreateJson(data SwitchData, t TokenBuilder) (string, error) {
	createJson, err := JsonReturn(data, t)
	if err != nil {
		return "", err
	}
	s, err := json.Marshal(createJson)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return string(s), nil
}

func Entry(data SwitchData, t TokenBuilder) (string, error) {
	d, err := CreateJson(data, t)
	if err != nil {
		return d, err
	}
	return fmt.Sprintf("%s", d), nil
}
