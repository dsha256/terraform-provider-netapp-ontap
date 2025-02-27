package interfaces

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/mitchellh/mapstructure"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/restclient"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/utils"
)

// ProtocolsCIFSShareGetDataModelONTAP describes the GET record data model using go types for mapping.
type ProtocolsCIFSShareGetDataModelONTAP struct {
	Name string `mapstructure:"name"`
	UUID string `mapstructure:"uuid"`

	Acls                  []AclsGet `mapstructure:"acls"`
	ChangeNotify          bool      `mapstructure:"change_notify"`
	Comment               string    `mapstructure:"comment"`
	ContinuouslyAvailable bool      `mapstructure:"continuously_available"`
	DirUmask              int64     `mapstructure:"dir_umask"`
	Encryption            bool      `mapstructure:"encryption"`
	FileUmask             int64     `mapstructure:"file_umask"`
	ForceGroupForCreate   string    `mapstructure:"force_group_for_create"`
	HomeDirectory         bool      `mapstructure:"home_directory"`
	NamespaceCaching      bool      `mapstructure:"namespace_caching"`
	NoStrictSecurity      bool      `mapstructure:"no_strict_security"`
	OfflineFiles          string    `mapstructure:"offline_files"`
	Oplocks               bool      `mapstructure:"oplocks"`
	Path                  string    `mapstructure:"path"`
	ShowSnapshot          bool      `mapstructure:"show_snapshot"`
	UnixSymlink           string    `mapstructure:"unix_symlink"`
	VscanProfile          string    `mapstructure:"vscan_profile"`
}

// Acls describes the acls data model using go types for mapping in ProtocolsCIFSShareResourceBodyDataModelONTAP.
type Acls struct {
	Permission  string `json:"permission"` // Because Acls is nested in ProtocolsCIFSShareResourceBodyDataModelONTAP, when it unmarshal from API to go struct, 'mapstructure' does not work, and result in "Permission" instead of "permission". So, we need to use "json" tag to map the field.
	Type        string `json:"type"`
	UserOrGroup string `json:"user_or_group"`
}

// AclsGet describes the acls data model using go types for mapping in ProtocolsCIFSShareGetDataModelONTAP.
type AclsGet struct {
	Permission  string `mapstructure:"permission"` // ProtocolsCIFSShareGetDataModelONTAP needs 'mapstructure'
	Type        string `mapstructure:"type"`
	UserOrGroup string `mapstructure:"user_or_group"`
}

// ProtocolsCIFSShareResourceBodyDataModelONTAP describes the body data model using go types for mapping.
type ProtocolsCIFSShareResourceBodyDataModelONTAP struct {
	Name                  string `mapstructure:"name,omitempty"` // can't be present in update, so omit empty.
	SVM                   svm    `mapstructure:"svm"`
	Acls                  []Acls `mapstructure:"acls,omitempty"` // API complains if this is not omit empty
	ChangeNotify          bool   `mapstructure:"change_notify"`
	Comment               string `mapstructure:"comment,omitempty"` // API complains if this is not omit empty
	ContinuouslyAvailable bool   `mapstructure:"continuously_available"`
	DirUmask              int64  `mapstructure:"dir_umask"`
	Encryption            bool   `mapstructure:"encryption"`
	FileUmask             int64  `mapstructure:"file_umask"`
	ForceGroupForCreate   string `mapstructure:"force_group_for_create"`
	HomeDirectory         bool   `mapstructure:"home_directory,omitempty"` // can't be present in update, so omit empty.
	NamespaceCaching      bool   `mapstructure:"namespace_caching"`
	NoStrictSecurity      bool   `mapstructure:"no_strict_security"`
	OfflineFiles          string `mapstructure:"offline_files,omitempty"` // API complains if this is not omit empty
	Oplocks               bool   `mapstructure:"oplocks"`
	Path                  string `mapstructure:"path,omitempty"` // can't be present in update, so omit empty.
	ShowSnapshot          bool   `mapstructure:"show_snapshot"`
	UnixSymlink           string `mapstructure:"unix_symlink,omitempty"`  // API complains if this is not omit empty
	VscanProfile          string `mapstructure:"vscan_profile,omitempty"` // API complains if this is not omit empty
}

// ProtocolsCIFSShareDataSourceFilterModel describes the data source data model for queries.
type ProtocolsCIFSShareDataSourceFilterModel struct {
	Name    string `mapstructure:"name"`
	SVMName string `mapstructure:"svm.name"`
}

// GetProtocolsCIFSShareByName to get protocols_cifs_share info
func GetProtocolsCIFSShareByName(errorHandler *utils.ErrorHandler, r restclient.RestClient, name string, svmName string) (*ProtocolsCIFSShareGetDataModelONTAP, error) {
	// api := fmt.Sprintf("protocols/cifs/shares/%s/%s", svmName, name)
	api := "protocols/cifs/shares"
	query := r.NewQuery()
	query.Add("name", name)
	query.Add("svm.name", svmName)
	query.Fields([]string{"name", "svm.name", "unix_symlink", "dir_umask", "file_umask", "acls", "home_directory", "force_group_for_create", "no_strict_security", "oplocks", "volume", "change_notify", "path", "encryption", "vscan_profile", "offline_files", "comment", "show_snapshot", "continuously_available", "namespace_caching"})
	statusCode, response, err := r.GetNilOrOneRecord(api, query, nil)
	if err == nil && response == nil {
		err = fmt.Errorf("no response for GET %s", api)
	}
	if err != nil {
		return nil, errorHandler.MakeAndReportError("error reading protocols_cifs_share info", fmt.Sprintf("error on GET %s: %s, statusCode %d", api, err, statusCode))
	}

	var dataONTAP ProtocolsCIFSShareGetDataModelONTAP
	if err := mapstructure.Decode(response, &dataONTAP); err != nil {
		return nil, errorHandler.MakeAndReportError(fmt.Sprintf("failed to decode response from GET %s", api),
			fmt.Sprintf("error: %s, statusCode %d, response %#v", err, statusCode, response))
	}
	tflog.Debug(errorHandler.Ctx, fmt.Sprintf("Read protocols_cifs_share: %#v", dataONTAP))
	return &dataONTAP, nil
}

