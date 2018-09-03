LDFLAGS     ?= -ldflags '-s -w -X main.version=dev -X main.commit=$(shell git rev-parse --short HEAD)'
BUILD_FILES ?= main.go


.PHONY: cmd-help
cmd-help:
	go run $(LDFLAGS) $(BUILD_FILES) help

.PHONY: cmd-version
cmd-version:
	go run $(LDFLAGS) $(BUILD_FILES) version


.PHONY: dev-server
dev-server:
	go run $(LDFLAGS) $(BUILD_FILES) run -H 127.0.0.1:8080 --with-profiling --with-monitoring
