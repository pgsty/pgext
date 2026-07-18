## Usage

Sources:

- [Official extension control file](https://github.com/df7cb/postgresql-hamradio/blob/4e2181d68b62d224b120ce1bb38d2c23ee63e937/hamradio.control)

`hamradio` — ham radio extension

The reviewed catalog snapshot records version `1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `postgis`, `uctext`.

```sql
CREATE EXTENSION "hamradio";
SELECT extversion
FROM pg_extension
WHERE extname = 'hamradio';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
