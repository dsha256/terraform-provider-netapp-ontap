---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_protocols_san_igroup_resource Resource - terraform-provider-netapp-ontap"
subcategory: "NAS"
description: |-
  ProtocolsSanIgroup resource
---

# netapp-ontap_protocols_san_igroup_resource (Resource)

Create/Modify/Delete a protocols_san_igroup resource

### Related ONTAP commands
```commandline
* lun igroup create
* lun igroup modify
* lun igroup delete
```

## Supported Platforms
* On-perm ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage
```
# Create a protocols_san_igroup
resource "netapp-ontap_protocols_san_igroup_resource" "protocols_san_igroups" {
  # required to know which system to interface with
  cx_profile_name = "cluster2"
  name = "test1"
  svm = {
    name = "test"
  }
  os_type = "linux"
  comment = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `name` (String) Existing SVM in which to create the initiator group.
- `os_type` (String) Operating system of the initiator group's initiators.
- `svm` (Attributes) SVM details for ProtocolsSanLunMaps (see [below for nested schema](#nestedatt--svm))

### Optional

- `comment` (String) Comment
- `igroups` (Attributes Set) List of initiator groups (see [below for nested schema](#nestedatt--igroups))
- `initiators` (Attributes Set) List of initiators (see [below for nested schema](#nestedatt--initiators))
- `portset` (Attributes) Required ONTAP 9.9 or greater. The portset to which the initiator group is bound. Binding the initiator group to a portset restricts the initiators of the group to accessing mapped LUNs only through network interfaces in the portset. (see [below for nested schema](#nestedatt--portset))
- `protocol` (String) If not specified, the default protocol is mixed.

### Read-Only

- `id` (String) Igroup UUID

<a id="nestedatt--svm"></a>
### Nested Schema for `svm`

Required:

- `name` (String) name of the SVM


<a id="nestedatt--igroups"></a>
### Nested Schema for `igroups`

Required:

- `name` (String) Initiator group name


<a id="nestedatt--initiators"></a>
### Nested Schema for `initiators`

Required:

- `name` (String) Initiator name


<a id="nestedatt--portset"></a>
### Nested Schema for `portset`

Required:

- `name` (String) Portset name

## Import
This resource supports import, which allows you to import existing protocols_san_igroup into the state of this resource.
Import require a unique ID composed of the protocols_san_igroup name, svm_name and connection profile, separated by a comma.

id = `name`, `svm.name`, `cx_profile_name`

### Terraform Import

For example
```shell
 terraform import netapp-ontap_protocols_san_igroup_resource.example name,svm_name,cluster5
```
!> The terraform import CLI command can only import resources into the state. Importing via the CLI does not generate configuration. If you want to generate the accompanying configuration for imported resources, use the import block instead.

### Terrafomr Import Block
This requires Terraform 1.5 or higher, and will auto create the configuration for you

First create the block
```terraform
import {
  to = netapp-ontap_protocols_san_igroup_resource.protocols_san_igroup_import
  id = "name,svm_name,cluster5"
}
```
Next run, this will auto create the configuration for you
```shell
terraform plan -generate-config-out=generated.tf
```
This will generate a file called generated.tf, which will contain the configuration for the imported resource
```terraform
# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.
# __generated__ by Terraform from "name,svm_name,cluster5"
resource "netapp-ontap_protocols_san_igroup_resource" "protocols_san_igroup_import" {
  cx_profile_name = "cluster4"
  id = "abcd"
  igroups = null
  initiators = null
  svm = {
    name = "test"
  }
  name = "test"
  os_type = "windows"
  portset = {
    name = "test"
  }
  protocol = "mixed"
}
```

