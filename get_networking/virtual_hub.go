package get_networking

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

func GetVirtualHubs(hubId string) (HubConnections, error) {
	var hubConnections HubConnections

	return hubConnections, nil
}
