package configmanager_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/cagedtornado/centralconfig-go"
)

var (
	unitTestAppName    = "UnitTest"
	defaultMachineName = "Kali"
)

//	Gets the service connection information from the environment
func getCentralConfigInfo() configmanager.ConfigClient {

	return configmanager.ConfigClient{
		Application: unitTestAppName,
		ServiceUrl:  os.Getenv("centralconfig_service_url")}

}

//	We should be able to get a blank configitem back
func TestConfigManager_Get_ReturnsConfigResponse(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()
	expected := ""
	if config.ServiceUrl == "" {
		t.Fatalf("Oops!  It looks like you need to set the 'centralconfig_service_url' environment variable to the service endpoint of the centralconfig service.")
	}

	//	Act
	response, err := config.Get("TestItem42")

	//	Assert
	if err != nil {
		t.Errorf("Can't get config: %v", err)
	}

	if response.Status != http.StatusOK {
		t.Errorf("ConfigResponse isn't OK: %v Response is %v", err, response.Status)
	}

	if response.Data.Value != expected {
		t.Errorf("Config value '%v' isn't expected.  Should be '%v'", response.Data.Value, expected)
	}

}

//	We should be able to set an item
func TestConfigManager_Set_ReturnsConfigResponse(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()

	request := configmanager.ConfigItem{
		Application: unitTestAppName,
		Name:        "TestItem42",
		Value:       "SetFirstTime"}

	expected := "SetFirstTime"

	//	Act
	response, err := config.Set(&request)
	defer config.Remove(&request)

	//	Assert
	if err != nil {
		t.Errorf("Can't set config: %v", err)
	}

	if response.Status != http.StatusOK {
		t.Errorf("ConfigResponse isn't OK: %v Response is %v:%v", err, response.Status, response.Message)
	}

	if response.Data.Value != expected {
		t.Errorf("Config value '%v' isn't expected.  Should be '%v'", response.Data.Value, expected)
	}
}

//	We should be able to set an item with a machine name
func TestConfigManager_Set_WithMachine_ReturnsConfigResponse(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()

	//	Set the hostname to the os Hostname or
	//	the default hostname
	hostName, err := os.Hostname()
	if err != nil {
		hostName = defaultMachineName
	}

	request := configmanager.ConfigItem{
		Application: unitTestAppName,
		Machine:     hostName,
		Name:        "TestItem42",
		Value:       "Set_WithMachine"}

	expected := "Set_WithMachine"

	//	Act
	response, err := config.Set(&request)
	defer config.Remove(&request)

	//	Assert
	if err != nil {
		t.Errorf("Can't set config: %v", err)
	}

	if response.Status != http.StatusOK {
		t.Errorf("ConfigResponse isn't OK: %v Response is %v:%v", err, response.Status, response.Message)
	}

	if response.Data.Value != expected {
		t.Errorf("Config value '%v' isn't expected.  Should be '%v'", response.Data.Value, expected)
	}

	if response.Data.Machine != hostName {
		t.Errorf("Config machine '%v' ins't expected.  Should be '%v'", response.Data.Machine, hostName)
	}
}

//	We should be able to set and remove an item
func TestConfigManager_Remove_Successful(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()

	request := configmanager.ConfigItem{
		Application: unitTestAppName,
		Name:        "TestItem42",
		Value:       "Remove"}

	//	Act
	config.Set(&request)
	err := config.Remove(&request)

	//	Assert
	if err != nil {
		t.Errorf("Can't remove config: %v", err)
	}
}

//	We should be able to set and remove an item
func TestConfigManager_RemoveWithMachine_Successful(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()

	//	Set the hostname to the os Hostname or
	//	the default hostname
	hostName, err := os.Hostname()
	if err != nil {
		hostName = defaultMachineName
	}

	requests := []configmanager.ConfigItem{}

	//	--- Config item without a machine name
	requests = append(requests, configmanager.ConfigItem{
		Application: unitTestAppName,
		Name:        "TestItem42",
		Value:       "RemoveWithMachine"})

	//	--- Config item with machine name
	requests = append(requests, configmanager.ConfigItem{
		Application: unitTestAppName,
		Machine:     hostName,
		Name:        "TestItem42",
		Value:       "RemoveWithMachine"})

	//	Act
	for _, item := range requests {
		config.Set(&item)
	}

	for _, item := range requests {
		err = config.Remove(&item)

		//	Assert
		if err != nil {
			t.Errorf("Can't remove config: %v", err)
		}
	}
}

//	We should be able to set multiple items and get them all back
func TestConfigManager_GetAll_ReturnsConfigResponseMultiple(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()

	//	Set the hostname to the os Hostname or
	//	the default hostname
	hostName, err := os.Hostname()
	if err != nil {
		hostName = defaultMachineName
	}

	requests := []configmanager.ConfigItem{}

	//	--- Config item without a machine name
	requests = append(requests, configmanager.ConfigItem{
		Application: unitTestAppName,
		Name:        "TestItem42",
		Value:       "GetAll"})

	//	--- Config item with machine name
	requests = append(requests, configmanager.ConfigItem{
		Application: unitTestAppName,
		Machine:     hostName,
		Name:        "TestItem42",
		Value:       "GetAll"})

	//	Clean up:
	defer func() {
		for _, item := range requests {
			config.Remove(&item)
		}
	}()

	//	Act
	for _, item := range requests {
		config.Set(&item)
	}

	response, err := config.GetAll()

	//	Assert
	if err != nil {
		t.Errorf("Can't get all configs: %v", err)
	}

	if response.Status != http.StatusOK {
		t.Errorf("ConfigResponse isn't OK: %v Response is %v:%v", err, response.Status, response.Message)
	}

	if len(response.Data) == 0 {
		t.Errorf("Got '%v' items back.  Should be at least '%v' items", len(response.Data), len(requests))
	}
}
