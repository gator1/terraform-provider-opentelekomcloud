package datastores

import "github.com/gophercloud/gophercloud"

func listURL(c *gophercloud.ServiceClient, dataStoreName string) string {
	return c.ServiceURL("datastores", dataStoreName, "versions")
}
