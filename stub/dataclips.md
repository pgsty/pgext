## Usage

Sources:

- [Official extension control file](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/dataclips.control)
- [Official upstream documentation](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/docs/dataclip.md)
- [Official upstream README](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/README.md)

`dataclips` — Archived, unfinished C prototype for querying Heroku Dataclips through a function and planned FDW.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "dataclips";
SELECT extversion
FROM pg_extension
WHERE extname = 'dataclips';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
