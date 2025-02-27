---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ONTAP: storage volume snapshot"
subcategory: "Storage"
description: |-
  Snapshot resource
---

# Resource Volume Snapshot

Create/Modify/Delete a Snapshot resource

### Related ONTAP commands
```commandline
* snapshot create
* snapshot modify
* snapshot delete
```

## Supported Platforms
* On-perm ONTAP system 9.6 or higher
* Amazon FSx for NetApp ONTAP

## Example Usage

```
resource "netapp-ontap_storage_volume_snapshot_resource" "example" {
  cx_profile_name = "cluster4"
  name = "snaptest"
  volume_name = "tf_test_root"
  svm_name = "tf-test"
}
```


<!-- schema generated by tfplugindocs -->
## Argument Reference

### Required

- `cx_profile_name` (String) Connection profile name
- `name` (String) Snapshot name
- `svm_nmae` (String) The name of the SVM the snapshot is on
- `volume_name` (String) The name of the volume the snapshot is on

### Optional

- `comment` (String) Comment
- `expiry_time` (String) Snapshot copies with an expiry time set are not allowed to be deleted until the retetion time is reached
- `snaplock_expiry_time` (String) Expiry time for Snapshot copy locking enabled volumes
- `snapmirror_label` (String) Label for SnapMirror Operations

### Read-Only

- `id` (String) storage/volumes/snapshots identifier

## Import
This Resource supports import, which allows you to import existing snapshot into the state of this resoruce.
Import require a unique ID composed of the snapshot name, volume_name, svm_name and cx_profile_name, separated by a comma.

 id = `name`,`volume_name`,`svm_name`,`cx_profile_name`

 ### Terraform Import

 For example
 ```shell
  terraform import netapp-ontap_storage_volume_snapshot_resource.example snapshotname,vol2,svm1,cluster4
 ```

!> The terraform import CLI command can only import resources into the state. Importing via the CLI does not generate configuration. If you want to generate the accompanying configuration for imported resources, use the import block instead.

### Terrafomr Import Block
This requires Terraform 1.5 or higher, and will auto create the configuration for you

First create the block
```terraform
import {
  to = netapp-ontap_storage_volume_snapshot_resource.snapshot_import
  id = "snapshot1,vol1,svm1,cluster4"
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
# __generated__ by Terraform from "snapshot1,vol1,svm1,cluster4"
resource "netapp-ontap_storage_volume_snapshot_resource" "snapshot_import" {
  cx_profile_name = "cluster4"
  name       = "snapshot1"
  volume_name       = "vol1"
  svm_name = "svm1"
}
```