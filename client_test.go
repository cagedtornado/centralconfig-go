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
	t.Logf("Using config client information: %v", config)

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

	//	Assert
	if err != nil {
		t.Errorf("Can't get config: %v", err)
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
		Value:       "SetFirstTime"}

	expected := "SetFirstTime"

	//	Act
	response, err := config.Set(&request)

	//	Assert
	if err != nil {
		t.Errorf("Can't get config: %v", err)
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
