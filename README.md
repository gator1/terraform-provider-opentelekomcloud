Terraform OpenTelekomCloud Provider
============================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Terraform-Provider-OpenTelekomcloud is a Terraform Provider for OpenTelekomcloud.
It is based on [Terraform-Provider-OpenStack](https://github.com/terraform-providers/terraform-provider-openstack)
which is a standard Terraform OpenStack Provider. Terraform-Provider-OpenTelekomcloud has renamed OpenStack with OpenTelekomcloud
and added some enhancements to interact with the many resources supported by OpenTelekomcloud.

## Enhancements

- **ces**: cloud eye service
- **rds**: relational database service
- **obs**: object storage service
- **elb**: elastic load balance service
- **smn**: simple message notification service

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)


Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/gator1/terraform-provider-opentelekomcloud`

```sh
$ go get github.com/gator1/terraform-provider-opentelekomcloud
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/gator1/terraform-provider-opentelekomcloud
$ make build
```

## Exact steps on clean Ubuntu 16.04

```sh
# prerequisites are sudo privileges, unzip, make, wget and git.  Use apt install if missing.
$ wget https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.9.1.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin # You should put in your .profile or .bashrc
$ go version # to verify it runs and version #
$ go get github.com/gator1/terraform-provider-opentelekomcloud
$ cd ~/go/src/github.com/gator1/terraform-provider-opentelekomcloud/
$ make build
$ export PATH=$PATH:~/go/bin # You should put in your .profile or .bashrc
$ wget https://releases.hashicorp.com/terraform/0.10.7/terraform_0.10.7_linux_amd64.zip
$ unzip terraform_0.10.7_linux_amd64.zip
$ mv terraform ~/go/bin
$ terraform version # to verify it runs and version #
$ vi test.tf # paste in Quick Start contents, fix authentication information
$ terraform init
$ terraform plan
$ terraform apply # Should all work if everything is correct.

```

## Quick Start

```hcl
# Configure the OpenTelekomCloud Provider
# This will work with a single defined/default network, otherwise you need to specify network
# to fix errrors about multiple networks found.
provider "opentelekomcloud" {
  user_name   = "user"
  tenant_name = "tenant"
  domain_name = "domain"
  password    = "pwd"
  auth_url    = "https://iam.eu-de.otc.t-systems.com/v3"
  region      = "eu-de"
}

# Create a web server
resource "opentelekomcloud_compute_instance_v2" "test-server" {
  name		  = "test-server"
  image_name  = "Standard_CentOS_7_latest"
  flavor_name = "s1.medium"
}
```

### Full Example
----------------------
Please see full example at https://github.com/gator1/terraform-provider-opentelekomcloud/tree/master/examples, 
you must fill in the required variables in variables.tf.

Using the provider
----------------------
Please see the documentation at [provider usage](website/docs/index.html.markdown).

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-opentelekomcloud
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## License

Terraform-Provider-OpenTelekomcloud is under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.

