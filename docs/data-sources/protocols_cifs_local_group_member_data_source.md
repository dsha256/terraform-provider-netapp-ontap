---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_protocols_cifs_local_group_member_data_source Data Source - terraform-provider-netapp-ontap"
subcategory: "NAS"
description: |-
  Retrieve CifsLocalGroupMember data source
---

# Data source netapp-ontap_protocols_cifs_local_group_member_data_source

Retreives protocols cifs local group member configuration

## Example Usage
```terraform
data "netapp-ontap_protocols_cifs_local_group_member_data_source" "protocols_cifs_local_group_member" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
  svm_name = "test3"
  group_name = "testme"
  member = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name
- `group_name` (String) Local group name
- `member` (String) Member name
- `svm_name` (String) Svm name


