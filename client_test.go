package configmanager_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/cagedtornado/centralconfig-go"
)

//	Gets the service connection information from the environment
func getCentralConfigInfo() configmanager.ConfigClient {

	return configmanager.ConfigClient{
		Application: "WickedCool",
		ServiceUrl:  os.Getenv("centralconfig_service_url")}

}

//	We should be able to get a blank configitem back
func TestConfigManager_Get_ReturnsConfigResponse(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()
	expected := ""

	//	Act
	response, err := config.Get("TestItem42")

	//	Assert
	if err != nil {
		t.Errorf("Get failed: Can't get config: %v", err)
	}

	if response.Status != http.StatusOK {
		t.Errorf("Get failed: ConfigResponse isn't OK: %v Response is %v", err, response.Status)
	}

	if response.Data.Value != expected {
		t.Errorf("Get failed: Config value '%v' isn't expected.  Should be '%v'", response.Data.Value, expected)
	}

}
