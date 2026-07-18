## Usage

Sources:

- [Official extension control file](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/ext/gitgres.control)
- [Official upstream documentation](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/README.md)

`gitgres` — Store Git objects and refs in PostgreSQL tables with SQL functions and a native git_oid type.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pgcrypto`.

```sql
CREATE EXTENSION "gitgres";
SELECT extversion
FROM pg_extension
WHERE extname = 'gitgres';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
