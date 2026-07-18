## Usage

Sources:

- [Official upstream documentation](https://github.com/za-arthur/dict_translate/blob/c95381026d760bdfae6a8cb2bab94c7ca9c06575/README.md)

`dict_translate` — Text search dictionary template that normalizes lexemes and expands them to configured translations.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dict_translate";
SELECT extversion
FROM pg_extension
WHERE extname = 'dict_translate';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
