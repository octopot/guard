.PHONY: protobuf
protobuf:
	protoc -Ipkg/transport/grpc/ --go_out=plugins=grpc:pkg/transport/grpc license.proto
