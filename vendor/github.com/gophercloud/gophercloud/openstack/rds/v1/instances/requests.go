package instances

import (
	"github.com/gophercloud/gophercloud"
)

var RequestOpts gophercloud.RequestOpts = gophercloud.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

//CreateOpsBuilder is used for creating instance parameters.
//any struct providing the parameters should implement this interface
type CreateOpsBuilder interface {
	ToInstanceCreateMap() (map[string]interface{}, error)
}

type UpdateOpsBuilder interface {
	ToInstanceUpdateMap() (map[string]interface{}, error)
}

type UpdatePolicyOpsBuilder interface {
    ToInstanceUpdatePolicyMap() (map[string]interface{}, error)
}

type UpdateFlavorOpsBuilder interface {
    ToInstanceFlavorUpdateMap() (map[string]interface{}, error)
}

type UpdateOps struct {
	Volume map[string]interface{} `json:"volume"`
}

type UpdatePolicyOps struct {
    StartTime string `json:"starttime"`
    KeepDays  int    `json:"keepday"`
}

type UpdateFlavorOps struct {
	FlavorRef string `json:"flavorRef"`
}

//CreateOps is a struct that contains all the parameters.
type CreateOps struct {

	Name string `json:"name" required:"true"`

	DataStore DataStoreOps `json:"datastore,omitempty"`

	FlavorRef string `json:"flavorRef" required:"true"`

	Volume VolumeOps `json:"volume,omitempty"`

	Region string `json:"region,omitempty"`

	AvailabilityZone string `json:"availabilityZone,omitempty"`

	Vpc string `json:"vpc,omitempty"`

	Nics NicsOps `json:"nics,omitempty"`

	SecurityGroup SecurityGroupOps `json:"securityGroup,omitempty"`

	DbPort string `json:"dbPort,omitempty"`

	BackupStrategy BackupStrategyOps `json:"backupStrategy,omitempty"`

	DbRtPd string `json:"dbRtPd,omitempty"`

	Ha HaOps `json:"ha,omitempty"`

	ReplicaOf string `json:"replicaOf,omitempty"`
}

type DataStoreOps struct {
	Type    string `json:"type" required:"true"`
	Version string `json:"version" required:"true"`
}

type VolumeOps struct {
	Type string `json:"type" required:"true"`
	Size int    `json:"size" required:"true"`
}

type NicsOps struct {
	SubnetId string `json:"subnetId" required:"true"`
}

type SecurityGroupOps struct {
	Id string `json:"id" required:"true"`
}

type BackupStrategyOps struct {
	StartTime string `json:"startTime" required:"true"`
	KeepDays  int    `json:"keepDays,omitempty"`
}

type HaOps struct {
	Enable          bool   `json:"enable" required:"true"`
	ReplicationMode string `json:"replicationMode" required:"true"`
}


func (ops CreateOps) ToInstanceCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(ops, "instance")
}

func (ops UpdateOps) ToInstanceUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(ops, "resize")
}

func (ops UpdatePolicyOps) ToInstanceUpdatePolicyMap() (map[string]interface{}, error) {
    return gophercloud.BuildRequestBody(ops, "policy")
}

func (ops UpdateFlavorOps) ToInstanceFlavorUpdateMap() (map[string]interface{}, error) {
    return gophercloud.BuildRequestBody(ops, "resize")
}

//Create a instance with given parameters.
func Create(client *gophercloud.ServiceClient, ops CreateOpsBuilder) (r CreateResult) {
	b, err := ops.ToInstanceCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})

	return
}

func UpdateVolumeSize(client *gophercloud.ServiceClient, ops UpdateOpsBuilder, id string) (r UpdateResult) {
	b, err := ops.ToInstanceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

func UpdatePolicy(client *gophercloud.ServiceClient, ops UpdatePolicyOpsBuilder, id string) (r UpdateResult) {
    b, err := ops.ToInstanceUpdatePolicyMap()
    if err != nil {
        r.Err = err
        return
    }

    _, r.Err = client.Put(updatePolicyURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
        OkCodes:     []int{200},
        MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
    })
    return
}


func UpdateFlavorRef(client *gophercloud.ServiceClient, ops UpdateFlavorOpsBuilder, id string) (r UpdateResult) {
	b, err := ops.ToInstanceFlavorUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

//delete a instance via id
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	RequestOpts.OkCodes = []int{202}
	RequestOpts.JSONBody = nil
	_, r.Err = client.Delete(deleteURL(client, id), &gophercloud.RequestOpts{
		OkCodes:     []int{202},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

//get a instance with detailed information by id
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {

	_, r.Err = client.Get(getURL(client, id), &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

//list all the instances
func List(client *gophercloud.ServiceClient) (r ListResult) {

	_, r.Err = client.Get(listURL(client), &r.Body, &gophercloud.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
