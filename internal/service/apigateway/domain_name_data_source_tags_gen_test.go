// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package apigateway_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccAPIGatewayDomainNameDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/DomainName/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtKey1: config.StringVariable(acctest.CtValue1),
					}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_NullMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/DomainName/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:          config.StringVariable(rName),
					acctest.CtResourceTags:   nil,
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_EmptyMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/DomainName/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:          config.StringVariable(rName),
					acctest.CtResourceTags:   config.MapVariable(map[string]config.Variable{}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}
