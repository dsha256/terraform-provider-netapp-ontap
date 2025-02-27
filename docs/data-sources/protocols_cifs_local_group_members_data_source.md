---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_protocols_cifs_local_group_members_data_source Data Source - terraform-provider-netapp-ontap"
subcategory: "NAS"
description: |-
  Retrieve CifsLocalGroupMembers data source
---

# Data source netapp-ontap_protocols_cifs_local_group_members_data_source

Retreives protocols cifs local group members configuration

## Example Usage
```terraform
data "netapp-ontap_protocols_cifs_local_group_members_data_source" "protocols_cifs_local_group_members" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  svm_name = "test3"
  group_name = "testme"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `group_name` (String) Local group name
- `svm_name` (String) Svm name

### Read-Only

- `members` (Set of String) List of members


