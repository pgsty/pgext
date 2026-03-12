

## Usage

> [pg_relusage: log relations actually used by SQL statements](https://github.com/adept/pg_relusage)

pg_relusage hooks into the query executor and logs which relations (tables, views, indexes, etc.) are actually used by each SQL statement. Unlike the statement log, it reports relations after view expansion and unused join elimination.

### How It Works

Once loaded, each SQL statement emits a log message listing all referenced relations:

```sql
SELECT * FROM pg_stats LIMIT 1;
```

Produces log output:
```
relations used: pg_stats,pg_statistic,pg_class,pg_attribute,pg_namespace
```

### Loading

```sql
-- Per-session
LOAD 'pg_relusage';

-- Or globally in postgresql.conf
shared_preload_libraries = 'pg_relusage'
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_relusage.log_level` | `LOG` | Log level for relation messages |
| `pg_relusage.rel_kinds` | `'riSvmfp'` | Relation kinds to report (one-letter codes from `pg_class.relkind`) |

Relation kind codes: `r` = table, `i` = index, `S` = sequence, `v` = view, `m` = materialized view, `f` = foreign table, `p` = partitioned table.

### Use Case

This extension is useful for discovering unused objects in legacy databases. By analyzing the logged relation usage over time, you can identify which tables, views, and indexes are actually accessed by your application.
