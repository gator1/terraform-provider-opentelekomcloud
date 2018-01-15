data "opentelekomcloud_rds_flavors_v1" "flavor" {
  region = "eu-de"
  datastore_name = "PostgreSQL"
  datastore_version = "9.5.5"
}

resource "opentelekomcloud_rds_instance_v1" "instance" {
  name = "${var.project}-instance"
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
    id = "${opentelekomcloud_compute_secgroup_v2.secgrp_web.id}"
  }
  dbport = "8635"
  backupstrategy = {
    starttime = "00:00:00"
    keepdays = 0
  }
  dbrtpd = "Huangwei!120521"
}
