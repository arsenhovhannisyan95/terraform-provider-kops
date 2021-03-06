package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCalicoNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_request":                        ComputedQuantity(),
			"cross_subnet":                       ComputedBool(),
			"log_severity_screen":                ComputedString(),
			"mtu":                                ComputedInt(),
			"prometheus_metrics_enabled":         ComputedBool(),
			"prometheus_metrics_port":            ComputedInt(),
			"prometheus_go_metrics_enabled":      ComputedBool(),
			"prometheus_process_metrics_enabled": ComputedBool(),
			"major_version":                      ComputedString(),
			"iptables_backend":                   ComputedString(),
			"ip_ip_mode":                         ComputedString(),
			"typha_prometheus_metrics_enabled":   ComputedBool(),
			"typha_prometheus_metrics_port":      ComputedInt(),
			"typha_replicas":                     ComputedInt(),
			"wireguard_enabled":                  ComputedBool(),
		},
	}
}

func ExpandDataSourceCalicoNetworkingSpec(in map[string]interface{}) kops.CalicoNetworkingSpec {
	if in == nil {
		panic("expand CalicoNetworkingSpec failure, in is nil")
	}
	return kops.CalicoNetworkingSpec{
		CPURequest: func(in interface{}) *resource.Quantity {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *resource.Quantity {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resource.Quantity) *resource.Quantity {
					return &in
				}(ExpandQuantity(in))
			}(in)
		}(in["cpu_request"]),
		CrossSubnet: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["cross_subnet"]),
		LogSeverityScreen: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["log_severity_screen"]),
		MTU: func(in interface{}) *int32 {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["mtu"]),
		PrometheusMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_metrics_enabled"]),
		PrometheusMetricsPort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["prometheus_metrics_port"]),
		PrometheusGoMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_go_metrics_enabled"]),
		PrometheusProcessMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["prometheus_process_metrics_enabled"]),
		MajorVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["major_version"]),
		IptablesBackend: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["iptables_backend"]),
		IPIPMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ip_ip_mode"]),
		TyphaPrometheusMetricsEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["typha_prometheus_metrics_enabled"]),
		TyphaPrometheusMetricsPort: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["typha_prometheus_metrics_port"]),
		TyphaReplicas: func(in interface{}) int32 {
			return int32(ExpandInt(in))
		}(in["typha_replicas"]),
		WireguardEnabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["wireguard_enabled"]),
	}
}

func FlattenDataSourceCalicoNetworkingSpecInto(in kops.CalicoNetworkingSpec, out map[string]interface{}) {
	out["cpu_request"] = func(in *resource.Quantity) interface{} {
		return func(in *resource.Quantity) interface{} {
			if in == nil {
				return nil
			}
			return func(in resource.Quantity) interface{} {
				return FlattenQuantity(in)
			}(*in)
		}(in)
	}(in.CPURequest)
	out["cross_subnet"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.CrossSubnet)
	out["log_severity_screen"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LogSeverityScreen)
	out["mtu"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MTU)
	out["prometheus_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusMetricsEnabled)
	out["prometheus_metrics_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.PrometheusMetricsPort)
	out["prometheus_go_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusGoMetricsEnabled)
	out["prometheus_process_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PrometheusProcessMetricsEnabled)
	out["major_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MajorVersion)
	out["iptables_backend"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IptablesBackend)
	out["ip_ip_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.IPIPMode)
	out["typha_prometheus_metrics_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.TyphaPrometheusMetricsEnabled)
	out["typha_prometheus_metrics_port"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TyphaPrometheusMetricsPort)
	out["typha_replicas"] = func(in int32) interface{} {
		return FlattenInt(int(in))
	}(in.TyphaReplicas)
	out["wireguard_enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.WireguardEnabled)
}

func FlattenDataSourceCalicoNetworkingSpec(in kops.CalicoNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCalicoNetworkingSpecInto(in, out)
	return out
}
