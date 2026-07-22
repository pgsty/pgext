## Usage

Sources:

- [Official extension control file](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/pg_agg.control)
- [Official extension SQL](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/pg_agg--1.0.sql)

`pg_agg` 1.0 does not implement the aggregate-query index support claimed by its control-file comment. The complete install SQL only executes `CREATE OR REPLACE LANGUAGE plpython3u`; it defines no aggregate, function, type, operator, table, index access method, or configuration parameter.

### Core Workflow

Inspect availability before installation, because the extension needs the server's untrusted PL/Python 3 module:

```sql
SELECT name, default_version
FROM pg_available_extensions
WHERE name IN ('pg_agg', 'plpython3u');

CREATE EXTENSION pg_agg;

SELECT lanname, lanpltrusted
FROM pg_language
WHERE lanname = 'plpython3u';
```

The only observable effect of `CREATE EXTENSION pg_agg` is creation or replacement of the `plpython3u` language. There is no documented query workflow beyond that side effect.

### Important Objects

- `plpython3u` is the sole object named by the version 1.0 install script.
- `pg_agg` is an extension record wrapping that one-line install action; it exposes no user-facing API of its own.

### Operational Notes

Creating an untrusted procedural language requires superuser-level authority and the matching PostgreSQL PL/Python server package. Untrusted Python functions execute with database-server privileges, so enabling the language is a security decision even though this extension creates no functions itself. Do not select `pg_agg` for aggregate indexes or query acceleration unless a later authoritative upstream version supplies an actual implementation; the pinned 1.0 source does not.
