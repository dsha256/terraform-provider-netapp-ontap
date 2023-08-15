package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/interfaces"
	"github.com/netapp/terraform-provider-netapp-ontap/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &StorageVolumesDataSource{}

// NewStorageVolumesDataSource is a helper function to simplify the provider implementation.
func NewStorageVolumesDataSource() datasource.DataSource {
	return &StorageVolumesDataSource{
		config: resourceOrDataSourceConfig{
			name: "storage_volumes_data_source",
		},
	}
}

// StorageVolumesDataSource defines the data source implementation.
type StorageVolumesDataSource struct {
	config resourceOrDataSourceConfig
}

// StorageVolumesDataSourceModel describes the data source data model.
type StorageVolumesDataSourceModel struct {
	CxProfileName  types.String                        `tfsdk:"cx_profile_name"`
	StorageVolumes []StorageVolumeDataSourceModel      `tfsdk:"storage_volumes"`
	Filter         *StorageVolumeDataSourceFilterModel `tfsdk:"filter"`
}

// StorageVolumeDataSourceFilterModel describes the data source data model for queries.
type StorageVolumeDataSourceFilterModel struct {
	Name    types.String `tfsdk:"name"`
	SVMName types.String `tfsdk:"svm_name"`
}

// Metadata returns the data source type name.
func (d *StorageVolumesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.config.name
}

// Schema defines the schema for the data source.
func (d *StorageVolumesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "StorageVolumes data source",

		Attributes: map[string]schema.Attribute{
			"cx_profile_name": schema.StringAttribute{
				MarkdownDescription: "Connection profile name",
				Required:            true,
			},
			"filter": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: "StorageVolume name",
						Optional:            true,
					},
					"svm_name": schema.StringAttribute{
						MarkdownDescription: "StorageVolume svm name",
						Optional:            true,
					},
				},
				Optional: true,
			},
			"storage_volumes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cx_profile_name": schema.StringAttribute{
							MarkdownDescription: "Connection profile name",
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "StorageVolume name",
							Required:            true,
						},
						"svm_name": schema.StringAttribute{
							MarkdownDescription: "Name of the svm to use",
							Required:            true,
						},
						"size": schema.Int64Attribute{
							MarkdownDescription: "The size of the volume",
							Computed:            true,
						},
						"size_unit": schema.StringAttribute{
							MarkdownDescription: "The unit used to interpret the size parameter",
							Computed:            true,
						},
						"is_online": schema.BoolAttribute{
							MarkdownDescription: "Whether the specified volume is online, or not",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The volume type, either read-write (RW) or data-protection (DP)",
							Computed:            true,
						},
						"export_policy": schema.StringAttribute{
							MarkdownDescription: "The name of the export policy",
							Computed:            true,
						},
						"junction_path": schema.StringAttribute{
							MarkdownDescription: "Junction path of the volume",
							Computed:            true,
						},
						"space_guarantee": schema.StringAttribute{
							MarkdownDescription: "Space guarantee style for the volume",
							Computed:            true,
						},
						"percent_snapshot_space": schema.Int64Attribute{
							MarkdownDescription: "Amount of space reserved for snapshot copies of the volume",
							Computed:            true,
						},
						"security_style": schema.StringAttribute{
							MarkdownDescription: "The security style associated to the volume",
							Computed:            true,
						},
						"encrypt": schema.BoolAttribute{
							MarkdownDescription: "Whether or not to enable Volume Encryption",
							Computed:            true,
						},
						"efficiency_policy": schema.StringAttribute{
							MarkdownDescription: "Allows a storage efficiency policy to be set on volume creation",
							Computed:            true,
						},
						"unix_permissions": schema.StringAttribute{
							MarkdownDescription: "Unix permission bits in octal or symbolic format. For example, 0 is equivalent to ------------, 777 is equivalent to ---rwxrwxrwx,both formats are accepted",
							Computed:            true,
						},
						"group_id": schema.Int64Attribute{
							MarkdownDescription: "The UNIX group ID for the volume",
							Computed:            true,
						},
						"user_id": schema.Int64Attribute{
							MarkdownDescription: "The UNIX user ID for the volume",
							Computed:            true,
						},
						"snapshot_policy": schema.StringAttribute{
							MarkdownDescription: "The name of the snapshot policy",
							Computed:            true,
						},
						"language": schema.StringAttribute{
							MarkdownDescription: "Language to use for volume",
							Computed:            true,
						},
						"qos_policy_group": schema.StringAttribute{
							MarkdownDescription: "Specifies a QoS policy group to be set on volume",
							Computed:            true,
						},
						"qos_adaptive_policy_group": schema.StringAttribute{
							MarkdownDescription: "Specifies a QoS adaptive policy group to be set on volume",
							Computed:            true,
						},
						"tiering_policy": schema.StringAttribute{
							MarkdownDescription: "The tiering policy that is to be associated with the volume",
							Computed:            true,
						},
						"comment": schema.StringAttribute{
							MarkdownDescription: "Sets a comment associated with the volume",
							Computed:            true,
						},
						"compression": schema.StringAttribute{
							MarkdownDescription: "Whether to enable compression for the volume (HDD and Flash Pool aggregates, AFF platforms)",
							Computed:            true,
						},
						"tiering_minimum_cooling_days": schema.Int64Attribute{
							MarkdownDescription: "Determines how many days must pass before inactive data in a volume using the Auto or Snapshot-Only policy is considered cold and eligible for tiering",
							Computed:            true,
						},
						"logical_space_enforcement": schema.BoolAttribute{
							MarkdownDescription: "Whether to perform logical space accounting on the volume",
							Computed:            true,
						},
						"logical_space_reporting": schema.BoolAttribute{
							MarkdownDescription: "Whether to report space logically",
							Computed:            true,
						},
						"aggregates": schema.ListAttribute{
							ElementType:         types.StringType,
							Computed:            true,
							MarkdownDescription: "List of aggregates in which to create the volume",
						},
						"snaplock_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The SnapLock type of the volume",
						},
						"analytics": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Set file system analytics state of the volume",
						},
						"uuid": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Volume identifier",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "",
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *StorageVolumesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	config, ok := req.ProviderData.(Config)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected Config, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
	}
	d.config.providerConfig = config
}

