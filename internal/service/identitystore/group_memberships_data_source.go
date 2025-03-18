// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	awstypes "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_identitystore_group_memberships", name="Group Memberships")
func newGroupMembershipsDataSource(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &groupMembershipsDataSource{}, nil
}

const (
	DSNameGroupMemberships = "Group Memberships Data Source"
)

type groupMembershipsDataSource struct {
	framework.DataSourceWithConfigure
}

func (d *groupMembershipsDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"group_memberships": schema.ListAttribute{
				CustomType: fwtypes.NewListNestedObjectTypeOf[groupMembershipModel](ctx),
				Computed:   true,
			},
			"group_id": schema.StringAttribute{
				Required: true,
			},
			"identity_store_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (d *groupMembershipsDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data groupMembershipsDataSourceModel
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	conn := d.Meta().IdentityStoreClient(ctx)

	input := identitystore.ListGroupMembershipsInput{}
	response.Diagnostics.Append(fwflex.Expand(ctx, data, &input)...)
	if response.Diagnostics.HasError() {
		return
	}

	output, err := findGroupMemberships(ctx, conn, input)
	if err != nil {
		response.Diagnostics.AddError(
			create.ProblemStandardMessage(names.IdentityStore, create.ErrActionReading, DSNameGroupMemberships, data.GroupID.String(), err),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(fwflex.Flatten(ctx, output, &data.GroupMemberships)...)
	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func findGroupMemberships(ctx context.Context, conn *identitystore.Client, input identitystore.ListGroupMembershipsInput) ([]awstypes.GroupMembership, error) {
	var output []awstypes.GroupMembership
	pages := identitystore.NewListGroupMembershipsPaginator(conn, &input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)
		if err != nil {
			return output, err
		}

		output = append(output, page.GroupMemberships...)
	}

	return output, nil
}

type groupMembershipsDataSourceModel struct {
	IdentityStoreID  types.String                                          `tfsdk:"identity_store_id"`
	GroupID          types.String                                          `tfsdk:"group_id"`
	GroupMemberships fwtypes.ListNestedObjectValueOf[groupMembershipModel] `tfsdk:"group_memberships"`
}

type groupMembershipModel struct {
	MemberID        types.String `tfsdk:"member_id"`
	MembershipID    types.String `tfsdk:"membership_id"`
	GroupID         types.String `tfsdk:"group_id"`
	IdentityStoreID types.String `tfsdk:"identity_store_id"`
}
