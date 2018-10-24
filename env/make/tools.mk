# TODO issue#environment
# - gomock
# - protoc
# - statik


.PHONY: protobuf
protobuf:
	protoc -Ipkg/transport/grpc/protobuf \
	       -Ivendor/github.com/grpc-ecosystem/grpc-gateway \
	       -Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	       --go_out=plugins=grpc,logtostderr=true:pkg/transport/grpc/protobuf \
	       --grpc-gateway_out=logtostderr=true:pkg/transport/grpc/protobuf \
	       --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=guard:env/client \
	       common.proto license.proto maintenance.proto


.PHONY: test
test:
	go test -race -v ./...

.PHONY: test-integration
test-integration:
	go test -tags integration -v ./env/test/...
