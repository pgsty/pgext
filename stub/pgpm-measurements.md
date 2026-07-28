## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/measurements/README.md)
- [Official extension control file (pgpm-measurements.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/measurements/pgpm-measurements.control)
- [Official extension SQL (pgpm-measurements--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/measurements/sql/pgpm-measurements--0.15.5.sql)

`pgpm-measurements` — @pgpm/measurements provides a standardized system for tracking measurements and quantities in PostgreSQL applications. This package defines a schema for storing measurement types with their units and descriptions, enabling consistent metric tracking across your application. Use it when an application needs this specific database capability. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-measurements";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `measurements.quantities` is a table installed or managed by the extension.
- `measurements` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
