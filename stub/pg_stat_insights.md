## Usage

Sources:

- [Official documentation](https://github.com/pgElephant/pg_stat_insights/blob/07d83c9e3c606f4ef1b6b7d571d48c63d4ad57c5/README.md)
- [Extension control file](https://github.com/pgElephant/pg_stat_insights/blob/07d83c9e3c606f4ef1b6b7d571d48c63d4ad57c5/pg_stat_insights.control)
- [Official repository](https://github.com/pgElephant/pg_stat_insights)

`pg_stat_insights` Query planning/execution statistics with replication, index, trend, and contention analysis.

### Enablement

Merge `pg_stat_insights` into the existing preload list, restart PostgreSQL, and then create `pg_stat_insights` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_stat_insights'
```

```sql
CREATE EXTENSION pg_stat_insights;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- ステップ 1：PostgreSQL 設定で拡張機能を有効化
ALTER SYSTEM SET shared_preload_libraries = 'pg_stat_insights';
-- PostgreSQL サーバーの再起動が必要

-- ステップ 2：データベースに拡張機能を作成
CREATE EXTENSION pg_stat_insights;

-- ステップ 3：最も遅いクエリを即座に表示
SELECT
    query,
    calls,
    total_exec_time,
    mean_exec_time,
    rows
FROM pg_stat_insights_top_by_time
LIMIT 10;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_stat_insights_top_by_time` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_by_bucket` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_errors` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_histogram_summary` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_index_by_bucket` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_index_size_by_bucket` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_plan_errors` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_insights_replication_by_bucket` | VIEW | Inspection or query view created by the extension. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_stat_insights` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
