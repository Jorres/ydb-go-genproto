proto:
	protoc --go_out=. --go_opt=module=github.com/ydb-platform/ydb-go-genproto --go-grpc_out=. --go-grpc_opt=module=github.com/ydb-platform/ydb-go-genproto -Iapi -Iapi/protos -Iapi/protos/annotations api/{*.proto,protos/*.proto}
