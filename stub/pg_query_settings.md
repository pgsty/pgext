## Usage

Sources:

- [Official upstream README](https://github.com/dalibo/pg_query_settings/blob/d1e9cbf00a33c4c11062c65605a7e129f3ebc5ea/README.md)
- [Official extension control file (pg_query_settings.control)](https://github.com/dalibo/pg_query_settings/blob/d1e9cbf00a33c4c11062c65605a7e129f3ebc5ea/pg_query_settings.control)
- [Official extension SQL (pg_query_settings--0.1.sql)](https://github.com/dalibo/pg_query_settings/blob/d1e9cbf00a33c4c11062c65605a7e129f3ebc5ea/pg_query_settings--0.1.sql)

`pg_query_settings` — The original idea was to configure a specific value of the work_mem parameter for a specific query. Currently, all query parameters may have a customizable value. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_query_settings;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgqs_config` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
