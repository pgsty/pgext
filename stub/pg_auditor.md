


## Usage

> [pg_auditor: Audit data changes with flashback capability](https://github.com/kouber/pg_auditor)

`pg_auditor` records each data modification (INSERT, UPDATE, DELETE) on specified tables and allows partial or complete flashback of transactions.

```sql
CREATE EXTENSION pg_auditor CASCADE;  -- also installs hstore
```

### Auditing Control

```sql
-- Start auditing a table (all DML by default)
SELECT auditor.attach('fruit');

-- Audit specific operations only
SELECT auditor.attach('fruit', ARRAY['INSERT', 'UPDATE']);

-- Audit specific columns only
SELECT auditor.attach('fruit', ARRAY['INSERT', 'UPDATE', 'DELETE'], ARRAY['name', 'weight']);

-- Stop auditing
SELECT auditor.detach('fruit');

-- Manage individual statements/columns
SELECT auditor.attach_statement('fruit', 'DELETE');
SELECT auditor.detach_statement('fruit', 'DELETE');
SELECT auditor.attach_column('fruit', 'weight');
SELECT auditor.detach_column('fruit', 'weight');

-- Protect against TRUNCATE
SELECT auditor.forbid_truncate('fruit');
```

### Viewing Audit Log

```sql
SELECT transaction_id, operation, old_rec, new_rec FROM auditor.log;
```

### Flashback Functions

```sql
-- Undo the last N transactions in current session
SELECT auditor.undo();          -- undo last 1
SELECT auditor.undo(3);         -- undo last 3
SELECT auditor.undo(1, true);   -- override other sessions

-- Cancel a specific transaction
SELECT auditor.cancel(5557);

-- Restore data to a specific transaction or timestamp
SELECT auditor.flashback(5556);
SELECT auditor.flashback('2021-02-08 14:40:00'::timestamp);
```

### Column Evolution Tracking

```sql
SELECT * FROM auditor.evolution('fruit'::regclass, 'weight', 'orange'::text);
-- Shows complete history of a column value for a given primary key
```
