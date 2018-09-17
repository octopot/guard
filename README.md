> # 💂‍♂️ Guard [![Tweet][icon_twitter]][twitter_publish] <img align="right" width="100" src=".github/character.png">
>
> Access Control as a Service &mdash; protect any API or sites you want.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![License][icon_license]](LICENSE)

## Roadmap

- [ ] v1: [MVP][project_v1]
  - [**September 20, 2018**][project_v1_dl]
  - Main concepts and working prototype.

## Motivation

...

## Quick start

...

## Specification

### API

You can find API specification [here](env/client/rest.http).

### CLI

...

#### Bash and Zsh completions

...

## Installation

### Brew

...

### Binary

...

### Docker Hub

...

### From source code

```bash
$ egg github.com/kamilsk/guard^1.0.0 -- make test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/guard^1.0.0 -- make test install
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

[icon_gitter]:     https://badges.gitter.im/Join%20Chat.svg
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[project_v1]:      https://github.com/kamilsk/guard/projects/1
[project_v1_dl]:   https://github.com/kamilsk/guard/milestone/1

[twitter_publish]: https://twitter.com/intent/tweet?text=Access%20Control%20as%20a%20Service&url=https://kamilsk.github.io/guard/&via=ikamilsk&hashtags=go,service
