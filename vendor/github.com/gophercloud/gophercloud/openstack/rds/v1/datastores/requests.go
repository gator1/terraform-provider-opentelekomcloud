package datastores

import (
	"github.com/gophercloud/gophercloud"
)

var RequestOpts gophercloud.RequestOpts = gophercloud.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

//list the version informations about a specified type of database
func List(client *gophercloud.ServiceClient, dataStoreName string) (r ListResult) {

	_, r.Err = client.Get(listURL(client, dataStoreName), &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
