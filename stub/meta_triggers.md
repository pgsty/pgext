## Usage

Sources:

- [Official extension control file](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/meta_triggers.control)

`meta_triggers` — Insert, update and delete triggers for the meta system catalog

The reviewed catalog snapshot records version `0.5.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `hstore`, `meta`, `plpgsql`.

```sql
CREATE EXTENSION "meta_triggers";
SELECT extversion
FROM pg_extension
WHERE extname = 'meta_triggers';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