// GetProtocolsCIFSShares to get protocols_cifs_share info for all resources matching a filter
func GetProtocolsCIFSShares(errorHandler *utils.ErrorHandler, r restclient.RestClient, filter *ProtocolsCIFSShareDataSourceFilterModel) ([]ProtocolsCIFSShareGetDataModelONTAP, error) {
	api := "protocols/cifs/shares"
	query := r.NewQuery()
	query.Fields([]string{"name", "svm.name", "unix_symlink", "dir_umask", "file_umask", "acls", "home_directory", "force_group_for_create", "no_strict_security", "oplocks", "volume", "change_notify", "path", "encryption", "vscan_profile", "offline_files", "comment", "show_snapshot", "continuously_available", "namespace_caching"})
	if filter != nil {
		var filterMap map[string]interface{}
		if err := mapstructure.Decode(filter, &filterMap); err != nil {
			return nil, errorHandler.MakeAndReportError("error encoding protocols_cifs_shares filter info", fmt.Sprintf("error on filter %#v: %s", filter, err))
		}
		query.SetValues(filterMap)
	}
	statusCode, response, err := r.GetZeroOrMoreRecords(api, query, nil)
	if err == nil && response == nil {
		err = fmt.Errorf("no response for GET %s", api)
	}
	if err != nil {
		return nil, errorHandler.MakeAndReportError("error reading protocols_cifs_shares info", fmt.Sprintf("error on GET %s: %s, statusCode %d", api, err, statusCode))
	}

	var dataONTAP []ProtocolsCIFSShareGetDataModelONTAP
	for _, info := range response {
		var record ProtocolsCIFSShareGetDataModelONTAP
		if err := mapstructure.Decode(info, &record); err != nil {
			return nil, errorHandler.MakeAndReportError(fmt.Sprintf("failed to decode response from GET %s", api),
				fmt.Sprintf("error: %s, statusCode %d, info %#v", err, statusCode, info))
		}
		dataONTAP = append(dataONTAP, record)
	}
	tflog.Debug(errorHandler.Ctx, fmt.Sprintf("Read protocols_cifs_shares data source: %#v", dataONTAP))
	return dataONTAP, nil
}

// CreateProtocolsCIFSShare to create protocols_cifs_share
func CreateProtocolsCIFSShare(errorHandler *utils.ErrorHandler, r restclient.RestClient, body ProtocolsCIFSShareResourceBodyDataModelONTAP) (*ProtocolsCIFSShareGetDataModelONTAP, error) {
	api := "/protocols/cifs/shares"

	var bodyMap map[string]interface{}
	if err := mapstructure.Decode(body, &bodyMap); err != nil {
		return nil, errorHandler.MakeAndReportError("error encoding protocols_cifs_share body", fmt.Sprintf("error on encoding %s body: %s, body: %#v", api, err, body))
	}
	query := r.NewQuery()
	query.Add("return_records", "true")
	statusCode, response, err := r.CallCreateMethod(api, query, bodyMap)
	if err != nil {
		return nil, errorHandler.MakeAndReportError("error creating protocols_cifs_share", fmt.Sprintf("error on POST %s: %s, statusCode %d", api, err, statusCode))
	}

	var dataONTAP ProtocolsCIFSShareGetDataModelONTAP
	if err := mapstructure.Decode(response.Records[0], &dataONTAP); err != nil {
		return nil, errorHandler.MakeAndReportError("error decoding protocols_cifs_share info", fmt.Sprintf("error on decode storage/protocols_cifs_shares info: %s, statusCode %d, response %#v", err, statusCode, response))
	}
	tflog.Debug(errorHandler.Ctx, fmt.Sprintf("Create protocols_cifs_share source - udata: %#v", dataONTAP))
	return &dataONTAP, nil
}

// UpdateProtocolsCIFSShare to update protocols_cifs_share
func UpdateProtocolsCIFSShare(errorHandler *utils.ErrorHandler, r restclient.RestClient, body ProtocolsCIFSShareResourceBodyDataModelONTAP, name string, svmUUID string) error {
	api := "/protocols/cifs/shares/"
	var bodyMap map[string]interface{}
	if err := mapstructure.Decode(body, &bodyMap); err != nil {
		return errorHandler.MakeAndReportError("error encoding protocols_cifs_share body", fmt.Sprintf("error on encoding %s body: %s, body: %#v", api, err, body))
	}
	statusCode, _, err := r.CallUpdateMethod(api+"/"+svmUUID+"/"+name, nil, bodyMap)
	if err != nil {
		return errorHandler.MakeAndReportError("error updating protocols_cifs_share", fmt.Sprintf("error on POST %s: %s, statusCode %d", api, err, statusCode))
	}
	return nil
}

// DeleteProtocolsCIFSShare to delete protocols_cifs_share
func DeleteProtocolsCIFSShare(errorHandler *utils.ErrorHandler, r restclient.RestClient, name string, svmUUID string) error {
	api := "/protocols/cifs/shares"
	statusCode, _, err := r.CallDeleteMethod(api+"/"+svmUUID+"/"+name, nil, nil)
	if err != nil {
		return errorHandler.MakeAndReportError("error deleting protocols_cifs_share", fmt.Sprintf("error on DELETE %s: %s, statusCode %d", api, err, statusCode))
	}
	return nil
}
