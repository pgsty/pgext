## Usage

Sources:

- [Official README](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/README.md)
- [Official extension SQL](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/sql/table_log_pl--0.1.sql)
- [Official extension control file](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/table_log_pl.control)

`table_log_pl` is a PL/pgSQL replacement for the older C `table_log` extension. It records row images from INSERT, UPDATE, and DELETE operations into a shadow table and can reconstruct a separate table as of a timestamp; it is a compatibility aid, not a tamper-resistant audit system.

### Core Workflow

Create the destination schema, then initialize logging for a table:

```sql
CREATE EXTENSION table_log_pl;

CREATE TABLE public.account (
  id bigint PRIMARY KEY,
  balance numeric
);
CREATE SCHEMA audit;

SELECT table_log_pl_init(
  5, 'public', 'account', 'audit', 'account_log'
);
```

`table_log_pl_init(integer, text, text, text, text)` creates a log table and an AFTER INSERT/UPDATE/DELETE row trigger. Level 3 records row metadata without an event ID or user; level 4 adds `trigger_id`; level 5 adds both `trigger_id` and `trigger_user`.

### Log and Restore Semantics

The shadow table copies the original columns and adds `trigger_mode`, `trigger_tuple`, `trigger_changed`, and the optional level-dependent fields. INSERT records the new row, DELETE records the old row, and UPDATE records both old and new row images.

`table_log_pl_restore_table(...)` reconstructs a separate table at a timestamp. Method 0 replays complete history forward, while method 1 works backward from the current source table. The default result is temporary, and the function does not overwrite the original table.

Logging requires exact row-shape compatibility; later column changes can break the trigger or restore logic. TRUNCATE and DDL are not captured. The trigger adds synchronous write cost, and owners able to change the trigger or log table can alter history. Protect the audit schema, test schema migrations and restore procedures, and use a purpose-built audit design when evidentiary integrity is required.
