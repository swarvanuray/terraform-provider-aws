// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	opensearchserverless_sdkv2 "github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceAccessPolicy,
			Name:    "Access Policy",
		},
		{
			Factory: newDataSourceCollection,
			Name:    "Collection",
		},
		{
			Factory: newDataSourceLifecyclePolicy,
			Name:    "Lifecycle Policy",
		},
		{
			Factory: newDataSourceSecurityConfig,
			Name:    "Security Config",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceAccessPolicy,
		},
		{
			Factory: newResourceCollection,
			Name:    "Collection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory: newResourceLifecyclePolicy,
			Name:    "Lifecycle Policy",
		},
		{
			Factory: newResourceSecurityConfig,
			Name:    "Security Config",
		},
		{
			Factory: newResourceSecurityPolicy,
		},
		{
			Factory: newResourceVPCEndpoint,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceSecurityPolicy,
			TypeName: "aws_opensearchserverless_security_policy",
		},
		{
			Factory:  DataSourceVPCEndpoint,
			TypeName: "aws_opensearchserverless_vpc_endpoint",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpenSearchServerless
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*opensearchserverless_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return opensearchserverless_sdkv2.NewFromConfig(cfg,
		opensearchserverless_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
