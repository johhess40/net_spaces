package get_networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
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

type AzureNetData []struct {
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
}

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
		if s.Size == "venti" {
			return n[8:508], nil
		} else if s.Size == "grande" {
			return n[8:2039], nil
		}
		return n, nil

	case "eastus2":
		n, err = MakeTheSwitch(s)
		if err != nil {
			return n, fmt.Errorf(err.Error())
		}

		if s.Size == "venti" {
			return n[512:], nil
		} else if s.Size == "grande" {
			return n[2040:], nil
		}

		return n, nil
	}
	return n, nil
}

func BuildCompositeNetworkData(t TokenBuilder, s Subscription) (AzureNetData, error) {
	var net AzureNetData
	for _, x := range s.Value {
		networks, err := GetNetworks(x.SubscriptionId, t)
		if err != nil {
			return net, fmt.Errorf(err.Error())
		}
		net = append(net, networks.Value...)
	}
	// fmt.Println(net)
	return net, nil
}

// func remove(s []string, index int) ([]string, error) {
//     if index >= len(s) {
//         return nil, errors.New("Out of Range Error")
//     }
//     return append(s[:index], s[index+1:]...), nil
// }

func EvaluateAvailableNetworks(data SwitchData, a AzureNetData) ([]string, error) {
	var used []string
	ret, err := ReturnNetworks(data)
	if err != nil {
		return used, err
	}

	// Create list of used address spaces.
	for _, b := range a {
		for _, z := range b.Properties.AddressSpace.AddressPrefixes {
			used = append(used, z)
		}
	}

	// Compare list of used address spaces to all calculated networks and remove used.
	for i := 0; i < len(ret); i++ {
		for _, u := range used {
			if ret[i] == u {
				ret = append(ret[:i], ret[i+1:]...)
				i--
				break
			}
		}
	}

	// Uncomment for troubleshooting.
	// fmt.Println(ret)
	// fmt.Println(used)
	// fmt.Println(ret)

	return ret, nil
}

func RunAll(data SwitchData, t TokenBuilder, n Subscription) (string, error) {
	var grail []string
	x, err := BuildCompositeNetworkData(t, n)
	if err != nil {
		return "", err
	}
	eval, err := EvaluateAvailableNetworks(data, x)
	if err != nil {
		return "", err
	}
	for _, m := range eval {
		grail = append(grail, m)
	}

	randomIndexString := strconv.Itoa(rand.Intn(time.Now().Nanosecond()))
	var num int
	nextRandom := strings.Split(randomIndexString, "")
	for _, v := range nextRandom {
		num, err = strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		if num < len(grail) {
			return grail[len(grail)-1], nil
		}
	}
	return grail[len(grail)-1], nil
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
