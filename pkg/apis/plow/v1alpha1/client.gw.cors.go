// Code generated by protoc-gen-grpc-gateway-cors
// source: client.proto
// DO NOT EDIT!

/*
Package v1alpha1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1alpha1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportClientsCorsPatterns returns an array of grpc gatway mux patterns for
// Clients service to enable CORS.
func ExportClientsCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Clients_List_0,
		pattern_Clients_Describe_0,
		pattern_Clients_Create_0,
		pattern_Clients_Delete_0,
		pattern_Clients_Update_0,
		pattern_Clients_Copy_0,
		pattern_Clients_EditConfigMap_0,
		pattern_Clients_EditSecret_0,
		pattern_Clients_RegisterPersistentVolume_0,
		pattern_Clients_UnregisterPersistentVolume_0,
		pattern_Clients_RegisterPersistentVolumeClaim_0,
		pattern_Clients_UnregisterPersistentVolumeClaim_0,
		pattern_Clients_ReverseIndex_0,
	}
}

// ExportDisksCorsPatterns returns an array of grpc gatway mux patterns for
// Disks service to enable CORS.
func ExportDisksCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Disks_List_0,
		pattern_Disks_Describe_0,
		pattern_Disks_Create_0,
		pattern_Disks_Delete_0,
	}
}
