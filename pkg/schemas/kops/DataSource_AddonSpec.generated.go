package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceAddonSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manifest": ComputedString(),
		},
	}
}

func ExpandDataSourceAddonSpec(in map[string]interface{}) kops.AddonSpec {
	if in == nil {
		panic("expand AddonSpec failure, in is nil")
	}
	return kops.AddonSpec{
		Manifest: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["manifest"]),
	}
}

func FlattenDataSourceAddonSpecInto(in kops.AddonSpec, out map[string]interface{}) {
	out["manifest"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Manifest)
}

func FlattenDataSourceAddonSpec(in kops.AddonSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceAddonSpecInto(in, out)
	return out
}