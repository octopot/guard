ifndef PACKAGE
$(error Please define PACKAGE variable)
endif

_commit   = -X $(PACKAGE)/pkg/cmd.commit=$(shell git rev-parse --short HEAD)
_date     = -X $(PACKAGE)/pkg/cmd.date=$(shell date -u +%FT%X%Z)
_version  = -X $(PACKAGE)/pkg/cmd.version=dev
LDFLAGS   = -ldflags '-s -w $(_commit) $(_date) $(_version)'
CTL_FILES = $(PACKAGE)/cmd/guardctl
CTL_FLAGS = $(LDFLAGS)
SRV_FILES = $(PACKAGE)/cmd/guard
SRV_FLAGS = $(LDFLAGS)


.PHONY: __cmd__
__cmd__:
	go run $(BUILD_FLAGS) $(BUILD_FILES) $(ARGS)

.PHONY: __install__
__install__:
	go install -i $(BUILD_FLAGS) $(BUILD_FILES)

.PHONY: __ctl__
__ctl__:
	$(eval BUILD_FLAGS = $(CTL_FLAGS))
	$(eval BUILD_FILES = $(CTL_FILES))

.PHONY: __srv__
__srv__:
	$(eval BUILD_FLAGS = $(SRV_FLAGS))
	$(eval BUILD_FILES = $(SRV_FILES))


.PHONY: control-cmd-help
control-cmd-help: ARGS = help
control-cmd-help: __ctl__ __cmd__

.PHONY: control-cmd-version
control-cmd-version: ARGS = version
control-cmd-version: __ctl__ __cmd__

.PHONY: control-install
control-install: __ctl__ __install__


.PHONY: service-cmd-help
service-cmd-help: ARGS = help
service-cmd-help: __srv__ __cmd__

.PHONY: service-cmd-migrate
service-cmd-migrate: ARGS = migrate
service-cmd-migrate: __srv__ __cmd__

.PHONY: service-cmd-run
service-cmd-run: ARGS = run -H 127.0.0.1:8080 --with-profiling --with-monitoring --with-grpc-gateway
service-cmd-run: __srv__ __cmd__

.PHONY: service-cmd-version
service-cmd-version: ARGS = version
service-cmd-version: __srv__ __cmd__

.PHONY: service-install
service-install: __srv__ __install__


.PHONY: install
install:
	@(make control-install)
	@(make service-install)
