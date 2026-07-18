## Usage

Sources:

- [Official extension control file](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/pg_adm.control)
- [Official upstream documentation](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/doc/pg_adm.md)
- [Official upstream README](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/README.md)

`pg_adm` — A PostgreSQL extension providing some administrative tools.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pg_buffercache`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_adm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_adm';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
