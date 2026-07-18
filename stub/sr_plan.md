## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/sr_plan.control)
- [Official upstream documentation](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/README.md)

`sr_plan` — Planner-hook extension that records, displays, and reuses saved execution plans, including parameterized-plan templates.

The reviewed catalog snapshot records version `1.2`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "sr_plan";
SELECT extversion
FROM pg_extension
WHERE extname = 'sr_plan';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
