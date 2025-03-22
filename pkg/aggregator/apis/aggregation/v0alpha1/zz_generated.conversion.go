//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by conversion-gen. DO NOT EDIT.

package v0alpha1

import (
	unsafe "unsafe"

	aggregation "github.com/grafana/grafana/pkg/aggregator/apis/aggregation"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DataPlaneService)(nil), (*aggregation.DataPlaneService)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_DataPlaneService_To_aggregation_DataPlaneService(a.(*DataPlaneService), b.(*aggregation.DataPlaneService), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.DataPlaneService)(nil), (*DataPlaneService)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_DataPlaneService_To_v0alpha1_DataPlaneService(a.(*aggregation.DataPlaneService), b.(*DataPlaneService), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DataPlaneServiceCondition)(nil), (*aggregation.DataPlaneServiceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_DataPlaneServiceCondition_To_aggregation_DataPlaneServiceCondition(a.(*DataPlaneServiceCondition), b.(*aggregation.DataPlaneServiceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.DataPlaneServiceCondition)(nil), (*DataPlaneServiceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_DataPlaneServiceCondition_To_v0alpha1_DataPlaneServiceCondition(a.(*aggregation.DataPlaneServiceCondition), b.(*DataPlaneServiceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DataPlaneServiceList)(nil), (*aggregation.DataPlaneServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_DataPlaneServiceList_To_aggregation_DataPlaneServiceList(a.(*DataPlaneServiceList), b.(*aggregation.DataPlaneServiceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.DataPlaneServiceList)(nil), (*DataPlaneServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_DataPlaneServiceList_To_v0alpha1_DataPlaneServiceList(a.(*aggregation.DataPlaneServiceList), b.(*DataPlaneServiceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DataPlaneServiceSpec)(nil), (*aggregation.DataPlaneServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_DataPlaneServiceSpec_To_aggregation_DataPlaneServiceSpec(a.(*DataPlaneServiceSpec), b.(*aggregation.DataPlaneServiceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.DataPlaneServiceSpec)(nil), (*DataPlaneServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_DataPlaneServiceSpec_To_v0alpha1_DataPlaneServiceSpec(a.(*aggregation.DataPlaneServiceSpec), b.(*DataPlaneServiceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DataPlaneServiceStatus)(nil), (*aggregation.DataPlaneServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_DataPlaneServiceStatus_To_aggregation_DataPlaneServiceStatus(a.(*DataPlaneServiceStatus), b.(*aggregation.DataPlaneServiceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.DataPlaneServiceStatus)(nil), (*DataPlaneServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_DataPlaneServiceStatus_To_v0alpha1_DataPlaneServiceStatus(a.(*aggregation.DataPlaneServiceStatus), b.(*DataPlaneServiceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*QueryDataResponse)(nil), (*aggregation.QueryDataResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_QueryDataResponse_To_aggregation_QueryDataResponse(a.(*QueryDataResponse), b.(*aggregation.QueryDataResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.QueryDataResponse)(nil), (*QueryDataResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_QueryDataResponse_To_v0alpha1_QueryDataResponse(a.(*aggregation.QueryDataResponse), b.(*QueryDataResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Service)(nil), (*aggregation.Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v0alpha1_Service_To_aggregation_Service(a.(*Service), b.(*aggregation.Service), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*aggregation.Service)(nil), (*Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_aggregation_Service_To_v0alpha1_Service(a.(*aggregation.Service), b.(*Service), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v0alpha1_DataPlaneService_To_aggregation_DataPlaneService(in *DataPlaneService, out *aggregation.DataPlaneService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v0alpha1_DataPlaneServiceSpec_To_aggregation_DataPlaneServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v0alpha1_DataPlaneServiceStatus_To_aggregation_DataPlaneServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v0alpha1_DataPlaneService_To_aggregation_DataPlaneService is an autogenerated conversion function.
func Convert_v0alpha1_DataPlaneService_To_aggregation_DataPlaneService(in *DataPlaneService, out *aggregation.DataPlaneService, s conversion.Scope) error {
	return autoConvert_v0alpha1_DataPlaneService_To_aggregation_DataPlaneService(in, out, s)
}

func autoConvert_aggregation_DataPlaneService_To_v0alpha1_DataPlaneService(in *aggregation.DataPlaneService, out *DataPlaneService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_aggregation_DataPlaneServiceSpec_To_v0alpha1_DataPlaneServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_aggregation_DataPlaneServiceStatus_To_v0alpha1_DataPlaneServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_aggregation_DataPlaneService_To_v0alpha1_DataPlaneService is an autogenerated conversion function.
func Convert_aggregation_DataPlaneService_To_v0alpha1_DataPlaneService(in *aggregation.DataPlaneService, out *DataPlaneService, s conversion.Scope) error {
	return autoConvert_aggregation_DataPlaneService_To_v0alpha1_DataPlaneService(in, out, s)
}

func autoConvert_v0alpha1_DataPlaneServiceCondition_To_aggregation_DataPlaneServiceCondition(in *DataPlaneServiceCondition, out *aggregation.DataPlaneServiceCondition, s conversion.Scope) error {
	out.Type = aggregation.DataPlaneServiceConditionType(in.Type)
	out.Status = aggregation.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v0alpha1_DataPlaneServiceCondition_To_aggregation_DataPlaneServiceCondition is an autogenerated conversion function.
func Convert_v0alpha1_DataPlaneServiceCondition_To_aggregation_DataPlaneServiceCondition(in *DataPlaneServiceCondition, out *aggregation.DataPlaneServiceCondition, s conversion.Scope) error {
	return autoConvert_v0alpha1_DataPlaneServiceCondition_To_aggregation_DataPlaneServiceCondition(in, out, s)
}

func autoConvert_aggregation_DataPlaneServiceCondition_To_v0alpha1_DataPlaneServiceCondition(in *aggregation.DataPlaneServiceCondition, out *DataPlaneServiceCondition, s conversion.Scope) error {
	out.Type = DataPlaneServiceConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_aggregation_DataPlaneServiceCondition_To_v0alpha1_DataPlaneServiceCondition is an autogenerated conversion function.
func Convert_aggregation_DataPlaneServiceCondition_To_v0alpha1_DataPlaneServiceCondition(in *aggregation.DataPlaneServiceCondition, out *DataPlaneServiceCondition, s conversion.Scope) error {
	return autoConvert_aggregation_DataPlaneServiceCondition_To_v0alpha1_DataPlaneServiceCondition(in, out, s)
}

func autoConvert_v0alpha1_DataPlaneServiceList_To_aggregation_DataPlaneServiceList(in *DataPlaneServiceList, out *aggregation.DataPlaneServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]aggregation.DataPlaneService)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v0alpha1_DataPlaneServiceList_To_aggregation_DataPlaneServiceList is an autogenerated conversion function.
func Convert_v0alpha1_DataPlaneServiceList_To_aggregation_DataPlaneServiceList(in *DataPlaneServiceList, out *aggregation.DataPlaneServiceList, s conversion.Scope) error {
	return autoConvert_v0alpha1_DataPlaneServiceList_To_aggregation_DataPlaneServiceList(in, out, s)
}

func autoConvert_aggregation_DataPlaneServiceList_To_v0alpha1_DataPlaneServiceList(in *aggregation.DataPlaneServiceList, out *DataPlaneServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]DataPlaneService)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_aggregation_DataPlaneServiceList_To_v0alpha1_DataPlaneServiceList is an autogenerated conversion function.
func Convert_aggregation_DataPlaneServiceList_To_v0alpha1_DataPlaneServiceList(in *aggregation.DataPlaneServiceList, out *DataPlaneServiceList, s conversion.Scope) error {
	return autoConvert_aggregation_DataPlaneServiceList_To_v0alpha1_DataPlaneServiceList(in, out, s)
}

func autoConvert_v0alpha1_DataPlaneServiceSpec_To_aggregation_DataPlaneServiceSpec(in *DataPlaneServiceSpec, out *aggregation.DataPlaneServiceSpec, s conversion.Scope) error {
	out.PluginID = in.PluginID
	out.PluginType = aggregation.PluginType(in.PluginType)
	out.Group = in.Group
	out.Version = in.Version
	out.Services = *(*[]aggregation.Service)(unsafe.Pointer(&in.Services))
	return nil
}

// Convert_v0alpha1_DataPlaneServiceSpec_To_aggregation_DataPlaneServiceSpec is an autogenerated conversion function.
func Convert_v0alpha1_DataPlaneServiceSpec_To_aggregation_DataPlaneServiceSpec(in *DataPlaneServiceSpec, out *aggregation.DataPlaneServiceSpec, s conversion.Scope) error {
	return autoConvert_v0alpha1_DataPlaneServiceSpec_To_aggregation_DataPlaneServiceSpec(in, out, s)
}

func autoConvert_aggregation_DataPlaneServiceSpec_To_v0alpha1_DataPlaneServiceSpec(in *aggregation.DataPlaneServiceSpec, out *DataPlaneServiceSpec, s conversion.Scope) error {
	out.PluginID = in.PluginID
	out.PluginType = PluginType(in.PluginType)
	out.Group = in.Group
	out.Version = in.Version
	out.Services = *(*[]Service)(unsafe.Pointer(&in.Services))
	return nil
}

// Convert_aggregation_DataPlaneServiceSpec_To_v0alpha1_DataPlaneServiceSpec is an autogenerated conversion function.
func Convert_aggregation_DataPlaneServiceSpec_To_v0alpha1_DataPlaneServiceSpec(in *aggregation.DataPlaneServiceSpec, out *DataPlaneServiceSpec, s conversion.Scope) error {
	return autoConvert_aggregation_DataPlaneServiceSpec_To_v0alpha1_DataPlaneServiceSpec(in, out, s)
}

func autoConvert_v0alpha1_DataPlaneServiceStatus_To_aggregation_DataPlaneServiceStatus(in *DataPlaneServiceStatus, out *aggregation.DataPlaneServiceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]aggregation.DataPlaneServiceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v0alpha1_DataPlaneServiceStatus_To_aggregation_DataPlaneServiceStatus is an autogenerated conversion function.
func Convert_v0alpha1_DataPlaneServiceStatus_To_aggregation_DataPlaneServiceStatus(in *DataPlaneServiceStatus, out *aggregation.DataPlaneServiceStatus, s conversion.Scope) error {
	return autoConvert_v0alpha1_DataPlaneServiceStatus_To_aggregation_DataPlaneServiceStatus(in, out, s)
}

func autoConvert_aggregation_DataPlaneServiceStatus_To_v0alpha1_DataPlaneServiceStatus(in *aggregation.DataPlaneServiceStatus, out *DataPlaneServiceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]DataPlaneServiceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_aggregation_DataPlaneServiceStatus_To_v0alpha1_DataPlaneServiceStatus is an autogenerated conversion function.
func Convert_aggregation_DataPlaneServiceStatus_To_v0alpha1_DataPlaneServiceStatus(in *aggregation.DataPlaneServiceStatus, out *DataPlaneServiceStatus, s conversion.Scope) error {
	return autoConvert_aggregation_DataPlaneServiceStatus_To_v0alpha1_DataPlaneServiceStatus(in, out, s)
}

func autoConvert_v0alpha1_QueryDataResponse_To_aggregation_QueryDataResponse(in *QueryDataResponse, out *aggregation.QueryDataResponse, s conversion.Scope) error {
	out.QueryDataResponse = in.QueryDataResponse
	return nil
}

// Convert_v0alpha1_QueryDataResponse_To_aggregation_QueryDataResponse is an autogenerated conversion function.
func Convert_v0alpha1_QueryDataResponse_To_aggregation_QueryDataResponse(in *QueryDataResponse, out *aggregation.QueryDataResponse, s conversion.Scope) error {
	return autoConvert_v0alpha1_QueryDataResponse_To_aggregation_QueryDataResponse(in, out, s)
}

func autoConvert_aggregation_QueryDataResponse_To_v0alpha1_QueryDataResponse(in *aggregation.QueryDataResponse, out *QueryDataResponse, s conversion.Scope) error {
	out.QueryDataResponse = in.QueryDataResponse
	return nil
}

// Convert_aggregation_QueryDataResponse_To_v0alpha1_QueryDataResponse is an autogenerated conversion function.
func Convert_aggregation_QueryDataResponse_To_v0alpha1_QueryDataResponse(in *aggregation.QueryDataResponse, out *QueryDataResponse, s conversion.Scope) error {
	return autoConvert_aggregation_QueryDataResponse_To_v0alpha1_QueryDataResponse(in, out, s)
}

func autoConvert_v0alpha1_Service_To_aggregation_Service(in *Service, out *aggregation.Service, s conversion.Scope) error {
	out.Type = aggregation.ServiceType(in.Type)
	out.Method = in.Method
	out.Path = in.Path
	return nil
}

// Convert_v0alpha1_Service_To_aggregation_Service is an autogenerated conversion function.
func Convert_v0alpha1_Service_To_aggregation_Service(in *Service, out *aggregation.Service, s conversion.Scope) error {
	return autoConvert_v0alpha1_Service_To_aggregation_Service(in, out, s)
}

func autoConvert_aggregation_Service_To_v0alpha1_Service(in *aggregation.Service, out *Service, s conversion.Scope) error {
	out.Type = ServiceType(in.Type)
	out.Method = in.Method
	out.Path = in.Path
	return nil
}

// Convert_aggregation_Service_To_v0alpha1_Service is an autogenerated conversion function.
func Convert_aggregation_Service_To_v0alpha1_Service(in *aggregation.Service, out *Service, s conversion.Scope) error {
	return autoConvert_aggregation_Service_To_v0alpha1_Service(in, out, s)
}
