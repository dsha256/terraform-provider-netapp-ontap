resource "netapp-ontap_protocols_san_lun-maps_resource" "protocols_san_lun-maps" {
  # required to know which system to interface with
  cx_profile_name = "cluster2"
  svm = {
    name = "carchi-test"
  }
  lun = {
    name = "/vol/lunTest/ACC-import-lun"
  }
  igroup = {
    name = "test"
  }
  # logical_unit_number = 1
}
