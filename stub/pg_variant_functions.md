## Usage

Sources:

- [Official extension control file](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/pg_variant_functions.control)
- [Official upstream documentation](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/README.md)

`pg_variant_functions` — Variant summary and allele-frequency functions over TINYINT[] genotype arrays

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `tinyint`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_variant_functions";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_variant_functions';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
