## Usage

Sources:

- [Official extension control file](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/emoji.control)
- [Official upstream documentation](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/README.md)

`emoji` — Pure SQL PostgreSQL extension to encode and decode bytea/text as emoji

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "emoji";
SELECT extversion
FROM pg_extension
WHERE extname = 'emoji';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
