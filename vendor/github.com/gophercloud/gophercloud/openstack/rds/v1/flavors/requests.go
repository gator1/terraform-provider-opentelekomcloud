package flavors

import (
	"github.com/gophercloud/gophercloud"
)

var RequestOpts gophercloud.RequestOpts = gophercloud.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

//list the flavors informations about a specified id of database
func List(client *gophercloud.ServiceClient, dataStoreID string, region string) (r ListResult) {

	_, r.Err = client.Get(listURL(client, dataStoreID, region), &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
