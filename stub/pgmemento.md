## Usage

Sources:

- [pgMemento v0.7.4 README](https://github.com/pgMemento/pgMemento/blob/v0.7.4/README.md)
- [pgMemento v0.7.4 control file](https://github.com/pgMemento/pgMemento/blob/v0.7.4/extension/pgmemento.control)
- [pgMemento wiki](https://github.com/pgMemento/pgMemento/wiki)
- [Changes from v0.7.3 to v0.7.4](https://github.com/pgMemento/pgMemento/compare/v0.7.3...v0.7.4)

pgmemento is a trigger-based audit trail for PostgreSQL. It records DML changes as JSONB deltas, groups them by transaction and table event, tracks audited schema changes, and provides restore and revert helpers. Use it when row history and transaction context must be queried inside PostgreSQL.

### Create and Initialize

    CREATE EXTENSION pgmemento;
    SELECT pgmemento.init('public');

init instruments eligible tables in the selected schema and adds the pgmemento_audit_id identity column used to track row history. Run it first on a staging copy: auditing changes table definitions and installs event and row triggers.

Use start and stop to control auditing for a schema, and use the documented drop function when intentionally removing pgMemento's instrumentation. Do not manually delete extension triggers or audit identifiers.

### Inspect the Audit Trail

The central data model includes:

- transaction_log: transaction metadata and optional application context.
- table_event_log: table-level events within a transaction.
- row_log: JSONB row deltas linked to a table event.
- audited_schema and audited_table metadata: tracked schemas, tables, columns, and versions.

A typical investigation joins a transaction to its table events and row deltas:

    SELECT t.id,
           t.txid_time,
           e.table_operation,
           r.audit_id,
           r.old_data,
           r.new_data
    FROM pgmemento.transaction_log AS t
    JOIN pgmemento.table_event_log AS e
      ON e.transaction_id = t.id
    JOIN pgmemento.row_log AS r
      ON r.event_key = e.event_key
    WHERE t.id = 12345;

Inspect the installed views and column names before embedding this query because audit schema details can differ across pgmemento versions.

### Restore and Revert

pgmemento provides restore functions that reconstruct table state from the audit trail and revert_transaction or related helpers that apply compensating changes. Treat these as recovery operations:

1. take and verify a backup;
2. identify the exact transaction and dependent changes;
3. preview reconstructed data where possible;
4. run the operation in a controlled transaction;
5. validate constraints, sequences, and application invariants.

### Version 0.7.4

Version 0.7.4 changes row serialization to avoid PostgreSQL's jsonb_build_object argument limit for very wide payloads and adds PostgreSQL 15 support. Upgrade using ALTER EXTENSION pgmemento UPDATE only after testing the version-specific upgrade script.

### Operational Boundaries

- Audit triggers add latency and write volume to every tracked change. Monitor row_log growth and index maintenance.
- The audit trail resides in the same database and is not a substitute for backups, WAL archives, or tamper-resistant external audit storage.
- Schema initialization and DDL tracking alter application tables. Coordinate migrations with pgmemento rather than bypassing its event triggers.
- Limit direct writes to the pgmemento schema and protect any transaction metadata that can contain user or application information.
