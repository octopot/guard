> # 💂‍♂️ Guard [![Tweet][icon_twitter]][twitter_publish] <img align="right" width="126" src=".github/character.png">
> [![Analytics][analytics_pixel]][page_promo]
> Access Control as a Service &mdash; your personal paywall to protect any API or site's content.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![Build Status][icon_build]][page_build]
[![License][icon_license]](LICENSE)

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**October 20, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [ ] v2: [Rate limiting][project_v2]
  - [**December 16, 2018**][project_v2_dl]
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

$ ./env/client/grpc/install.sh localhost:443

$ open http://spec.127.0.0.1.xip.io/

$ make help
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

$ guardctl install -f env/client/grpc/install.yaml

$ export GUARD_TOKEN=10000000-2000-4000-8000-160000000003

$ guardctl license register -f env/client/grpc/license.register.yml
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

$ cat env/client/grpc/license.update.yml | guardctl license update
id: 10000000-2000-4000-8000-160000000004
updated_at: "2018-10-04T09:33:32.487454Z"

$ guardctl license create -f env/client/grpc/license.create.yml | guardctl license delete | guardctl license read
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

$ guardctl license delete -f env/client/grpc/license.delete.yml
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

$ which guard guardctl
/usr/local/bin/guard
/usr/local/bin/guardctl
```

### Binary

```bash
$ export REQ_VER=0.0.1  # all available versions are on https://github.com/kamilsk/guard/releases/
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
# wget -q -O guard.tar.gz
$ curl -sL -o guard.tar.gz \
       https://github.com/kamilsk/guard/releases/download/"${REQ_VER}/guard_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf guard.tar.gz -C "${GOPATH}"/bin/ && rm guard.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/guard:latest
```

### From source code

```bash
$ egg github.com/kamilsk/guard -- make test install
# or use mirror
$ egg bitbucket.org/kamilsk/guard -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

## Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

<sup id="egg">1</sup> The project is still in prototyping. [↩](#anchor-egg)

---

[![@kamilsk][icon_tw_author]](https://twitter.com/ikamilsk)
[![@octolab][icon_tw_sponsor]](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[analytics_pixel]: https://ga-beacon.appspot.com/UA-109817251-26/guard/readme?pixel

[icon_gitter]:     https://badges.gitter.im/Join%20Chat.svg
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[page_build]:      https://travis-ci.org/kamilsk/guard
[page_promo]:      https://github.com/kamilsk/guard
[page_research]:   ../../tree/research

[project_v1]:      https://github.com/kamilsk/guard/projects/1
[project_v1_dl]:   https://github.com/kamilsk/guard/milestone/1
[project_v2]:      https://github.com/kamilsk/guard/projects/2
[project_v2_dl]:   https://github.com/kamilsk/guard/milestone/2

[twitter_publish]: https://twitter.com/intent/tweet?text=Access%20Control%20as%20a%20Service&url=https://kamilsk.github.io/guard/&via=ikamilsk&hashtags=go,service
