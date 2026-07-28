## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/streaming_lag/streaming_lag-0.0.1/README.md)
- [Official extension control file (streaming_lag.control)](https://api.pgxn.org/src/streaming_lag/streaming_lag-0.0.1/streaming_lag.control)
- [Official extension SQL (streaming_lag--0.0.1.sql)](https://api.pgxn.org/src/streaming_lag/streaming_lag-0.0.1/streaming_lag--0.0.1.sql)

`streaming_lag` — streaming_lag is an experimental extension to measure the lag of a streaming slave in units of time instead of bytes. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION streaming_lag;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `streaming_lag` is an extension-defined view.
- `streaming_lag_data` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
