---
layout: "opentelekomcloud"
page_title: "OpenStack: opentelekomcloud_rds_instance_v1"
sidebar_current: "docs-opentelekomcloud-resource-rds-instance-v1"
description: |-
  Manages rds instance resource within OpenTelekomCloud
---

# opentelekomcloud\_rds\_instance\_v1

Manages rds instance resource within OpenTelekomCloud

## Example Usage

```hcl
data "opentelekomcloud_rds_flavors_v1" "flavor" {
    region = "eu-de"
    datastore_name = "PostgreSQL"
    datastore_version = "9.5.5"
}

resource "opentelekomcloud_compute_secgroup_v2" "secgrp_rds" {
  name        = "secgrp-rds-instance"
  description = "Rds Security Group"
}


resource "opentelekomcloud_networking_network_v2" "network" {
  name = "network_1"
  admin_state_up = "true"
}

resource "opentelekomcloud_networking_subnet_v2" "subnet" {
  name = "subnet_1"
  cidr = "192.168.10.0/24"
  ip_version = 4
  network_id = "${opentelekomcloud_networking_network_v2.network.id}"
}

resource "opentelekomcloud_networking_router_v2" "router" {
  name = "router_1"
}

resource "opentelekomcloud_networking_router_interface_v2" "interface" {
  router_id = "${opentelekomcloud_networking_router_v2.router.id}"
  subnet_id = "${opentelekomcloud_networking_subnet_v2.subnet.id}"
}

resource "opentelekomcloud_rds_instance_v1" "instance" {
  name = "rds-instance"
  datastore {
    type = "PostgreSQL"
    version = "9.5.5"
  }
  flavorref = "${data.opentelekomcloud_rds_flavors_v1.flavor.id}"
  volume {
    type = "COMMON"
    size = 100
  }
  region = "eu-de"
  availabilityzone = "eu-de-01"
  vpc = "${opentelekomcloud_networking_router_v2.router.id}"
  nics {
    subnetid = "${opentelekomcloud_networking_network_v2.network.id}"
  }
  securitygroup {
    id = "${opentelekomcloud_compute_secgroup_v2.secgrp_rds.id}"
  }
  dbport = "8635"
  backupstrategy = {
    starttime = "00:00:00"
    keepdays = 0
  }
  dbrtpd = "Huangwei!120521"
  depends_on = ["opentelekomcloud_networking_router_v2.router",
                "opentelekomcloud_networking_network_v2.network",
                "opentelekomcloud_networking_subnet_v2.subnet",
                "opentelekomcloud_networking_router_interface_v2.interface",
                "opentelekomcloud_compute_secgroup_v2.secgrp_rds"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Specifies the DB instance name. The DB instance name of
    the same type is unique in the same tenant.

* `datastore` - (Required) Specifies database information. The structure is
    described below.

* `flavorref` - (Required) Specifies the specification ID (flavors.id in the
    response message in Obtaining All DB Instance Specifications).

* `volume` - (Required) Specifies the volume information. The structure is described
    below.

* `region` - (Required) Specifies the region ID.

* `availabilityzone` - (Required) Specifies the ID of the AZ.

* `vpc` - (Required) Specifies the VPC ID. For details about how to obtain this
    parameter value, see section "Virtual Private Cloud" in the Virtual Private
    Cloud API Reference.

* `nics` - (Required) Specifies the nics information. For details about how
    to obtain this parameter value, see section "Subnet" in the Virtual Private
    Cloud API Reference. The structure is described below.

* `securitygroup` - (Required) Specifies the security group which the RDS DB
    instance belongs to. The structure is described below.

* `dbport` - (Optional) Specifies the database port number.

* `backupstrategy` - (Optional) Specifies the advanced backup policy. The structure
    is described below.

* `dbrtpd` - (Required) Specifies the password for user root of the database.

* `ha` - (Optional) Specifies the parameters configured on HA and is used when
    creating HA DB instances. The structure is described below.

The `datastore` block supports:

* `type` - (Required) Specifies the DB engine. Currently, MySQL, PostgreSQL, and
    Microsoft SQL Server are supported. The value is MySQL, PostgreSQL, or SQLServer.

* `version` - (Required) Specifies the DB instance version.


The `volume` block supports:

* `type` - (Required) Specifies the volume type. Valid value:
    It must be COMMON (SATA) or ULTRAHIGH (SSD) and is case-sensitive.

* `size` - (Required) Specifies the volume size.
    Its value must be a multiple of 10 and the value range is 100 GB to 2000 GB.

The `nics` block supports:

* `subnetId` - (Required) Specifies the subnet ID obtained from the VPC.

The `securitygroup ` block supports:

* `id` - (Required) Specifies the ID obtained from the securitygroup.

The `backupstrategy ` block supports:

* `starttime` - (Optional) Indicates the backup start time that has been set.
    The backup task will be triggered within one hour after the backup start time.
    Valid value: The value cannot be empty. It must use the hh:mm:ss format and
    must be valid. The current time is the UTC time.

* `keepdays` - (Optional) Specifies the number of days to retain the generated backup files.
    Its value range is 0 to 35. If this parameter is not specified or set to 0, the
    automated backup policy is disabled.

The `ha` block supports:

* `enable` - (Optional) Specifies the configured parameters on the HA.
    Valid value: The value is true or false. The value true indicates creating
    HA DB instances. The value false indicates creating a single DB instance.

* `replicationmode` - (Optional) Specifies the replication mode for the standby DB instance.
    The value cannot be empty.
    For MySQL, the value is async or semisync.
    For PostgreSQL, the value is async or sync.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `flavorref` - See Argument Reference above.
* `volume` - See Argument Reference above.
* `availabilityzone` - See Argument Reference above.
* `vpc` - See Argument Reference above.
* `nics` - See Argument Reference above.
* `securitygroup` - See Argument Reference above.
* `dbport` - See Argument Reference above.
* `backupstrategy` - See Argument Reference above.
* `dbrtpd` - See Argument Reference above.
* `ha` - See Argument Reference above.
* `status` - Indicates the DB instance status.
* `hostname` - Indicates the instance connection address. It is a blank string.
* `type` - Indicates the DB instance type, which can be master or readreplica.
* `created` - Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.
* `updated` - Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ.