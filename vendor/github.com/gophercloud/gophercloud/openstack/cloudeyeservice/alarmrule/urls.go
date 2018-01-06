package alarmrule

import "github.com/gophercloud/gophercloud"

const (
	rootPath = "alarms"
)

func rootURL(c *gophercloud.ServiceClient1) string {
	return c.ServiceURL(c.ProjectID, rootPath)
}

func resourceURL(c *gophercloud.ServiceClient1, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, id)
}

func actionURL(c *gophercloud.ServiceClient1, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, id, "action")
}
