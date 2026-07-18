## Usage

Sources:

- [Official extension control file](https://github.com/jasonmassie01/pg_sage/blob/a77458aaa3bed42adb7e47e481f476b13413cc3b/pg_sage.control)

`pg_sage` — Provides a preload C DBA agent extension and a newer standalone Go monitoring and guarded-action sidecar.

The reviewed catalog snapshot records version `0.5.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_sage";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sage';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
