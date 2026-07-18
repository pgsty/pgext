## Usage

Sources:

- [Official extension control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/f600a9def937ff5614fbecac805494777d4e69d8/web/web_common/cas_codes/cas_codes.control)

`cas_codes` — CAS Registry Number data type and validation support for PostgreSQL.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cas_codes";
SELECT extversion
FROM pg_extension
WHERE extname = 'cas_codes';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
