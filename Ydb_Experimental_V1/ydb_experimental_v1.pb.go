// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ydb_experimental_v1.proto

package Ydb_Experimental_V1

import (
	Ydb_Experimental "github.com/YandexDatabase/ydb-go-genproto/protos/Ydb_Experimental"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ydb_experimental_v1_proto protoreflect.FileDescriptor

var file_ydb_experimental_v1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x79, 0x64, 0x62, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x59, 0x64, 0x62,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31,
	0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xcf, 0x02, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x71, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x6c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x5f, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x79, 0x64, 0x62, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x79,
	0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59,
	0x64, 0x62, 0x5f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_experimental_v1_proto_goTypes = []interface{}{
	(*Ydb_Experimental.UploadRowsRequest)(nil),          // 0: Ydb.Experimental.UploadRowsRequest
	(*Ydb_Experimental.ExecuteStreamQueryRequest)(nil),  // 1: Ydb.Experimental.ExecuteStreamQueryRequest
	(*Ydb_Experimental.GetDiskSpaceUsageRequest)(nil),   // 2: Ydb.Experimental.GetDiskSpaceUsageRequest
	(*Ydb_Experimental.UploadRowsResponse)(nil),         // 3: Ydb.Experimental.UploadRowsResponse
	(*Ydb_Experimental.ExecuteStreamQueryResponse)(nil), // 4: Ydb.Experimental.ExecuteStreamQueryResponse
	(*Ydb_Experimental.GetDiskSpaceUsageResponse)(nil),  // 5: Ydb.Experimental.GetDiskSpaceUsageResponse
}
var file_ydb_experimental_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Experimental.V1.ExperimentalService.UploadRows:input_type -> Ydb.Experimental.UploadRowsRequest
	1, // 1: Ydb.Experimental.V1.ExperimentalService.ExecuteStreamQuery:input_type -> Ydb.Experimental.ExecuteStreamQueryRequest
	2, // 2: Ydb.Experimental.V1.ExperimentalService.GetDiskSpaceUsage:input_type -> Ydb.Experimental.GetDiskSpaceUsageRequest
	3, // 3: Ydb.Experimental.V1.ExperimentalService.UploadRows:output_type -> Ydb.Experimental.UploadRowsResponse
	4, // 4: Ydb.Experimental.V1.ExperimentalService.ExecuteStreamQuery:output_type -> Ydb.Experimental.ExecuteStreamQueryResponse
	5, // 5: Ydb.Experimental.V1.ExperimentalService.GetDiskSpaceUsage:output_type -> Ydb.Experimental.GetDiskSpaceUsageResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_experimental_v1_proto_init() }
func file_ydb_experimental_v1_proto_init() {
	if File_ydb_experimental_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_experimental_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_experimental_v1_proto_goTypes,
		DependencyIndexes: file_ydb_experimental_v1_proto_depIdxs,
	}.Build()
	File_ydb_experimental_v1_proto = out.File
	file_ydb_experimental_v1_proto_rawDesc = nil
	file_ydb_experimental_v1_proto_goTypes = nil
	file_ydb_experimental_v1_proto_depIdxs = nil
}
