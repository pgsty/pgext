## Usage

Sources:

- [Official extension control file](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/text_domains.control)

`text_domains` — Pure-SQL collection of constrained PostgreSQL text domains for nonempty, alphabetic, alphanumeric, case-specific, and fixed-length values.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "text_domains";
SELECT extversion
FROM pg_extension
WHERE extname = 'text_domains';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
