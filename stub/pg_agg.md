## Usage

Sources:

- [Official extension control file](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/pg_agg.control)
- [Official upstream README](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/README.md)

`pg_agg` — Experimental aggregate-query support extension; its install script currently creates the untrusted PL/Python language.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`.
The curated compatibility set is `12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_agg";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_agg';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
