## Usage

Sources:

- [Official control file](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/audit.control)
- [Official version 1.0.0 SQL](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/audit--1.0.0.sql)
- [Official repository](https://github.com/ttfkam/pg_audit)

`audit` is a small SQL-only row-change audit prototype. It provides an `audit` table and a trigger function intended to capture the old or new row as `jsonb` after `INSERT`, `UPDATE`, and `DELETE`. Version `1.0.0` is an old, unfinished artifact; inspect and patch its SQL before relying on it.

### Core Workflow

The intended workflow is to install the extension, then attach its `audit()` trigger function to every table that needs row-level history:

```sql
CREATE EXTENSION audit;

CREATE TRIGGER accounts_audit
AFTER INSERT OR UPDATE OR DELETE ON accounts
FOR EACH ROW EXECUTE FUNCTION audit();

SELECT id, relation::regclass, op, row_data, archived
FROM audit
ORDER BY id DESC;
```

For inserts and updates the trigger is intended to store `NEW`; for deletes it stores `OLD`. The `archived` flag is application-managed and is not an automatic retention policy.

### Objects

- `audit`: audit-row table with a `jsonb` payload, relation OID, operation enum, and archive flag.
- `operation`: enum with `INSERT`, `UPDATE`, and `DELETE` values.
- `audit()`: row-trigger function.
- `audit_id(jsonb)`: extracts an `id` property as `bigint`.
- `audit_oid(name, name)`: intended to resolve a schema and relation name to an OID.
- `oid_metadata`: view over relation and schema metadata.

### Caveats

The published `1.0.0` SQL is internally inconsistent: the table column is named `op`, while the trigger inserts into `operation`; `audit_oid` also compares parameter names that shadow its selected columns. The script's final `\echo` names a different extension. As published, the trigger example therefore requires an upstream SQL review and corrective patch before it can work reliably.

The trigger stores complete row images and has no pruning, tamper protection, actor/session metadata, or DDL auditing. Restrict access to the audit table, define retention explicitly, and test storage and write-amplification costs on a disposable database before considering real use.
