## Usage

Sources:

- [README at the reviewed commit](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/README.md)
- [Extension control file](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/pg_queryid.control)
- [SQL function definition](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/pg_queryid--0.0.1.sql)
- [Fingerprinting implementation](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/pg_queryid.c)

`pg_queryid` is a proof-of-concept external query-fingerprinting module for PostgreSQL 14 and later. When preloaded, it replaces core query-ID calculation and can fingerprint referenced objects by name rather than OID, optionally ignoring schemas or queries that touch temporary tables. Creating the extension also exposes `pg_queryid(text) returns bigint` for one statement at a time.

### Configuration and Query

The module must be the last entry in `shared_preload_libraries`, and upstream requires core `compute_query_id` to be disabled when the replacement hook is active.

```conf
shared_preload_libraries = 'pg_stat_statements,pg_queryid'
compute_query_id = off
pg_queryid.use_object_names = on
pg_queryid.ignore_schema = off
pg_queryid.ignore_temp_tables = off
```

After restarting PostgreSQL:

```sql
CREATE EXTENSION pg_queryid;

SELECT pg_queryid('SELECT * FROM public.orders WHERE id = 42');
```

### Caveats

- Upstream labels version `0.0.1` a proof of concept that is not ready for production. The reviewed source was last updated in 2022 and provides no current compatibility statement beyond PostgreSQL 14 or later.
- Name-based fingerprinting can add substantial overhead. With `pg_queryid.ignore_schema = on`, same-named objects in different schemas intentionally share a fingerprint.
- Changing any fingerprint option changes query IDs seen by consumers. Reset `pg_stat_statements` and any other query-ID-indexed state after a change.
- With `pg_queryid.ignore_temp_tables = on`, statements that reference temporary relations are deliberately left without this module's fingerprint and can disappear from query-ID-based monitoring.
- Automatic replacement requires preloading and a restart. Merely running `CREATE EXTENSION` only installs the SQL helper and does not install the parser hook.
