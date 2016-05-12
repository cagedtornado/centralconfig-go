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

//	We should be able to get a
func TestConfigManager_Get_ReturnsConfigResponse(t *testing.T) {
	//	Arrange
	config := getCentralConfigInfo()

	//	Act
	configItem, err := config.Get("TestItem42")

	//	Assert
	if err != nil {
		t.Errorf("Get failed: Can't get config: %v", err)
	}

	if configItem.Status != http.StatusOK {
		t.Errorf("Get failed: ConfigResponse isn't OK: %v Response is %v", err, configItem.Status)
	}

}
