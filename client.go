package configmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//	The centralconfig client
type ConfigClient struct {
	ServiceUrl  string //	The service url
	Application string //	The application to fetch config information for
	Machine     string //	The (optional) machine to fetch config information for.
}

//	Get a specific config item
func (client ConfigClient) Get(name string) (ConfigResponse, error) {
	retval := ConfigResponse{}

	//	First, see if we have a service url set:
	if client.ServiceUrl == "" {
		err := fmt.Errorf("Please specify a 'ServiceUrl' for the centralconfig service")
		return retval, err
	}

	//	Check to see if we have an application
	if client.Application == "" {
		err := fmt.Errorf("Please specify an 'Application' to get configuration for")
		return retval, err
	}

	//	If we have a serviceUrl set, do a get
	res, err := http.Get(client.ServiceUrl)
	defer res.Body.Close()
	if err != nil {
		return retval, err
	}

	//	Read the body of the response if we have one:
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return retval, err
	}

	//	Unmarshall from JSON into our struct:
	cres := &ConfigResponse{}
	if err := json.Unmarshal(body, &cres); err != nil {
		return retval, err
	}

	return retval, nil
}
