package get_networking

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//TokenBuilder is used to generate and return token needed to query resources in Azure
type TokenBuilder struct {
	ClientId       string
	ClientSecret   string
	TenantId       string
	SubscriptionId string
	BearerToken    struct {
		TokenType    string `json:"token_type"`
		ExpiresIn    string `json:"expires_in"`
		ExtExpiresIn string `json:"ext_expires_in"`
		ExpiresOn    string `json:"expires_on"`
		NotBefore    string `json:"not_before"`
		Resource     string `json:"resource"`
		AccessToken  string `json:"access_token"`
	}
}

func GenerateToken() (TokenBuilder, error) {
	var t TokenBuilder

	t.ClientId = os.Getenv("ARM_CLIENT_ID")
	if len(t.ClientId) == 0 {
		return t, fmt.Errorf("env var client_id must be set")
	}
	t.ClientSecret = os.Getenv("ARM_CLIENT_SECRET")
	if len(t.ClientId) == 0 {
		return t, fmt.Errorf("env var client_secret must be set")
	}
	t.TenantId = os.Getenv("ARM_TENANT__ID")
	if len(t.ClientId) == 0 {
		return t, fmt.Errorf("env var tenant_id must be set")
	}
	t, err := DoTokenRequest(t)
	if err != nil {
		return t, err
	}

	return t, nil
}

//DoTokenRequest creates a token for use later on with pulling data on networking
func DoTokenRequest(t TokenBuilder) (TokenBuilder, error) {
	endpoint := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", t.TenantId)

	values := url.Values{}
	values.Set("grant_type", "client_credentials")
	values.Set("client_id", t.ClientId)
	values.Set("client_secret", t.ClientSecret)
	values.Set("resource", "https://management.azure.com")

	client := &http.Client{}

	r, err := http.NewRequest("POST", endpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return t, fmt.Errorf("error calling http.NewRequest: %s", err.Error())
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	res, err := client.Do(r)
	if err != nil {
		return t, fmt.Errorf("error calling http.Client.Do: %s", err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return t, fmt.Errorf("error reading response body for token generation: %s", err.Error())
	}

	err = json.Unmarshal(body, &t.BearerToken)
	if err != nil {
		return t, fmt.Errorf("error unmarshalling JSON for token generation: %s", err.Error())

	}
	return t, nil
}

func ExecToken() (TokenBuilder, error) {
	token, err := GenerateToken()
	if err != nil {
		return TokenBuilder{}, err
	}

	request, err := DoTokenRequest(token)
	if err != nil {
		return TokenBuilder{}, err
	}

	return request, nil
}
