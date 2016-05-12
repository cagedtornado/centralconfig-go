package configmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	apiUrl := client.ServiceUrl + "/config/get"

	//	Check to see if we have an application
	if client.Application == "" {
		err := fmt.Errorf("Please specify an 'Application' to get configuration for")
		return retval, err
	}

	//	Get the machine name if it hasn't been set

	//	Create our request
	request := ConfigItem{
		Application: client.Application,
		Name:        name}

	//	Serialize our request to JSON:
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return retval, err
	}

	// Convert bytes to a reader.
	requestJSON := strings.NewReader(string(requestBytes))

	//	Post the JSON to the api url
	res, err := http.Post(apiUrl, "application/json", requestJSON)
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
	if err := json.Unmarshal(body, &retval); err != nil {
		return retval, err
	}

	return retval, nil
}
