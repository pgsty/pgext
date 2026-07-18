## Usage

Sources:

- [Official extension control file](https://github.com/frost242/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask.control)
- [Official upstream documentation](https://github.com/frost242/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/README.md)

`pg_decode_infomask` — PostgreSQL 9.6 C functions that decode heap tuple infomask and infomask2 hint bits into queryable fields.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_decode_infomask";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_decode_infomask';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
