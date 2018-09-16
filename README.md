> # Guard <img align="right" width="100" height="100" src=".github/character.png">
>
> 💂‍♂️ Access Control as a Service &mdash; protect any API or sites you want.

## Data structure

```
/defaults
|
|--[rate limiting] (strings in format "%d (rps|rpm|rph)")
|
|--[request limiting] (integer)
|
|--[workplace limiting] (integer)
|
/licenses
|
|--/10000000-2000-4000-8000-160000000001
|  |
|  |--/workplaces
|  |  |
|  |  |--/10000000-2000-4000-8000-160000000002
|  |  |  |
|  |  |  `--/workplace level limits (in future releases)
|  |  |
|  |  `--...
|  |
|  |--creation date (UNIX timestamp)
|  |
|  |--activation date (>= creation date)
|  |
|  |--[active period] (strings passed to time.ParseDuration)
|  |
|  |--[rate limiting] (strings in format "%d (rps|rpm|rph)")
|  |
|  |--[request limiting] (integer)
|  |
|  `--[workplace limiting] (integer)
|
`--/10000000-2000-4000-8000-160000000003
   |
   `--...
```

## API Specification

### POST `/api/v1/license/register`

```json
{
  "license": "10000000-2000-4000-8000-160000000001",
  "activation": {
    "after": "1h", "xor": "^", "when": "2018-01-01 10:55:00",
    "duration": "24h"
  },
  "limits": {
    "rate_limit": "1 rps",
    "request_limit": 1000,
    "workplace_limit": 10
  }
}
```

### PUT `/api/v1/license/extend`

```json
{
  "license": "10000000-2000-4000-8000-160000000001",
  "activation": [
    {"when": "2018-01-01 11:00:00"},
    {"duration": "24h", "rel": "-"}
  ],
  "limits": [
    {"rate_limit": "1 rpm"},
    {"request_limit": null},
    {"workplace_limit": 3, "rel": "+"}
  ]
}
```

### POST `/api/v1/license/check`

```
X-UDID: 10000000-2000-4000-8000-160000000002
X-User: 10000000-2000-4000-8000-160000000001
```

---

[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)
