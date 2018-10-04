# TODO
# - gomock
# - protoc
# - statik


.PHONY: protobuf
protobuf:
	protoc -Ipkg/transport/grpc \
	       -Ivendor/github.com/grpc-ecosystem/grpc-gateway \
	       -Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	       --go_out=plugins=grpc,logtostderr=true:pkg/transport/grpc \
	       --grpc-gateway_out=logtostderr=true:pkg/transport/grpc \
	       --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=guard:env/client \
	       common.proto license.proto maintenance.proto


.PHONY: test
test: test-control test-service

.PHONY: test-control
test-control:
	go test -race -tags ctl -v .

.PHONY: test-service
test-service:
	go test -race -v ./...

.PHONY: test-integration
test-integration:
	go test -tags integration -v ./env/test/...