// Read refreshes the Terraform state with the latest data.
func (d *StorageVolumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data StorageVolumesDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	errorHandler := utils.NewErrorHandler(ctx, &resp.Diagnostics)
	// we need to defer setting the client until we can read the connection profile name
	client, err := getRestClient(errorHandler, d.config, data.CxProfileName)
	if err != nil {
		// error reporting done inside NewClient
		return
	}

	var filter *interfaces.StorageVolumeDataSourceFilterModel = nil
	if data.Filter != nil {
		filter = &interfaces.StorageVolumeDataSourceFilterModel{
			Name:    data.Filter.Name.ValueString(),
			SVMName: data.Filter.SVMName.ValueString(),
		}
	}
	restInfo, err := interfaces.GetStorageVolumes(errorHandler, *client, filter)
	if err != nil {
		// error reporting done inside GetStorageVolumes
		return
	}

	data.StorageVolumes = make([]StorageVolumeDataSourceModel, len(restInfo))
	for index, record := range restInfo {
		var aggregates []types.String
		for _, e := range record.Aggregates {
			aggregates = append(aggregates, types.StringValue(e.Name))
		}

		size, sizeUnit := interfaces.ByteFormat(int64(record.Space.Size))

		data.StorageVolumes[index] = StorageVolumeDataSourceModel{
			CxProfileName:             types.String(data.CxProfileName),
			Name:                      types.StringValue(record.Name),
			UUID:                      types.StringValue(record.UUID),
			SVMName:                   types.StringValue(record.SVM.Name),
			Aggregates:                aggregates,
			Size:                      types.Int64Value(size),
			SizeUnit:                  types.StringValue(sizeUnit),
			IsOnline:                  types.BoolValue(interfaces.OnlineToBool(record.State)),
			Type:                      types.StringValue(record.Type),
			ExportPolicy:              types.StringValue(record.NAS.ExportPolicy.Name),
			JunctionPath:              types.StringValue(record.NAS.JunctionPath),
			SpaceGuarantee:            types.StringValue(record.SpaceGuarantee.Type),
			PercentSnapshotSpace:      types.Int64Value(int64(record.Space.Snapshot.ReservePercent)),
			SecurityStyle:             types.StringValue(record.NAS.SecurityStyle),
			Encrypt:                   types.BoolValue(record.Encryption.Enabled),
			EfficiencyPolicy:          types.StringValue(record.Efficiency.Policy.Name),
			UnixPermissions:           types.StringValue(strconv.Itoa(record.NAS.UnixPermissions)),
			GroupID:                   types.Int64Value(int64(record.NAS.GroupID)),
			UserID:                    types.Int64Value(int64(record.NAS.UserID)),
			SnapshotPolicy:            types.StringValue(record.SnapshotPolicy.Name),
			Language:                  types.StringValue(record.Language),
			QosPolicyGroup:            types.StringValue(record.QOS.Policy.Name),
			QosAdaptivePolicyGroup:    types.StringValue(record.QOS.Policy.Name),
			TieringPolicy:             types.StringValue(record.TieringPolicy.Policy),
			Comment:                   types.StringValue(record.Comment),
			Compression:               types.StringValue(record.Efficiency.Compression),
			TieringMinimumCoolingDays: types.Int64Value(int64(record.TieringPolicy.MinCoolingDays)),
			LogicalSpaceEnforcement:   types.BoolValue(record.Space.LogicalSpace.Enforcement),
			LogicalSpaceReporting:     types.BoolValue(record.Space.LogicalSpace.Reporting),
			SnaplockType:              types.StringValue(record.Snaplock.Type),
			Analytics:                 types.StringValue(record.Analytics.State),
		}
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Debug(ctx, fmt.Sprintf("read a data source: %#v", data))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
