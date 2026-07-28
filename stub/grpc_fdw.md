## Usage

Sources:

- [Official upstream README](https://github.com/hydradb/grpc_fdw/blob/a594a3574b51b7beb79a10abbd522324ed35dc6e/README.md)
- [Official extension control file (grpc_fdw.control)](https://github.com/hydradb/grpc_fdw/blob/a594a3574b51b7beb79a10abbd522324ed35dc6e/grpc_fdw.control)
- [Official implementation source](https://github.com/hydradb/grpc_fdw/blob/a594a3574b51b7beb79a10abbd522324ed35dc6e/src/lib.rs)

`grpc_fdw` — Foreign data wrapper that delegates table operations to services over gRPC. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION grpc_fdw;

CREATE FOREIGN DATA WRAPPER grpc_fdw_handler HANDLER grpc_fdw_handler NO VALIDATOR;
CREATE SERVER user_srv FOREIGN DATA WRAPPER grpc_fdw_handler OPTIONS (server_uri 'http://[::1]:50051');
CREATE FOREIGN TABLE users (
    id integer,
    name text,
    email text
) SERVER user_srv OPTIONS (
    table_option '1',
    table_option2 '2'
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `grpc_fdw_handler()` is an extension function and returns `pg_sys::Datum`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
