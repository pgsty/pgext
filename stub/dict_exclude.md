## Usage

Sources:

- [Official extension control file](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/dict_exclude.control)
- [Official upstream documentation](https://github.com/no0p/dict_exclude/blob/2058a16ca44942a0c11eae8f7bffa9a733d7cc5a/README.md)

`dict_exclude` — Full-text search dictionary that filters stop words using regular-expression rules.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dict_exclude";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_exclude';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
