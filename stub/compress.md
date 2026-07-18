## Usage

Sources:

- [Official upstream documentation](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/README.md)

`compress` — C/C++ functions to compress text into Base64-encoded ZIP archives and extract text from them.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "compress";
SELECT extversion
FROM pg_extension
WHERE extname = 'compress';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
