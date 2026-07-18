## Usage

Sources:

- [Official extension control file](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/schedule.control)

`schedule` — Cron-formatted schedule type and UTC timestamp-matching functions for PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "schedule";
SELECT extversion
FROM pg_extension
WHERE extname = 'schedule';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
