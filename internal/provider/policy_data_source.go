package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PolicyDataSource{}

func NewPolicyDataSource() datasource.DataSource {
	return &PolicyDataSource{}
}

// PolicyDataSource defines the data source implementation.
type PolicyDataSource struct{}

type Hash struct {
	Algorithm types.String `tfsdk:"algorithm"`
	Value     types.String `tfsdk:"value"`
}

type Directive struct {
	Name types.String `tfsdk:"name"`

	Keywords []types.String `tfsdk:"keywords"`
	Hosts    []types.String `tfsdk:"hosts"`
	Schemes  []types.String `tfsdk:"schemes"`
	Nonces   []types.String `tfsdk:"nonces"`
	Values   []types.String `tfsdk:"values"`

	Hashes []Hash `tfsdk:"hash"`
}

// PolicyDataSourceModel describes the data source data model.
type PolicyDataSourceModel struct {
	Directives []Directive `tfsdk:"directive"`

	Value types.String `tfsdk:"value"`
}

func (d *PolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy"
}

func (d *PolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Generate a Content-Security-Policy header value.",

		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "The generated Content-Security-Policy header value.",
				Computed:    true,
			},
		},

		Blocks: map[string]schema.Block{
			"directive": schema.ListNestedBlock{
				Description: "Directives to include in the policy.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "The name of the directive.",
							Required:    true,
						},
						"keywords": schema.ListAttribute{
							Description: "Keywords to include as values for the directive.",
							ElementType: types.StringType,
							Optional:    true,
						},
						"hosts": schema.ListAttribute{
							Description: "Hosts to include as values for the directive.",
							ElementType: types.StringType,
							Optional:    true,
						},
						"schemes": schema.ListAttribute{
							Description: "Schemes to include as values for the directive.",
							ElementType: types.StringType,
							Optional:    true,
						},
						// TODO: Figure out whether or not to include this
						"nonces": schema.ListAttribute{
							Description: "Nonces to include as values for the directive.",
							ElementType: types.StringType,
							Optional:    true,
						},
						"values": schema.ListAttribute{
							Description: "Any extra values to include in the directive.",
							ElementType: types.StringType,
							Optional:    true,
						},
					},
					Blocks: map[string]schema.Block{
						"hash": schema.ListNestedBlock{
							Description: "Hashes to include as values for the directive.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"algorithm": schema.StringAttribute{
										Description: "The algorithm used to generate the hash.",
										Required:    true,
									},
									"value": schema.StringAttribute{
										Description: "The base64-encoded hash value.",
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *PolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *PolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PolicyDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Value = types.StringValue(data.GeneratePolicy())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
