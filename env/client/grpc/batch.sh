#!/usr/bin/env bash

guardctl install -f env/client/grpc/install.yml
export GUARD_TOKEN=10000000-2000-4000-8000-160000000003
guardctl license register -f env/client/grpc/license.register.yml
guardctl license update -f env/client/grpc/license.update.yml
guardctl license create -f env/client/grpc/license.create.yml | guardctl license delete | guardctl license read
