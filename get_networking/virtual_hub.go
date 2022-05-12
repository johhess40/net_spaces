package get_networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Confirmed struct {
	Region             string
	ClientId           string
	TenantId           string
	AddressSpace       string
	RemoteConnectionId string
}

type Connect struct {
	HubId   string
	HubType string
}

func (c Connect) Address(sw SwitchData, t TokenBuilder) (string, error) {
	entry, errEntry := Entry(sw, t)
	if errEntry != nil {
		return "", errEntry
	}
	return fmt.Sprintf("%s", strings.TrimSpace(entry)), nil
}

func (c Connect) Gen(sw SwitchData, t TokenBuilder) (string, error) {
	entry, errEntry := Entry(sw, t)
	if errEntry != nil {
		return "", errEntry
	}
	return fmt.Sprintf("%s", strings.TrimSpace(entry)), nil
}

func (c Connect) Generate(j string, con Connect) (string, error) {
	t, err := ExecToken()
	if err != nil {
		return "", err
	}

	switches, err := MakeConnectionSwitches(j, con, t)
	if err != nil {
		return "", err
	}

	connectable, err := ReturnConnectable(switches)
	if err != nil {
		return "", err
	}
	return connectable, nil
}

func (c Connect) CheckLength() error {
	if len(c.HubId) == 0 || len(c.HubType) == 0 {
		return fmt.Errorf("all flags must have a value")
	} else {
		return nil
	}
}

type NetData struct {
	VnetIds      []string
	VnetSpaces   []string
	NetworkStuff struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		Type       string `json:"type"`
		Location   string `json:"location"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
			AddressSpace      struct {
				AddressPrefixes []string `json:"addressPrefixes"`
			} `json:"addressSpace"`
			Subnets []struct {
				Name       string `json:"name"`
				ID         string `json:"id"`
				Properties struct {
					ProvisioningState string `json:"provisioningState"`
					AddressPrefix     string `json:"addressPrefix"`
					IPConfigurations  []struct {
						ID string `json:"id"`
					} `json:"ipConfigurations"`
				} `json:"properties"`
			} `json:"subnets"`
			VirtualNetworkPeerings []interface{} `json:"virtualNetworkPeerings"`
		} `json:"properties"`
	}
}

type HubConnections struct {
	Value []struct {
		Name       string `json:"name"`
		Id         string `json:"id"`
		Etag       string `json:"etag"`
		Type       string `json:"type"`
		Properties struct {
			ProvisioningState    string `json:"provisioningState"`
			ResourceGuid         string `json:"resourceGuid"`
			RoutingConfiguration struct {
				AssociatedRouteTable struct {
					Id string `json:"id"`
				} `json:"associatedRouteTable"`
				PropagatedRouteTables struct {
					Labels []string `json:"labels"`
					Ids    []struct {
						Id string `json:"id"`
					} `json:"ids"`
				} `json:"propagatedRouteTables"`
				VnetRoutes struct {
					StaticRoutes []interface{} `json:"staticRoutes"`
				} `json:"vnetRoutes"`
			} `json:"routingConfiguration"`
			RemoteVirtualNetwork struct {
				Id string `json:"id"`
			} `json:"remoteVirtualNetwork"`
			AllowHubToRemoteVnetTransit         bool   `json:"allowHubToRemoteVnetTransit"`
			AllowRemoteVnetToUseHubVnetGateways bool   `json:"allowRemoteVnetToUseHubVnetGateways"`
			EnableInternetSecurity              bool   `json:"enableInternetSecurity"`
			ConnectivityStatus                  string `json:"connectivityStatus"`
		} `json:"properties"`
	} `json:"value"`
}

type VnetConnections struct {
	Value []struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Properties struct {
			AllowVirtualNetworkAccess bool `json:"allowVirtualNetworkAccess"`
			AllowForwardedTraffic     bool `json:"allowForwardedTraffic"`
			AllowGatewayTransit       bool `json:"allowGatewayTransit"`
			UseRemoteGateways         bool `json:"useRemoteGateways"`
			RemoteVirtualNetwork      struct {
				Id string `json:"id"`
			} `json:"remoteVirtualNetwork"`
			RemoteAddressSpace struct {
				AddressPrefixes []string `json:"addressPrefixes"`
			} `json:"remoteAddressSpace"`
			RemoteVirtualNetworkAddressSpace struct {
				AddressPrefixes []string `json:"addressPrefixes"`
			} `json:"remoteVirtualNetworkAddressSpace"`
			RemoteBgpCommunities struct {
				VirtualNetworkCommunity string `json:"virtualNetworkCommunity"`
				RegionalCommunity       string `json:"regionalCommunity"`
			} `json:"remoteBgpCommunities"`
			PeeringState      string `json:"peeringState"`
			PeeringSyncLevel  string `json:"peeringSyncLevel"`
			ProvisioningState string `json:"provisioningState"`
		} `json:"properties"`
	} `json:"value"`
}

type VirtualHub struct {
	Id         string `json:"id"`
	Etag       string `json:"etag"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		VirtualWan        struct {
			Id string `json:"id"`
		} `json:"virtualWan"`
		VirtualHubRouteTableV2S    []interface{} `json:"virtualHubRouteTableV2s"`
		AddressPrefix              string        `json:"addressPrefix"`
		Sku                        string        `json:"sku"`
		RoutingState               string        `json:"routingState"`
		VirtualRouterAsn           int           `json:"virtualRouterAsn"`
		VirtualRouterIps           []string      `json:"virtualRouterIps"`
		AllowBranchToBranchTraffic bool          `json:"allowBranchToBranchTraffic"`
		PreferredRoutingGateway    string        `json:"preferredRoutingGateway"`
	} `json:"properties"`
}

