.PHONY: protobuf
protobuf:
	protoc -Ipkg/transport/grpc/ --go_out=plugins=grpc:pkg/transport/grpc license.proto


.PHONY: test
test: test-control test-service

.PHONY: test-control
test-control:
	go test -race -tags 'cli ctl' -v .

.PHONY: test-service
test-service:
	go test -race -v ./...
