## Usage

Sources:

- [Official upstream documentation](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/README.md)

`replace_empty_string` — C trigger function that replaces empty strings with NULL

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "replace_empty_string";
SELECT extversion
FROM pg_extension
WHERE extname = 'replace_empty_string';
```

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
