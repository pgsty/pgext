## Usage

Sources:

- [README.md](https://github.com/GerardSmit/pg_temporal_tables/blob/bd849266516293ca7e3b41225fdc61284ef66843/README.md)
- [sql/pg_temporal_tables--0.1.1.sql](https://github.com/GerardSmit/pg_temporal_tables/blob/bd849266516293ca7e3b41225fdc61284ef66843/sql/pg_temporal_tables--0.1.1.sql)
- [pg_temporal_tables.control](https://github.com/GerardSmit/pg_temporal_tables/blob/bd849266516293ca7e3b41225fdc61284ef66843/pg_temporal_tables.control)

`pg_temporal_tables` records row history and rewrites queries for a selected historical time. The current implementation targets PostgreSQL 18 and requires preload and a server restart.

### Core Workflow

```ini
shared_preload_libraries = 'pg_temporal_tables'
```

```sql
CREATE EXTENSION pg_temporal_tables;
CREATE TABLE users (id bigint PRIMARY KEY, name text);
SELECT temporal.enable('public.users');
SET temporal.user_id = 'operator';
INSERT INTO users VALUES (1, 'Alice');
UPDATE users SET name = 'Alicia' WHERE id = 1;
```

### Operational Boundaries

A tracked table needs a primary key. `temporal.enable()` creates history tracking; `temporal.disable()` stops it. `temporal.changes()` and `temporal.column_history()` inspect changes; retention and compaction functions maintain history. Set `temporal.user_id` for attribution and reset `temporal.as_of` after historical reads. Track the root of a partitioned table; legacy inheritance and TimescaleDB hypertables are unsupported. Do not change or reuse primary keys without reviewing the lineage option. Installation requires a superuser. Back up the history together with the base data and review retention before deleting history.
