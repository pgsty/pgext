## Usage

Sources:

- [Official extension control file](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/content_utils.control)
- [Official upstream documentation](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/README.md)

`content_utils` — Utility functions for news management and full-text search.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `file_fdw`.

```sql
CREATE EXTENSION "content_utils";
SELECT extversion
FROM pg_extension
WHERE extname = 'content_utils';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
