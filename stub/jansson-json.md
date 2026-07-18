## Usage

Sources:

- [Official upstream documentation](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/README.md)

`jansson-json` — JSON data type, operators, and functions backed by the Jansson JSON library.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "jansson-json";
SELECT extversion
FROM pg_extension
WHERE extname = 'jansson-json';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
