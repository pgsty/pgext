## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pgjsonrpc/pgjsonrpc-0.0.1/README.md)
- [Official extension control file (pgjsonrpc.control)](https://api.pgxn.org/src/pgjsonrpc/pgjsonrpc-0.0.1/pgjsonrpc.control)
- [Official extension SQL (pgjsonrpc.sql)](https://api.pgxn.org/src/pgjsonrpc/pgjsonrpc-0.0.1/sql/pgjsonrpc.sql)

`pgjsonrpc` — Implementation of JSON-RPC 2.0 as a PostgreSQL extension. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgjsonrpc;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `jsonrpc.echo(p_request JSON)` is an extension function and returns `JSON`.
- `jsonrpc.error_response(p_request JSON, p_code INTEGER, p_message TEXT, p_data JSON)` is an extension function and returns `JSON`.
- `jsonrpc.execute(p_request TEXT)` is an extension function and returns `JSON`.
- `jsonrpc.get_response(p_request JSON, p_jsonrpc TEXT, p_id TEXT, p_result JSON, p_code INTEGER, p_message TEXT)` is an extension function and returns `JSON`.
- `jsonrpc.success_response(p_request JSON, p_result ANYELEMENT)` is an extension function and returns `JSON`.
- `jsonrpc.methods` is a table installed or managed by the extension.
- `jsonrpc` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
