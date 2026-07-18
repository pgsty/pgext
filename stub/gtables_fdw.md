## Usage

Sources:

- [Extension control file](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/gtables_fdw.control)
- [Version 0.1 installation SQL](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/gtables_fdw--0.1.sql)
- [Native module](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/gtables_fdw.c)
- [Only integration test](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/t/001_basic.pl)

`gtables_fdw` 0.1 is an unfinished global-table prototype, not a functional foreign data wrapper. Its native module has an empty `_PG_init()` function. The installation SQL only creates metadata tables `gt.GlobalTables` and `gt.GlobalTableSpace`; proposed functions and triggers are commented out.

### Inspect without assuming functionality

If the extension already exists in a disposable test database, its entire effective SQL surface can be inspected with:

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'gtables_fdw';

SELECT * FROM gt.GlobalTables;
SELECT * FROM gt.GlobalTableSpace;
```

Adding rows to these tables does not create, synchronize, route, or replicate a global table. The repository has no working management API, query planner, write path, failure handling, or consistency protocol.

### Destructive installation warning

The 0.1 installation script begins with `DROP SCHEMA IF EXISTS gt CASCADE`, then recreates `gt`. Installing it can therefore delete an unrelated existing `gt` schema and every dependent object. Do not execute `CREATE EXTENSION gtables_fdw` on a database containing valuable objects with that schema name. Review the script and use an isolated disposable database if source archaeology is required.

The sole test preloads `postgres_fdw, gtables_fdw`, but the library performs no preload work and the test does not prove table synchronization. The last source change is from 2020 and no compatibility matrix or release exists. Use maintained logical replication, foreign tables, or a distributed PostgreSQL design for actual global-table requirements.
