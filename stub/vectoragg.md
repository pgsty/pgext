## Usage

Sources:

- [Official extension control file](https://github.com/cybertec-postgresql/vectoragg/blob/ffbd58f1c355a1302df9ec1b5a65caad812df5df/vectoragg.control)

`vectoragg` — Float-array range sum, clamp, and decimation functions

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "vectoragg";
SELECT extversion
FROM pg_extension
WHERE extname = 'vectoragg';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
