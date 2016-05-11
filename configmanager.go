package configmanager

import (
	"time"
)

//	ConfigItem represents a configuration item
type ConfigItem struct {
	Id          int64     `sql:"id" json:"id"`
	Application string    `sql:"application" json:"application"`
	Machine     string    `sql:"machine" json:"machine"`
	Name        string    `sql:"name" json:"name"`
	Value       string    `sql:"value" json:"value"`
	LastUpdated time.Time `sql:"updated" json:"updated"`
}

//	ConfigResponse represents a response from the centralconfig service
type ConfigResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//	ConfigService encapsulates account (user) based operations
//	This allows us to create a testable service layer.  See
//	https://github.com/tonyhb/tonyhb.com/blob/master/posts/Building%20a%20testable%20Golang%20database%20layer.md
//	for more information
type ConfigService interface {

	//	Get a specific config item
	Get(name string) ConfigResponse

	//	Get a specific config string
	GetString(name string, defaultValue string) string

	//	Create / update a config item
	Set(c *ConfigItem) (ConfigItem, error)

	//	Get all config items for the given application
	GetAllForApplication(application string) ([]ConfigItem, error)

	//	Get all config items for all applications (including global)
	GetAll() ([]ConfigItem, error)

	//	Get all applications (including global)
	GetAllApplications() ([]string, error)

	//	Remove a config item
	Remove(c *ConfigItem) error
}
