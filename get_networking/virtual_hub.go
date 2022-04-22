package get_networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Connect struct {
	HubId   string
	HubType string
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

type HubConnections []struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	Properties struct {
		ProvisioningState    string `json:"provisioningState"`
		RemoteVirtualNetwork struct {
			ID string `json:"id"`
		} `json:"remoteVirtualNetwork"`
		EnableInternetSecurity bool `json:"enableInternetSecurity"`
		RoutingConfiguration   struct {
			AssociatedRouteTable struct {
				ID string `json:"id"`
			} `json:"associatedRouteTable"`
			PropagatedRouteTables struct {
				Labels []string `json:"labels"`
				Ids    []struct {
					ID string `json:"id"`
				} `json:"ids"`
			} `json:"propagatedRouteTables"`
			VnetRoutes struct {
				StaticRoutes []struct {
					Name             string   `json:"name"`
					AddressPrefixes  []string `json:"addressPrefixes"`
					NextHopIPAddress string   `json:"nextHopIpAddress"`
				} `json:"staticRoutes"`
				BgpConnections []struct {
					ID string `json:"id"`
				} `json:"bgpConnections"`
			} `json:"vnetRoutes"`
		} `json:"routingConfiguration"`
	} `json:"properties"`
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

func GetVirtualHubConnections(hubId string, t TokenBuilder) (HubConnections, error) {
	var hubConnections HubConnections
	u := fmt.Sprintf("GET https://management.azure.com%s/hubVirtualNetworkConnections?api-version=2021-05-01", hubId)

	token := t.BearerToken.AccessToken

	bearer := "Bearer " + token

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hubConnections, err
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

func GetVirtualNetworkPeerings(hubId string, t TokenBuilder) (VnetConnections, error) {
	var vnetConnections VnetConnections
	u := fmt.Sprintf("GET https://management.azure.com%s/virtualNetworkPeerings?api-version=2021-05-01", hubId)

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

	for _, v := range connections {
		network.VnetIds = append(network.VnetIds, v.Properties.RemoteVirtualNetwork.ID)
	}

	for _, x := range network.VnetIds {
		u := fmt.Sprintf("GET https://management.azure.com%s?api-version=2021-05-01", x)

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

func MakeConnectionSwitches(hubId string, c Connect, t TokenBuilder) {
	switch c.HubType {
	case "vhub":
		connections, err := ParseHubConnections()
		if err != nil {
			return
		}

	case "vnet":
		connections, err := ParseVnetConnections()
		if err != nil {
			return
		}
	}
}
