package flavors

import "github.com/gophercloud/gophercloud"

type Flavor struct {
    ID               string                 `json:"id"`
    Name             string                 `json:"name"`
    Ram              int                    `json:"ram"`
    SpecCode         string                 `json:"specCode"`
}

type ListResult struct {
	gophercloud.Result
}

func (lr ListResult) Extract() ([]Flavor, error) {
	var a struct {
        Flavors []Flavor `json:"flavors"`
	}
	err := lr.Result.ExtractInto(&a)
	return a.Flavors, err
}
