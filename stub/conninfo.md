## Usage

Sources:

- [Official README](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/README.md)
- [Extension SQL](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/conninfo--1.0.sql)
- [libpq wrapper implementation](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/conninfo.c)

`conninfo` exposes the server's linked libpq parser and defaults as SQL functions. Use it to split a connection string, check that libpq can parse it, or inspect effective default connection parameters; it never attempts a database connection.

### Core Workflow

```sql
CREATE EXTENSION conninfo;

SELECT *
FROM conninfo_parse(
    'host=db.internal port=5432 dbname=app user=reader connect_timeout=5'
);

SELECT conninfo_validate('host=db.internal connect_timeout=5');

SELECT keyword, envvar, compiled, value
FROM conninfo_defaults()
WHERE keyword IN ('host', 'port', 'user', 'dbname');
```

`conninfo_parse(text)` returns one `keyword`/`value` row per supplied option and raises an error for a malformed or unknown option. `conninfo_validate(text)` returns `false` and emits a warning for the same parse failure. A `true` result means only that libpq accepted the syntax and option names; it does not validate reachability, credentials, TLS, or the semantic suitability of values.

### Function Index

- `conninfo_parse(text)` returns `SETOF record` with `keyword text` and `value text`.
- `conninfo_validate(text)` returns `boolean`.
- `conninfo_defaults()` returns `keyword`, `envvar`, `compiled`, `value`, `label`, `dispchar`, and `dispsize` from `PQconndefaults()`.

### Security and Compatibility

Parsed values and defaults can contain passwords, including values inherited from `PGPASSWORD` or a libpq service/environment setting. Do not log query results, expose the functions through an untrusted API, or grant access without considering credential disclosure.

The functions use the libpq library and environment attached to the PostgreSQL server process, which may differ from a client application's libpq build and environment. Validate with the same client runtime when client-specific behavior matters.
