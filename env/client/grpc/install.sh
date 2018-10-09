#!/usr/bin/env bash

export GUARD_TOKEN=10000000-2000-4000-8000-160000000003
guardctl install -f env/client/grpc/install.yml
guardctl license register -f env/client/grpc/license.register.yml
guardctl license update -f env/client/grpc/license.update.yml
guardctl license create -f env/client/grpc/license.create.yml | guardctl license delete | guardctl license read
