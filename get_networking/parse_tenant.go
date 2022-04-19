package get_networking

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Subscription struct {
	Value []struct {
		Id             string `json:"id"`
		SubscriptionId string `json:"subscriptionId"`
		TenantId       string `json:"tenantId"`
		DisplayName    string `json:"displayName"`
		State          string `json:"state"`
	} `json:"value"`
}

//GetSubscriptions reads all the subscriptions available for the tenant
func GetSubscriptions(t TokenBuilder) (Subscription, error) {
	var s Subscription

	u := "https://management.azure.com/subscriptions?api-version=2020-01-01"

	bearer := "Bearer " + t.BearerToken.AccessToken

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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(response, &s)
	if err != nil {
		log.Fatal(err)
	}

	return s, nil
}
