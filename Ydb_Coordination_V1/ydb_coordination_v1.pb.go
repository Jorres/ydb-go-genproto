// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: ydb_coordination_v1.proto

package Ydb_Coordination_V1

import (
	Ydb_Coordination "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Coordination"
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

var File_ydb_coordination_v1_proto protoreflect.FileDescriptor

var file_ydb_coordination_v1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x79, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x59, 0x64, 0x62,
	0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xca, 0x03, 0x0a, 0x13, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x44, 0x72,
	0x6f, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x72, 0x6f,
	0x70, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a,
	0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6b, 0x0a, 0x18,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_ydb_coordination_v1_proto_goTypes = []interface{}{
	(*Ydb_Coordination.SessionRequest)(nil),       // 0: Ydb.Coordination.SessionRequest
	(*Ydb_Coordination.CreateNodeRequest)(nil),    // 1: Ydb.Coordination.CreateNodeRequest
	(*Ydb_Coordination.AlterNodeRequest)(nil),     // 2: Ydb.Coordination.AlterNodeRequest
	(*Ydb_Coordination.DropNodeRequest)(nil),      // 3: Ydb.Coordination.DropNodeRequest
	(*Ydb_Coordination.DescribeNodeRequest)(nil),  // 4: Ydb.Coordination.DescribeNodeRequest
	(*Ydb_Coordination.SessionResponse)(nil),      // 5: Ydb.Coordination.SessionResponse
	(*Ydb_Coordination.CreateNodeResponse)(nil),   // 6: Ydb.Coordination.CreateNodeResponse
	(*Ydb_Coordination.AlterNodeResponse)(nil),    // 7: Ydb.Coordination.AlterNodeResponse
	(*Ydb_Coordination.DropNodeResponse)(nil),     // 8: Ydb.Coordination.DropNodeResponse
	(*Ydb_Coordination.DescribeNodeResponse)(nil), // 9: Ydb.Coordination.DescribeNodeResponse
}
var file_ydb_coordination_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Coordination.V1.CoordinationService.Session:input_type -> Ydb.Coordination.SessionRequest
	1, // 1: Ydb.Coordination.V1.CoordinationService.CreateNode:input_type -> Ydb.Coordination.CreateNodeRequest
	2, // 2: Ydb.Coordination.V1.CoordinationService.AlterNode:input_type -> Ydb.Coordination.AlterNodeRequest
	3, // 3: Ydb.Coordination.V1.CoordinationService.DropNode:input_type -> Ydb.Coordination.DropNodeRequest
	4, // 4: Ydb.Coordination.V1.CoordinationService.DescribeNode:input_type -> Ydb.Coordination.DescribeNodeRequest
	5, // 5: Ydb.Coordination.V1.CoordinationService.Session:output_type -> Ydb.Coordination.SessionResponse
	6, // 6: Ydb.Coordination.V1.CoordinationService.CreateNode:output_type -> Ydb.Coordination.CreateNodeResponse
	7, // 7: Ydb.Coordination.V1.CoordinationService.AlterNode:output_type -> Ydb.Coordination.AlterNodeResponse
	8, // 8: Ydb.Coordination.V1.CoordinationService.DropNode:output_type -> Ydb.Coordination.DropNodeResponse
	9, // 9: Ydb.Coordination.V1.CoordinationService.DescribeNode:output_type -> Ydb.Coordination.DescribeNodeResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_coordination_v1_proto_init() }
func file_ydb_coordination_v1_proto_init() {
	if File_ydb_coordination_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_coordination_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_coordination_v1_proto_goTypes,
		DependencyIndexes: file_ydb_coordination_v1_proto_depIdxs,
	}.Build()
	File_ydb_coordination_v1_proto = out.File
	file_ydb_coordination_v1_proto_rawDesc = nil
	file_ydb_coordination_v1_proto_goTypes = nil
	file_ydb_coordination_v1_proto_depIdxs = nil
}
