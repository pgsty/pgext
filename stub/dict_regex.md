## Usage

Sources:

- [Official extension control file](https://github.com/obartunov/dict_regex/blob/a8b8d3270c3d50d59eb8064efd1154128b39c548/dict_regex.control)
- [Official upstream documentation](https://github.com/obartunov/dict_regex/blob/a8b8d3270c3d50d59eb8064efd1154128b39c548/README.dict_regex)

`dict_regex` — Full-text-search dictionary with regular-expression rules.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dict_regex";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_regex';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
