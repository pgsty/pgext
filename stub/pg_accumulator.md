


## Usage

Sources: [README](https://github.com/Treedo/pg_accumulator/blob/main/README.MD), [documentation](https://github.com/Treedo/pg_accumulator/blob/main/doc/pg_accumulator.md), [control file](https://github.com/Treedo/pg_accumulator/blob/main/pg_accumulator.control)

`pg_accumulator` implements declarative accumulation registers for PostgreSQL. A register records movements across arbitrary dimensions and resources, then exposes generated functions for balances, turnovers, movement history, verification, and rebuild operations.

### Enable

```sql
CREATE EXTENSION pg_accumulator;
```

For high-write workloads that use the delta buffer and background worker, preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_accumulator'
```

### Create A Register

```sql
SELECT accum.register_create(
  name := 'inventory',
  dimensions := '{"warehouse": "int", "product": "int"}',
  resources := '{"quantity": "numeric", "amount": "numeric"}',
  kind := 'balance'
);
```

The call creates the register metadata, movement table, total tables, balance cache, triggers, generated query functions, and supporting indexes.

### Post And Query Movements

```sql
SELECT accum.register_post('inventory', '{
  "period": "2026-04-01",
  "document": "receipt:1",
  "dimensions": {"warehouse": 1, "product": 42},
  "resources": {"quantity": 10, "amount": 250}
}');

SELECT * FROM accum.inventory_balance(
  dimensions := '{"warehouse": 1, "product": 42}'
);

SELECT * FROM accum.inventory_turnover(
  from_date := '2026-04-01',
  to_date := '2026-04-30',
  dimensions := '{"warehouse": 1}',
  group_by := '["product"]'
);
```

### Corrections And Maintenance

Use `register_unpost` and `register_repost` for retroactive corrections. Use `register_verify`, `register_rebuild_totals`, and `register_rebuild_cache` to check and repair derived state when needed.

```sql
SELECT accum.register_unpost('inventory', 'receipt:1');
SELECT accum.register_verify('inventory');
```

### Configuration

The documented GUCs include `pg_accumulator.background_workers`, `pg_accumulator.delta_merge_interval`, `pg_accumulator.delta_merge_delay`, `pg_accumulator.delta_merge_batch_size`, `pg_accumulator.partitions_ahead`, and `pg_accumulator.maintenance_interval`.
