## Usage

Sources:

- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/README.md)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/ledgerext/ledgerext.control)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/ledgerext/ledgerext--0.1.sql)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/LICENSE)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/docs/installation.md)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/docs/usage.md)
- [Official source](https://github.com/geovedi/ledgerext/blob/34c13c0ce3fbb6b1095aac5823fd578d8d994632/docs/operations.md)

`ledgerext` 0.1 is an experimental double-entry ledger for PostgreSQL 18. It keeps account and transfer state in shared memory and persists its own journal outside PostgreSQL WAL. It is not an ordinary SQL-table ledger: backup, replication, transaction and upgrade procedures must account for that independent storage.

### Enable and Inspect

Preload the library and restart before creating the extension as an administrator. The default memory grid is approximately 1192 MB; configure capacity before startup and check the actual host budget. Explicit huge pages are optional.

```ini
shared_preload_libraries = 'ledgerext'
```

```sql
CREATE EXTENSION ledgerext;
SELECT ledger_capacity();
SELECT ledger_journal_count();
SELECT ledger_expiry_due();
```

An administrator binds application roles to book numbers in `ledger_tenants` and grants only the required functions. Installation revokes PUBLIC execution on extension functions. The tenant mapping controls the ledger's access boundary; a client cannot supply a tenant in its batch. Several administrative journal/capacity functions expose cluster-wide information and should stay restricted.

### Accounts, Batches and Statuses

`ledger_create_account` creates a single account; `ledger_create_accounts` and `ledger_apply` accept packed BYTEA batches. Account and transfer codes and ledger numbers must be nonzero. UUIDs are client-provided identities, and transfer IDs support idempotent retries. Amounts are unsigned 128-bit integer units; `ledger_amount` encodes a NUMERIC amount for the documented wire format.

Batch calls queue work until pre-commit. Inspect `ledger_last_statuses` after commit and decode results with `ledger_status_name` or `ledger_account_status_name`; a returned count is not proof that every event succeeded. `ledger_balance`, `ledger_lookup_transfer` and journal query functions inspect state/history. Linked, pending, post and void flags have distinct meanings; use the exact batch layout and examples in the official usage guide.

### Durability and Maintenance

`$PGDATA/ledgerext.journal` is the durable ledger; `$PGDATA/ledgerext.snap` accelerates replay. An SQL dump or PostgreSQL WAL archive alone does not capture this state. Follow the upstream consistent journal/snapshot backup procedure and rehearse restore. Losing the journal can start a new empty ledger. The append-only journal grows without automatic pruning; monitor disk space and fixed account/transfer limits.

Replication ships journal records separately and is not a quorum protocol. The upstream design does not fence an old primary; promotion requires an external fencing procedure. Do not infer PostgreSQL PITR, logical replication, ordinary table MVCC or cross-storage atomicity from the SQL API. Capacity exhaustion returns explicit statuses, and the working set cannot spill to disk.

The release supplies no upgrade scripts; incompatible journal/snapshot versions are refused. Do not replace a deployed library without a reviewed migration/recovery plan.
