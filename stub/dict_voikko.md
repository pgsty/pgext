## Usage

Sources:

- [Official extension control file](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko.control)
- [Official upstream documentation](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/README.md)

`dict_voikko` — Finnish full-text-search dictionary using Voikko morphological analysis.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dict_voikko";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_voikko';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
