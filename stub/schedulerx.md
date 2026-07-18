## Usage

Sources:

- [Official extension control file](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/schedulerx.control)
- [Official upstream documentation](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/README.md)

`schedulerx` — Prototype background-worker scheduler for periodic, one-shot, and NOTIFY-triggered SQL jobs.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "schedulerx";
SELECT extversion
FROM pg_extension
WHERE extname = 'schedulerx';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
