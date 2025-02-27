---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ONTAP: IP Interface"
subcategory: "Networking"
description: |-
  IPInterface resource
---

# Resource IP Interface

Create/Update/Delete an IPInterface resource

### Related ONTAP commands
```commandline
* network interface create
* network interface modify
* network interface delete
```

## Supported Platforms
* On-perm ONTAP system 9.6 or higher

## Example Usage

```terraform
resource "netapp-ontap_networking_ip_interface_resource" "example" {
	cx_profile_name = "cluster4"
	name = "test-interface"
	svm_name = "carchi-test"
  	ip = {
    	address = "10.10.10.10"
    	netmask = 18
    }
  	location = {
    	home_port = "e0d"
    	home_node = "ontap_cluster_1-01"
  	}
}
```



<!-- schema generated by tfplugindocs -->
## Argument Reference

### Required

- `cx_profile_name` (String) Connection profile name
- `ip` (Attributes) (see [below for nested schema](#nestedatt--ip))
- `location` (Attributes) (see [below for nested schema](#nestedatt--location))
- `name` (String) IPInterface name
- `svm_name` (String) IPInterface svm name

### Read-Only

- `id` (String) IPInterface UUID

<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Required:

- `address` (String) IPInterface IP address
- `netmask` (Number) IPInterface IP netmask


<a id="nestedatt--location"></a>
### Nested Schema for `location`

Required:

- `home_node` (String) IPInterface home node
- `home_port` (String) IPInterface home port

## Import
This Resource supports import, which allows you to import existing network ip interface into the state of this resoruce.
Import require a unique ID composed of the interface name, svm_name and cx_profile_name, separated by a comma.
 id = `name`,`svm_name`,`cx_profile_name`
 ### Terraform Import
 For example
 ```shell
  terraform import netapp-ontap_networking_ip_interface_resource.example if1,svm1,cluster4
 ```

!> The terraform import CLI command can only import resources into the state. Importing via the CLI does not generate configuration. If you want to generate the accompanying configuration for imported resources, use the import block instead.

### Terrafomr Import Block
This requires Terraform 1.5 or higher, and will auto create the configuration for you

First create the block
```terraform
import {
  to = netapp-ontap_networking_ip_interface_resource.if_import
  id = "if1,svm1,cluster4"
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
# __generated__ by Terraform from "if1,svm1,cluster4"
resource "netapp-ontap_networking_ip_interface_resource" "if1_import" {
  cx_profile_name = "cluster4"
  name       = "if1"
  svm_name   = "svm1"
}
```
