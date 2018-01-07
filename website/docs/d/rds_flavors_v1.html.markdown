---
layout: "opentelekomcloud"
page_title: "OpenTelekomCloud: opentelekomcloud_rds_flavors_v1"
sidebar_current: "docs-opentelekomcloud-datasource-rds-flavors-v1"
description: |-
  Get the flavor information on an OpenTelekomCloud rds service.
---

# opentelekomcloud\_rds\_flavors\_v1

Use this data source to get the ID of an available OpenTelekomCloud rds flavor.

## Example Usage

```hcl
data "opentelekomcloud_rds_flavors_v1" "flavor" {
    region = "eu-de"
    datastore_name = "PostgreSQL"
    datastore_version = "9.5.5"
    speccode = "rds.pg.s1.medium"
}
```

## Argument Reference

* `region` - (Required) The region in which to obtain the V1 rds client.

* `datastore_name` - (Required) The datastore name of the rds.

* `datastore_version` - (Required) The datastore version of the rds.

* `speccode` - (Optional) The spec code of a rds flavor.


## Attributes Reference

`id` is set to the ID of the found rds flavor. In addition, the following attributes
are exported:

* `region` - See Argument Reference above.
* `datastore_name` - See Argument Reference above.
* `datastore_version` - See Argument Reference above.
* `speccode` - See Argument Reference above.
* `name` - The name of the rds flavor.
* `ram` - The name of the rds flavor.
