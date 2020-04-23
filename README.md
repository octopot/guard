> # 💂‍♂️ Guard
>
> Paywall as a Service &mdash; your personal guard to restrict access to content via a paid subscription.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**October 20, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [ ] v2: [Rate limiting][project_v2]
  - [**Someday, 20xx**][project_v2_dl]
  - Request rate limiting and metrics related to license checking.

## Motivation

- We have to limit access to some part of our content and APIs.

## Quick start

Requirements:

- Docker 18.06.0-ce or above
- Docker Compose 1.22.0 or above
- Go 1.11 or above
- GNU Make 3.81 or above

```bash
$ make demo up clean status

     Name                    Command               State                        Ports
-----------------------------------------------------------------------------------------------------------
guard_db_1        docker-entrypoint.sh postgres    Up      5432/tcp
guard_legacy_1    docker-php-entrypoint php-fpm    Up      9000/tcp
guard_server_1    nginx -g daemon off;             Up      0.0.0.0:443->443/tcp, 0.0.0.0:80->80/tcp
guard_service_1   service run --with-profili ...   Up      8080/tcp, 8090/tcp, 8091/tcp, 8092/tcp, 8093/tcp
guard_spec_1      sh /usr/share/nginx/docker ...   Up      80/tcp, 8080/tcp

$ make install

$ ./env/client/rpc/install.sh localhost:443

$ open http://spec.127.0.0.1.xip.io/

$ make help

$ make demo destroy
```

## Specification

### API

You can find API specification [here](env/client/rest.http).

### CLI

You can use `guard` to start the HTTP server and `guardctl` to execute
[CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) operations.

<details>
<summary><strong>Service command-line interface</strong></summary>

```bash
$ make service-install

$ guard --help
  Guard Service
  
  Usage:
    guard [command]
  
  Available Commands:
    completion  Print Bash or Zsh completion
    help        Help about any command
    run         Start HTTP server
    version     Show application version
  
  Flags:
    -h, --help   help for guard
  
  Use "guard [command] --help" for more information about a command.
```
</details>

<details>
<summary><strong>Client command-line interface</strong></summary>

```bash
$ make control-install

$ guardctl --help
  Guard Control
  
  Usage:
    guardctl [command]
  
  Available Commands:
    completion  Print Bash or Zsh completion
    help        Help about any command
    license     Guard License
    version     Show application version
  
  Flags:
    -h, --help   help for guardctl
  
  Use "guardctl [command] --help" for more information about a command.

$ guardctl install -f env/client/rpc/install.yaml

$ export GUARD_TOKEN=10000000-2000-4000-8000-160000000003

$ guardctl license register -f env/client/rpc/license.register.yml
id: 10000000-2000-4000-8000-160000000004

$ echo '{id: 10000000-2000-4000-8000-160000000004}' | guardctl license read
contract:
  rate:
    unit: rph
    value: 10
  requests: 1000
  since: "2018-09-29T17:11:43.264Z"
  workplaces: 10
created_at: "2018-10-04T09:32:19.102216Z"
id: 10000000-2000-4000-8000-160000000004

$ cat env/client/rpc/license.update.yml | guardctl license update
id: 10000000-2000-4000-8000-160000000004
updated_at: "2018-10-04T09:33:32.487454Z"

$ guardctl license create -f env/client/rpc/license.create.yml | guardctl license delete | guardctl license read
contract:
  rate:
    unit: rph
    value: 10
  requests: 1000
  since: "2018-09-29T17:11:43.264Z"
  workplaces: 10
created_at: "2018-10-04T09:57:16.656346Z"
deleted_at: "2018-10-04T09:57:16.666664Z"
id: 9ba7b564-3248-4401-b853-9dc32559b95b
updated_at: "2018-10-04T09:57:16.666664Z"

$ guardctl license delete -f env/client/rpc/license.delete.yml
deleted_at: "2018-10-04T09:58:27.365193Z"
id: 10000000-2000-4000-8000-160000000004

$ echo '{id: 10000000-2000-4000-8000-160000000004}' | guardctl license restore | guardctl license read
contract:
  rate:
    unit: rpd
    value: 10
  requests: 1000
  since: "2018-09-29T17:11:43.264Z"
  until: "2018-09-29T17:11:43.264Z"
  workplaces: 10
created_at: "2018-10-04T09:54:57.643041Z"
id: 10000000-2000-4000-8000-160000000004
updated_at: "2018-10-04T09:59:18.833134Z"
```
</details>

#### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

Service:

```bash
$ guard completion -f bash > /path/to/bash_completion.d/guard.sh
$ guard completion -f zsh  > /path/to/zsh-completions/_guard.zsh
```

Client:

```bash
$ guardctl completion -f bash > /path/to/bash_completion.d/guardctl.sh
$ guardctl completion -f zsh  > /path/to/zsh-completions/_guardctl.zsh
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/guard
```

### Binary

```bash
$ export REQ_VER=1.0.0  # all available versions are on https://github.com/kamilsk/guard/releases/
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
# wget -q -O guard.tar.gz
$ curl -sL -o guard.tar.gz \
       https://github.com/kamilsk/guard/releases/download/"${REQ_VER}/guard_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf guard.tar.gz -C "${GOPATH}"/bin/ && rm guard.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/guard:1.x
# or use mirror
$ docker pull quay.io/kamilsk/guard:1.x
```

### From source code

```bash
$ egg github.com/kamilsk/guard@^1.0.0 -- make test install
# or use mirror
$ egg bitbucket.org/kamilsk/guard@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/octopot/guard
[build.icon]:       https://travis-ci.com/octopot/guard.svg?branch=master
[design.page]:      https://www.notion.so/octolab/Guard-b6579c8a78714b8787f631508eff5451?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://octopot.github.io/guard/
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg

[project_v1]:       https://github.com/octopot/guard/projects/1
[project_v1_dl]:    https://github.com/octopot/guard/milestone/1
[project_v2]:       https://github.com/octopot/guard/projects/2
[project_v2_dl]:    https://github.com/octopot/guard/milestone/2
