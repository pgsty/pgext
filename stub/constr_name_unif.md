## Usage

Sources:

- [Official extension control file](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/constr_name_unif.control)
- [Official upstream documentation](https://github.com/katrinaibast/constr_name_unif/blob/e69b3a79fcaa4d013e648b812807a607bd14b7ff/doc/constr_name_unif.md)
- [Official PGXN distribution page](https://pgxn.org/dist/constr_name_unif/)

`constr_name_unif` — Uniforms the names of constraints on base tables.

The reviewed catalog snapshot records version `1.0.1`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "constr_name_unif";
SELECT extversion
FROM pg_extension
WHERE extname = 'constr_name_unif';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
