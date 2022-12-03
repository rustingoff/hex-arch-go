proto_msg:
	protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto \
	  internal/adapters/framework/left/grpc/proto/number_msg.proto

proto_svc:
	protoc --go-grpc_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto \
    	  internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto
