## Usage

Sources:

- [Official extension control file](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/pg_markdown.control)
- [Official upstream documentation](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/README.md)

`pg_markdown` — Pre-release Markdown-to-HTML conversion extension for PostgreSQL.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_markdown";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_markdown';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
