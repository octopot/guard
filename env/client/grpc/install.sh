#!/usr/bin/env bash

set -euo pipefail

HOST=${1:-127.0.0.1:8092}
GUARD_TOKEN=10000000-2000-4000-8000-160000000003

guardctl install                  -f env/client/grpc/install.yml           --grpc-host=${HOST}

guardctl license register         -f env/client/grpc/license.register.yml  --grpc-host=${HOST}
guardctl license update           -f env/client/grpc/license.update.yml    --grpc-host=${HOST}
guardctl license delete           -f env/client/grpc/license.delete.yml    --grpc-host=${HOST}
guardctl license restore          -f env/client/grpc/license.restore.yml   --grpc-host=${HOST}
guardctl license read             -f env/client/grpc/license.read.yml      --grpc-host=${HOST}

guardctl license employee add     -f env/client/grpc/license.employee.yml  --grpc-host=${HOST}
guardctl license employee delete  -f env/client/grpc/license.employee.yml  --grpc-host=${HOST}
guardctl license employee add     -f env/client/grpc/license.employee.yml  --grpc-host=${HOST}

guardctl license workplace add    -f env/client/grpc/license.workplace.yml --grpc-host=${HOST}
guardctl license workplace push   -f env/client/grpc/license.workplace.yml --grpc-host=${HOST}
guardctl license workplace delete -f env/client/grpc/license.workplace.yml --grpc-host=${HOST}
guardctl license workplace add    -f env/client/grpc/license.workplace.yml --grpc-host=${HOST}

guardctl license create           -f env/client/grpc/license.create.yml    --grpc-host=${HOST} | \
guardctl license delete                                                    --grpc-host=${HOST} | \
guardctl license read                                                      --grpc-host=${HOST}
