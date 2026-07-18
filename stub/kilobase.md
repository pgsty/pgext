## Usage

Sources:

- [Official extension control file](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/kilobase.control)
- [Official upstream documentation](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/README.md)
- [Official project or provider page](https://kbve.com/)

`kilobase` — Preloaded pgrx background worker for scheduling and refreshing PostgreSQL materialized views.

The reviewed catalog snapshot records version `17.4.1`, kind `preload`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "kilobase";
SELECT extversion
FROM pg_extension
WHERE extname = 'kilobase';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
