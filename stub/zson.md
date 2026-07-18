## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/zson.control)

`zson` — Dictionary-trained compressed JSONB-compatible type with transparent casts and JSONB operators.

The reviewed catalog snapshot records version `1.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "zson";
SELECT extversion
FROM pg_extension
WHERE extname = 'zson';
```

The upstream project is associated with `Postgres Professional`; verify its current support, license, packaging, and deployment boundary from the linked source.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
