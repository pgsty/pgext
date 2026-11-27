## Usage

Sources:

- [doc/chdb.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/doc/chdb.md)
- [doc/chdb_hook.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/doc/chdb_hook.md)
- [chdb.control](https://api.pgxn.org/src/chdb/chdb-0.1.4/chdb.control)
- [sql/chdb.sql](https://api.pgxn.org/src/chdb/chdb-0.1.4/sql/chdb.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/CHANGELOG.md)
- [LICENSE.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/LICENSE.md)

`chdb` runs chDB queries in an isolated helper process. Distribution 0.1.4 retains SQL/control version 0.1. It supports PostgreSQL 15 and later on Linux or macOS, using the supplied helper and a compatible libchdb.

### Query Workflow

A superuser installs this untrusted, relocatable extension; shared preload is not required:

```sql
CREATE EXTENSION chdb;
SELECT *
FROM chdb_query('SELECT number, number * number FROM numbers(5)')
AS result(number bigint, square bigint);
SELECT pgchdb_version();
```

`chdb_query(text)` returns records and requires an explicit column definition matching the query result. Ordinary roles receive no execution permission by default: grant `EXECUTE` on `chdb_query(text)` only to approved, existing roles. Each call creates and removes a temporary chDB database, so objects created by one call do not persist to the next.

### Resource and Credential Settings

Superuser settings `chdb.max_memory`, `chdb.max_threads` and `chdb.max_parsing_threads` control helper resource usage. The memory default is now 4 GB; `0` removes that limit. Both thread settings default to `0`, allowing chDB to choose, so set finite limits for large workloads.

Version 0.1.4 changes access to host cloud credentials. `chdb.host_credentials_roles` defaults to an empty list. Only members of the existing PostgreSQL roles listed by a superuser may use host AWS credentials; other callers make anonymous requests when they supply no keys. This setting does not grant execution permission on the query function. Listed members inherit the access available to the host credentials.

### Data and Upgrade Boundaries

Queries can reach data sources accessible to chDB. Review endpoint access, explicit credentials and result types. Complex arrays, maps, nested values and encoding conversions require matching PostgreSQL types; use representative round-trip checks. The helper isolates its memory from PostgreSQL, but still consumes host CPU, memory, disk and network resources.

The library reports its full version through `pgchdb_version()`. Existing SQL version 0.1 needs no SQL upgrade for the 0.1.4 library update. Install matching library/helper files and use fresh sessions to load the new binaries. The release also adds an optional extension wrapper for `chdb_hook`, documented separately.
