## Usage

Sources:

- [Official upstream documentation](https://github.com/cipherstash/encrypt-query-language/blob/4119c4246427c54ffb29da0f70a6935a1c7b5215/README.md)
- [Official project or provider page](https://cipherstash.com/docs)
- [Official primary documentation](https://database.dev/cipherstash/eql)

`eql` — Encrypt Query Language provides PostgreSQL types, functions, operators, and index support for searchable encrypted data.

The reviewed catalog snapshot records version `2.2.1`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "eql";
SELECT extversion
FROM pg_extension
WHERE extname = 'eql';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
