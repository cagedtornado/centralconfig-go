package configmanager

//	The centralconfig client
type ConfigClient struct {
	Application string
	Machine     string
}

//	Get a specific config item
func (client ConfigClient) Get(name string) ConfigResponse {
	retval := ConfigResponse{}

	return retval
}
