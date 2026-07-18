## Usage

Sources:

- [Official extension control file](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/ajversion.control)
- [Official upstream documentation](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/README.md)

`ajversion` — Semantic version type for PostgreSQL backed by compact integer storage.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "ajversion";
SELECT extversion
FROM pg_extension
WHERE extname = 'ajversion';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
