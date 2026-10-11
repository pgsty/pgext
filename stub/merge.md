## Usage

Sources:

- [docs/merge.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/merge.rst)
- [merge/merge--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/merge/merge--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [merge/merge.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/merge/merge.control)

`merge` is Oliphant’s table-transfer toolkit. It constructs SQL by matching columns by name and uses keys to reconcile a source relation with a destination.

### Core Workflow

```sql
CREATE EXTENSION "merge";
CREATE TABLE merge_source (id integer PRIMARY KEY, label text);
CREATE TABLE merge_target (id integer PRIMARY KEY, label text);
INSERT INTO merge_source VALUES (1, 'example');
SELECT auto_insert('merge_source', 'merge_target');
SELECT * FROM merge_target;
```

### Operational Boundaries

These are legacy pure-SQL/PLpgSQL utilities originally documented for PostgreSQL 9.x. Current-major compatibility is unverified. Installation uses the control defaults and needs administrator privileges; no shared library or preload is required. `auto_insert` copies shared columns, leaving destination-only columns to their defaults. `auto_merge` uses a writable CTE to update matching keys and insert missing rows; it is not a concurrency-safe replacement for native conflict handling. `auto_delete` removes destination rows absent from the source and needs deliberate scope. The caller needs read and write privileges. Validate duplicates, keys, constraints and the source selection before a bulk operation.
