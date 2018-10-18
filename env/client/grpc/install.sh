#!/usr/bin/env bash

set -euo pipefail

export GUARD_TOKEN=10000000-2000-4000-8000-160000000003

guardctl install -f env/client/grpc/install.yml

guardctl license register -f env/client/grpc/license.register.yml
guardctl license update   -f env/client/grpc/license.update.yml
guardctl license delete   -f env/client/grpc/license.delete.yml
guardctl license restore  -f env/client/grpc/license.restore.yml
guardctl license read     -f env/client/grpc/license.read.yml

guardctl license employee add    -f env/client/grpc/license.employee.yml
guardctl license employee delete -f env/client/grpc/license.employee.yml

guardctl license workplace add    -f env/client/grpc/license.workplace.yml
guardctl license workplace push   -f env/client/grpc/license.workplace.yml
guardctl license workplace delete -f env/client/grpc/license.workplace.yml

guardctl license create -f env/client/grpc/license.create.yml | \
guardctl license delete | \
guardctl license read
