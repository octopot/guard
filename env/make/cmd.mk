ifndef PACKAGE
$(error Please define PACKAGE variable)
endif


LDFLAGS   = -ldflags '-s -w -X $(PACKAGE)/cmd.version=dev -X $(PACKAGE)/cmd.commit=$(shell git rev-parse --short HEAD)'
CTL_FLAGS = -tags 'cli ctl' $(LDFLAGS)
CTL_BUILD = cli.go guardctl.go
SRV_FLAGS = $(LDFLAGS)
SRV_BUILD = cli.go main.go


.PHONY: __cmd__
__cmd__:
	go run $(BUILD_FLAGS) $(BUILD_FILES) $(ARGS)

.PHONY: __build__
__build__:
	go build -o $(BIN) -i $(BUILD_FLAGS) $(BUILD_FILES)
	chmod +x $(BIN)
	mv $(BIN) $(GOPATH)/bin/$(BIN)


.PHONY: control-cmd-help
control-cmd-help: BUILD_FLAGS = $(CTL_FLAGS)
control-cmd-help: BUILD_FILES = $(CTL_BUILD)
control-cmd-help: ARGS = help
control-cmd-help: __cmd__

.PHONY: control-cmd-version
control-cmd-version: BUILD_FLAGS = $(CTL_FLAGS)
control-cmd-version: BUILD_FILES = $(CTL_BUILD)
control-cmd-version: ARGS = version
control-cmd-version: __cmd__

.PHONY: control-install
control-install: BIN = guardctl
control-install: BUILD_FLAGS = $(CTL_FLAGS)
control-install: BUILD_FILES = $(CTL_BUILD)
control-install: __build__


.PHONY: service-cmd-help
service-cmd-help: BUILD_FLAGS = $(SRV_FLAGS)
service-cmd-help: BUILD_FILES = $(SRV_BUILD)
service-cmd-help: ARGS = help
service-cmd-help: __cmd__

.PHONY: service-cmd-version
service-cmd-version: BUILD_FLAGS = $(SRV_FLAGS)
service-cmd-version: BUILD_FILES = $(SRV_BUILD)
service-cmd-version: ARGS = version
service-cmd-version: __cmd__

.PHONY: service-cmd-run
service-cmd-run: BUILD_FLAGS = $(SRV_FLAGS)
service-cmd-run: BUILD_FILES = $(SRV_BUILD)
service-cmd-run: ARGS = run -H 127.0.0.1:8080 --with-profiling --with-monitoring
service-cmd-run: __cmd__

.PHONY: service-install
service-install: BIN = guard
service-install: BUILD_FLAGS = $(SRV_FLAGS)
service-install: BUILD_FILES = $(SRV_BUILD)
service-install: __build__
