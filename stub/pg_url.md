## Usage

Sources:

- [Official upstream README](https://github.com/semtexzv/pg_url/blob/2741caa7c7ed2a4435253f0d998aac306af7f2da/README.md)
- [Official extension control file (pg_url.control)](https://github.com/semtexzv/pg_url/blob/2741caa7c7ed2a4435253f0d998aac306af7f2da/pg_url.control)
- [Official implementation source](https://github.com/semtexzv/pg_url/blob/2741caa7c7ed2a4435253f0d998aac306af7f2da/src/lib.rs)

`pg_url` — Implements URL manipulation methods as a postgres extension. Using this extension will allow you to work with URLS directly inside database. You could use it to create an index on a URL host. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_url;

CREATE INDEX tbl_url_host on tbl (url_host(url));
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `url_clear_host` is an extension function.
- `url_fragment` is an extension function.
- `url_host` is an extension function.
- `url_path` is an extension function.
- `url_query` is an extension function.
- `url_query_param` is an extension function.
- `url_scheme` is an extension function.
- `url_set_fragment` is an extension function.
- `url_set_host` is an extension function.
- `url_set_path` is an extension function.
- `url_set_query` is an extension function.
- `url_set_query_param` is an extension function.
- `url_set_scheme` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
