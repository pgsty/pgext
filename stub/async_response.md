## Usage

Sources:

- [Official upstream README](https://github.com/arturformella/async_response/blob/24a278e4e6eb244009efa96c1859f359a028cf70/README.md)
- [Official extension control file (async_response.control)](https://github.com/arturformella/async_response/blob/24a278e4e6eb244009efa96c1859f359a028cf70/async_response.control)
- [Official extension SQL (async_response--1.0.sql)](https://github.com/arturformella/async_response/blob/24a278e4e6eb244009efa96c1859f359a028cf70/async_response--1.0.sql)

`async_response` — boolean async_response(port INTEGER, channel TEXT, aspect TEXT, data TEXT) in your complex query to send immediately to REDIS from PostgreSQL:. Use it for the corresponding SQL or database utility workflow. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION async_response;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `async_response(INTEGER,TEXT,TEXT,TEXT)` is an extension function and returns `boolean`.
- `async_response(TEXT,TEXT,TEXT,TEXT)` is an extension function and returns `boolean`.
- `async_response_async(INTEGER,TEXT,TEXT,TEXT)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
