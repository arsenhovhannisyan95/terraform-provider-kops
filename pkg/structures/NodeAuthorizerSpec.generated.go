package structures

import (
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandNodeAuthorizerSpec(in map[string]interface{}) kops.NodeAuthorizerSpec {
	if in == nil {
		panic("expand NodeAuthorizerSpec failure, in is nil")
	}
	return kops.NodeAuthorizerSpec{
		Authorizer: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["authorizer"]),
		Features: func(in interface{}) *[]string {
			return func(in interface{}) *[]string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in []string) *[]string {
					return &in
				}(func(in interface{}) []string {
					var out []string
					for _, in := range in.([]interface{}) {
						out = append(out, string(ExpandString(in)))
					}
					return out
				}(in))
			}(in)
		}(in["features"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		NodeURL: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["node_url"]),
		Port: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["port"]),
		Interval: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["interval"]),
		Timeout: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["timeout"]),
		TokenTTL: func(in interface{}) *v1.Duration {
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in v1.Duration) *v1.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["token_ttl"]),
	}
}

func FlattenNodeAuthorizerSpec(in kops.NodeAuthorizerSpec) map[string]interface{} {
	return map[string]interface{}{
		"authorizer": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Authorizer),
		"features": func(in *[]string) interface{} {
			return func(in *[]string) interface{} {
				if in == nil {
					return nil
				}
				return func(in []string) interface{} {
					return func(in []string) []interface{} {
						var out []interface{}
						for _, in := range in {
							out = append(out, FlattenString(string(in)))
						}
						return out
					}(in)
				}(*in)
			}(in)
		}(in.Features),
		"image": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.Image),
		"node_url": func(in string) interface{} {
			return FlattenString(string(in))
		}(in.NodeURL),
		"port": func(in int) interface{} {
			return FlattenInt(int(in))
		}(in.Port),
		"interval": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.Interval),
		"timeout": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.Timeout),
		"token_ttl": func(in *v1.Duration) interface{} {
			return func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
		}(in.TokenTTL),
	}
}