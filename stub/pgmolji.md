## Usage

Sources:

- [Official extension control file](https://github.com/NguyenLe1605/pgmolji/blob/42b9486e29319026df7c957b484eaf47fceabecd/pgmolji.control)

`pgmolji` — pgrx extension that rewrites strings to random emoji strings.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgmolji";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmolji';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
