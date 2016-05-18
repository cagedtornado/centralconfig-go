package configmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	if client.Machine == "" {
		hostname, err := os.Hostname()
		if err == nil {
			client.Machine = hostname
		}
	}

	//	Create our request
	request := ConfigItem{
		Application: client.Application,
		Name:        name}

	//	Serialize our request to JSON:
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(&request)
	if err != nil {
		return retval, err
	}

	// Convert bytes to a reader.
	requestJSON := strings.NewReader(requestBytes.String())

	//	Post the JSON to the api url
	res, err := http.Post(apiUrl, "application/json", requestJSON)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return retval, err
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(&retval)
	if err != nil {
		return retval, err
	}

	return retval, nil
}

//	Get a specific config item
func (client ConfigClient) GetAll() (ConfigResponseMultiple, error) {
	retval := ConfigResponseMultiple{}

	//	First, see if we have a service url set:
	if client.ServiceUrl == "" {
		err := fmt.Errorf("Please specify a 'ServiceUrl' for the centralconfig service")
		return retval, err
	}

	apiUrl := client.ServiceUrl + "/config/getall"

	// Convert bytes to a reader.
	requestJSON := strings.NewReader("")

	//	Post the JSON to the api url
	res, err := http.Post(apiUrl, "application/json", requestJSON)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return retval, err
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(&retval)
	if err != nil {
		return retval, err
	}

	return retval, nil
}

//	Set creates or updates a config item
func (client ConfigClient) Set(request *ConfigItem) (ConfigResponse, error) {
	retval := ConfigResponse{}

	//	First, see if we have a service url set:
	if client.ServiceUrl == "" {
		err := fmt.Errorf("Please specify a 'ServiceUrl' for the centralconfig service")
		return retval, err
	}

	apiUrl := client.ServiceUrl + "/config/set"

	//	Check to see if we have an application
	if request.Application == "" {
		err := fmt.Errorf("Please specify an 'Application' to set configuration for")
		return retval, err
	}

	//	Serialize our request to JSON:
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(request)
	if err != nil {
		return retval, err
	}

	// Convert bytes to a reader.
	requestJSON := strings.NewReader(requestBytes.String())

	//	Post the JSON to the api url
	res, err := http.Post(apiUrl, "application/json", requestJSON)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return retval, err
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(&retval)
	if err != nil {
		return retval, err
	}

	return retval, nil
}

//	Remove removes a config item
func (client ConfigClient) Remove(request *ConfigItem) error {
	retval := ConfigResponse{}

	//	First, see if we have a service url set:
	if client.ServiceUrl == "" {
		err := fmt.Errorf("Please specify a 'ServiceUrl' for the centralconfig service")
		return err
	}

	apiUrl := client.ServiceUrl + "/config/remove"

	//	Check to see if we have an application
	if request.Application == "" {
		err := fmt.Errorf("Please specify an 'Application' to remove configuration for")
		return err
	}

	//	Check to see if we have a name
	if request.Name == "" {
		err := fmt.Errorf("Please specify a 'Name' to remove configuration for")
		return err
	}

	//	Serialize our request to JSON:
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(request)
	if err != nil {
		return err
	}

	// Convert bytes to a reader.
	requestJSON := strings.NewReader(requestBytes.String())

	//	Post the JSON to the api url
	res, err := http.Post(apiUrl, "application/json", requestJSON)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	//	Decode the return object
	err = json.NewDecoder(res.Body).Decode(&retval)
	if err != nil {
		return err
	}

	return nil
}