func GetVirtualHubConnections(hubId string, t TokenBuilder) (HubConnections, error) {
	var hubConnections HubConnections
	u := strings.TrimSpace(fmt.Sprintf("https://management.azure.com%s/hubVirtualNetworkConnections?api-version=2021-05-01", hubId))

	token := t.BearerToken.AccessToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hubConnections, fmt.Errorf("%s for http.NewRequest: %v", err, err)
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return hubConnections, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return hubConnections, err
	}

	err = json.Unmarshal(response, &hubConnections)
	if err != nil {
		return hubConnections, err
	}

	return hubConnections, nil
}

type Vnet struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Type       string `json:"type"`
	Location   string `json:"location"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		AddressSpace      struct {
			AddressPrefixes []string `json:"addressPrefixes"`
		} `json:"addressSpace"`
		Subnets []struct {
			Name       string `json:"name"`
			Id         string `json:"id"`
			Properties struct {
				ProvisioningState string `json:"provisioningState"`
				AddressPrefix     string `json:"addressPrefix"`
				IpConfigurations  []struct {
					Id string `json:"id"`
				} `json:"ipConfigurations"`
			} `json:"properties"`
		} `json:"subnets"`
		VirtualNetworkPeerings []interface{} `json:"virtualNetworkPeerings"`
	} `json:"properties"`
}

func GetVirtualNetworkPeerings(hubId string, t TokenBuilder) (VnetConnections, error) {
	var vnetConnections VnetConnections
	u := fmt.Sprintf("https://management.azure.com%s/virtualNetworkPeerings?api-version=2021-05-01", hubId)

	token := t.BearerToken.AccessToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return vnetConnections, err
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return vnetConnections, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return vnetConnections, err
	}

	err = json.Unmarshal(response, &vnetConnections)
	if err != nil {
		return vnetConnections, err
	}

	return vnetConnections, nil
}

func ParseHubConnections(hubId string, t TokenBuilder) ([]string, error) {
	var network NetData
	var netspaces []string

	connections, err := GetVirtualHubConnections(hubId, t)
	if err != nil {
		return netspaces, err
	}

	for _, v := range connections.Value {
		network.VnetIds = append(network.VnetIds, v.Properties.RemoteVirtualNetwork.Id)
	}

	for _, x := range network.VnetIds {
		u := fmt.Sprintf("https://management.azure.com%s?api-version=2021-05-01", x)

		token := t.BearerToken.AccessToken

		bearer := "Bearer " + token

		request, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return netspaces, err
		}

		request.Header.Add("Authorization", bearer)

		client := &http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			return netspaces, err
		}

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return netspaces, err
		}

		err = json.Unmarshal(response, &network.NetworkStuff)
		if err != nil {
			return netspaces, err
		}
		netspaces = append(netspaces, network.NetworkStuff.Properties.AddressSpace.AddressPrefixes...)
	}
	return netspaces, nil
}

func ParseVnetConnections(hubId string, t TokenBuilder) ([]string, error) {
	var network NetData

	connections, err := GetVirtualNetworkPeerings(hubId, t)
	if err != nil {
		return network.VnetSpaces, err
	}

	for _, v := range connections.Value {
		network.VnetSpaces = append(network.VnetSpaces, v.Properties.RemoteVirtualNetworkAddressSpace.AddressPrefixes...)
	}

	return network.VnetSpaces, nil
}

func MakeConnectionSwitches(j string, c Connect, t TokenBuilder) (Confirmed, error) {
	var conf Confirmed
	switch c.HubType {
	case "vhub":
		connections, err := ParseHubConnections(c.HubId, t)
		if err != nil {
			return conf, err
		}

		jsonData := struct {
			Region       string
			ClientId     string
			TenantId     string
			AddressSpace string
		}{}

		err = json.Unmarshal([]byte(j), &jsonData)
		if err != nil {
			return conf, err
		}

		for _, v := range connections {
			if v == jsonData.AddressSpace {
				return conf, err
			}
		}
		conf.RemoteConnectionId = c.HubId
		conf.Region = jsonData.Region
		conf.ClientId = jsonData.ClientId
		conf.TenantId = jsonData.TenantId
		conf.AddressSpace = jsonData.AddressSpace

		return conf, nil
	case "vnet":
		connections, err := ParseVnetConnections(c.HubId, t)
		if err != nil {
			return conf, err
		}

		jsonData := struct {
			Region       string
			ClientId     string
			TenantId     string
			AddressSpace string
		}{}

		err = json.Unmarshal([]byte(j), &jsonData)
		if err != nil {
			return conf, err
		}

		for _, v := range connections {
			if v == jsonData.AddressSpace {
				return conf, err
			}
		}
		conf.RemoteConnectionId = c.HubId
		conf.Region = jsonData.Region
		conf.ClientId = jsonData.ClientId
		conf.TenantId = jsonData.TenantId
		conf.AddressSpace = jsonData.AddressSpace

		return conf, nil
	default:
		return conf, nil
	}
}

func ReturnConnectable(c Confirmed) (string, error) {
	s, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return string(s), nil
}

func GetVirtualHubData(hubId string, t TokenBuilder) (VirtualHub, error) {
	var hub VirtualHub
	u := strings.TrimSpace(fmt.Sprintf("https://management.azure.com%s/?api-version=2021-05-01", hubId))

	token := t.BearerToken.AccessToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hub, fmt.Errorf("%s for http.NewRequest: %v", err, err)
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return hub, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return hub, err
	}

	err = json.Unmarshal(response, &hub)
	if err != nil {
		return hub, err
	}

	return hub, nil
}

func GetVirtualNetworkHubData(hubId string, t TokenBuilder) (Vnet, error) {
	var hub Vnet
	u := strings.TrimSpace(fmt.Sprintf("https://management.azure.com%s/?api-version=2021-05-01", hubId))

	token := t.BearerToken.AccessToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hub, fmt.Errorf("%s for http.NewRequest: %v", err, err)
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return hub, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return hub, err
	}

	err = json.Unmarshal(response, &hub)
	if err != nil {
		return hub, err
	}

	return hub, nil
}
