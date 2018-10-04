#!/usr/bin/env bash

guardctl license register -f env/client/grpc/license.register.yml
guardctl license update -f env/client/grpc/license.update.yml
guardctl license create -f env/client/grpc/license.create.yml | guardctl license delete | guardctl license read
