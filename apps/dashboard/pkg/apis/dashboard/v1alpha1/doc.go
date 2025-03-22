// +k8s:openapi-gen=true
// +k8s:defaulter-gen=TypeMeta
// +k8s:conversion-gen=github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard
// +groupName=dashboard.grafana.app

// NOTE (@radiohead): we do not use package-wide deepcopy generation
// because grafana-app-sdk already provides deepcopy functions.
// Kinds which are not generated by the SDK are explicitly opted in to deepcopy generation.

package v1alpha1 // import "github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1"
