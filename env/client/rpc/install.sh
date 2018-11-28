#!/usr/bin/env bash

set -euo pipefail

HOST=${1:-127.0.0.1:8092}
GUARD_TOKEN=10000000-2000-4000-8000-160000000003

guardctl install                  -f env/client/rpc/install.yml           --grpc-host=${HOST}

guardctl license register         -f env/client/rpc/license.register.yml  --grpc-host=${HOST}
guardctl license update           -f env/client/rpc/license.update.yml    --grpc-host=${HOST}
guardctl license delete           -f env/client/rpc/license.delete.yml    --grpc-host=${HOST}
guardctl license restore          -f env/client/rpc/license.restore.yml   --grpc-host=${HOST}
guardctl license read             -f env/client/rpc/license.read.yml      --grpc-host=${HOST}

guardctl license employee add     -f env/client/rpc/license.employee.yml  --grpc-host=${HOST}
guardctl license employee delete  -f env/client/rpc/license.employee.yml  --grpc-host=${HOST}
guardctl license employee add     -f env/client/rpc/license.employee.yml  --grpc-host=${HOST}

guardctl license workplace add    -f env/client/rpc/license.workplace.yml --grpc-host=${HOST}
guardctl license workplace push   -f env/client/rpc/license.workplace.yml --grpc-host=${HOST}
guardctl license workplace delete -f env/client/rpc/license.workplace.yml --grpc-host=${HOST}
guardctl license workplace add    -f env/client/rpc/license.workplace.yml --grpc-host=${HOST}

guardctl license create           -f env/client/rpc/license.create.yml    --grpc-host=${HOST} | \
guardctl license delete                                                   --grpc-host=${HOST} | \
guardctl license read                                                     --grpc-host=${HOST}
