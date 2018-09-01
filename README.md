> # Paymaster
>
> 👨‍💼 Payment Access as a Service. The prototype.

## Data structure

```
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
|  |  |--/10000000-2000-4000-8000-160000000003
|  |  |  |
|  |  |  `--/workplace level limits (in future releases)
|  |  |
|  |  `--/10000000-2000-4000-8000-160000000004
|  |     |
|  |     `--/workplace level limits (in future releases)
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
`--/10000000-2000-4000-8000-160000000005
   |
   `--...
```

## API Specification

### POST `/api/v1/register`

```json
{
  "license": "10000000-2000-4000-8000-160000000001",
  "activation": {
    "when": "1h3m | 2018-01-01 10:55",
    "duration": "24h"
  },
  "limits": {
    "rate_limit": "1 rps",
    "request_limit": 1000,
    "workplace_limit": 10
  }
}
```

---

[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)
