## Usage

Sources:

- [Official extension control file](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/dict_roman.control)
- [Official upstream documentation](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/README.dict_roman)

`dict_roman` — Full-text search dictionary that converts Roman numeral tokens to Arabic integers.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dict_roman";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_roman';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
