## Usage

Sources:

- [Official extension control file](https://github.com/MorphingDB/MorphingDB/blob/f37774fe633f1e8f228c93d61956c80a768545f8/pgdl.control)

`pgdl` — PostgreSQL extension for deep learning model inference and vector storage inside the database.

The reviewed catalog snapshot records version `1.3.0`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgdl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgdl';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
