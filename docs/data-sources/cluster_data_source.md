---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netapp-ontap_cluster_data_source Data Source - terraform-provider-netapp-ontap"
subcategory: "Cluster"
description: |-
  Cluster data source
---

# netapp-ontap_cluster_data_source (Data Source)

Retrieves the details of a cluster data source

### Related ONTAP commands
cluster show

## Example Usage
```terraform
data "netapp-ontap_cluster_data_source" "cluster" {
  # required to know which system to interface with
  cx_profile_name = "cluster4"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cx_profile_name` (String) Connection profile name

### Read-Only

- `certificate` (Attributes) Certificate (see [below for nested schema](#nestedatt--certificate))
- `contact` (String) Contact information. Example: support@company.com
- `dns_domains` (Set of String) A list of DNS domains.
- `location` (String) Location information
- `management_interfaces` (Attributes Set) A list of network interface (see [below for nested schema](#nestedatt--management_interfaces))
- `name` (String) Cluster name
- `name_servers` (Set of String) The list of IP addresses of the DNS servers. Addresses can be either IPv4 or IPv6 addresses.
- `nodes` (Attributes List) Cluster Nodes (see [below for nested schema](#nestedatt--nodes))
- `ntp_servers` (Set of String) Host name, IPv4 address, or IPv6 address for the external NTP time servers.
- `time_zone` (Attributes) Time zone (see [below for nested schema](#nestedatt--time_zone))
- `version` (Attributes) ONTAP software version (see [below for nested schema](#nestedatt--version))

<a id="nestedatt--certificate"></a>
### Nested Schema for `certificate`

Read-Only:

- `id` (String)


<a id="nestedatt--management_interfaces"></a>
### Nested Schema for `management_interfaces`

Read-Only:

- `id` (String) ID
- `ip` (Attributes) IP address (see [below for nested schema](#nestedatt--management_interfaces--ip))
- `name` (String) Name

<a id="nestedatt--management_interfaces--ip"></a>
### Nested Schema for `management_interfaces.ip`

Read-Only:

- `address` (String) IP address



<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Read-Only:

- `management_ip_addresses` (List of String)
- `name` (String)


<a id="nestedatt--time_zone"></a>
### Nested Schema for `time_zone`

Read-Only:

- `name` (String) Time zone


<a id="nestedatt--version"></a>
### Nested Schema for `version`

Read-Only:

- `full` (String) ONTAP software version

