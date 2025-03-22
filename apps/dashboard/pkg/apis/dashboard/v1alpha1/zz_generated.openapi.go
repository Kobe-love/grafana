//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

package v1alpha1

import (
	v0alpha1 "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1"
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationActions":         schema_pkg_apis_dashboard_v1alpha1_AnnotationActions(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationPermission":      schema_pkg_apis_dashboard_v1alpha1_AnnotationPermission(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.Dashboard":                 schema_pkg_apis_dashboard_v1alpha1_Dashboard(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardAccess":           schema_pkg_apis_dashboard_v1alpha1_DashboardAccess(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardConversionStatus": schema_pkg_apis_dashboard_v1alpha1_DashboardConversionStatus(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardJSONCodec":        schema_pkg_apis_dashboard_v1alpha1_DashboardJSONCodec(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardList":             schema_pkg_apis_dashboard_v1alpha1_DashboardList(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardMetadata":         schema_pkg_apis_dashboard_v1alpha1_DashboardMetadata(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardStatus":           schema_pkg_apis_dashboard_v1alpha1_DashboardStatus(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardVersionInfo":      schema_pkg_apis_dashboard_v1alpha1_DashboardVersionInfo(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardVersionList":      schema_pkg_apis_dashboard_v1alpha1_DashboardVersionList(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardWithAccessInfo":   schema_pkg_apis_dashboard_v1alpha1_DashboardWithAccessInfo(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanel":              schema_pkg_apis_dashboard_v1alpha1_LibraryPanel(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelList":          schema_pkg_apis_dashboard_v1alpha1_LibraryPanelList(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelSpec":          schema_pkg_apis_dashboard_v1alpha1_LibraryPanelSpec(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelStatus":        schema_pkg_apis_dashboard_v1alpha1_LibraryPanelStatus(ref),
		"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.VersionsQueryOptions":      schema_pkg_apis_dashboard_v1alpha1_VersionsQueryOptions(ref),
		"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured":                   v0alpha1.Unstructured{}.OpenAPIDefinition(),
	}
}

func schema_pkg_apis_dashboard_v1alpha1_AnnotationActions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"canAdd": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"canAdd", "canEdit", "canDelete"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_AnnotationPermission(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"dashboard": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationActions"),
						},
					},
					"organization": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationActions"),
						},
					},
				},
				Required: []string{"dashboard", "organization"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationActions"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_Dashboard(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Spec is the spec of the Dashboard",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardStatus", "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardAccess(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Information about how the requesting user can use a given dashboard",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"slug": {
						SchemaProps: spec.SchemaProps{
							Description: "Metadata fields",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"canSave": {
						SchemaProps: spec.SchemaProps{
							Description: "The permissions part",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canAdmin": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canStar": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"annotationsPermissions": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationPermission"),
						},
					},
				},
				Required: []string{"canSave", "canEdit", "canAdmin", "canStar", "canDelete", "annotationsPermissions"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.AnnotationPermission"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardConversionStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConversionStatus is the status of the conversion of the dashboard.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"failed": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether from another version has failed. If true, means that the dashboard is not valid, and the caller should instead fetch the stored version.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"storedVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The version which was stored when the dashboard was created / updated. Fetching this version should always succeed.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"error": {
						SchemaProps: spec.SchemaProps{
							Description: "The error message from the conversion. Empty if the conversion has not failed.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"failed", "storedVersion", "error"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardJSONCodec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DashboardJSONCodec is an implementation of resource.Codec for kubernetes JSON encoding",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.Dashboard"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.Dashboard", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardMetadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "metadata contains embedded CommonMetadata and can be extended with custom string fields without external reference as using the CommonMetadata reference breaks thema codegen.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"updateTimestamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"createdBy": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"uid": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"creationTimestamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"deletionTimestamp": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"finalizers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"resourceVersion": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"generation": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int64",
						},
					},
					"updatedBy": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"updateTimestamp", "createdBy", "uid", "creationTimestamp", "finalizers", "resourceVersion", "generation", "updatedBy", "labels"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"conversion": {
						SchemaProps: spec.SchemaProps{
							Description: "Optional conversion status.",
							Ref:         ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardConversionStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardConversionStatus"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardVersionInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "The internal ID for this version (will be replaced with resourceVersion)",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"parentVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "If the dashboard came from a previous version, it is set here",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"created": {
						SchemaProps: spec.SchemaProps{
							Description: "The creation timestamp for this version",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"createdBy": {
						SchemaProps: spec.SchemaProps{
							Description: "The user who created this version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message passed while saving the version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"version", "created"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardVersionList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardVersionInfo"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardVersionInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_DashboardWithAccessInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "This is like the legacy DTO where access and metadata are all returned in a single call",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Spec is the spec of the Dashboard",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardStatus"),
						},
					},
					"access": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardAccess"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status", "access"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardAccess", "github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.DashboardStatus", "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_LibraryPanel(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Panel properties",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status will show errors",
							Ref:         ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelSpec", "github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanelStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_LibraryPanelList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanel"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/dashboard/pkg/apis/dashboard/v1alpha1.LibraryPanel", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_LibraryPanelSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "The panel type",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pluginVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The panel type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"title": {
						SchemaProps: spec.SchemaProps{
							Description: "The panel title",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Library panel description",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"options": {
						SchemaProps: spec.SchemaProps{
							Description: "The options schema depends on the panel type",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"fieldConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "The fieldConfig schema depends on the panel type",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"datasource": {
						SchemaProps: spec.SchemaProps{
							Description: "The default datasource type",
							Ref:         ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataSourceRef"),
						},
					},
					"targets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The datasource queries",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery"),
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "options", "fieldConfig"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery", "github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataSourceRef", "github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_LibraryPanelStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"warnings": {
						SchemaProps: spec.SchemaProps{
							Description: "Translation warnings (mostly things that were in SQL columns but not found in the saved body)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"missing": {
						SchemaProps: spec.SchemaProps{
							Description: "The properties previously stored in SQL that are not included in this model",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_dashboard_v1alpha1_VersionsQueryOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path is the URL path",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}
