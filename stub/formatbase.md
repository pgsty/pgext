## Usage

Sources:

- [Official extension control file](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase.control)

`formatbase` — C functions to format numbers and parse strings by numeric base.

The reviewed catalog snapshot records version `2.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "formatbase";
SELECT extversion
FROM pg_extension
WHERE extname = 'formatbase';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
