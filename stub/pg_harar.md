## Usage

Sources:

- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/README.md)
- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/pg_harar.control)
- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/Cargo.toml)
- [Official source](https://github.com/addiscorp/pg_harar/blob/b9f7a466d43ae65f9419013187dd1c0ff6490081/sql/audit.sql)

`pg_harar` 0.1.0 provides an append-only hash-chained audit table for Harar on PostgreSQL 17. Token identity and its integration into row security are planned work, not part of this audit implementation. The extension does not decrypt secrets or retain a decryption key.

### Core Workflow

A superuser installs the untrusted extension in its fixed `harar` schema. No preload or server restart is needed for this SQL workflow.

```sql
CREATE EXTENSION pg_harar;
INSERT INTO harar.audit (txn, org, actor, zone, via, action, resource, decision, by_role, policy_version)
VALUES ('txn-1', 'example', 'service.api', 'prod', 'broker',
        'read secret', 'DATABASE_URL', 'allow', 'reader', 1);
SELECT * FROM harar.verify_chain();
```

The extension assigns `seq`, `prev_hash`, `hash` and database recording time `at`. The optional `occurred_at` is a writer-supplied event time and is included in the hash. `harar.verify_chain` returns the checked count and first broken sequence; a NULL `first_broken` means the chain and head agree.

### Privileges and Concurrency

Writers need schema USAGE and table INSERT, usually SELECT as well. Keep the internal head row restricted to its owner. UPDATE, DELETE and TRUNCATE are rejected; the audit triggers remain active with replica session behavior.

Inserts serialize on one head row, held until transaction end. Use short transactions and an idle-transaction timeout. Under repeatable-read or serializable isolation, handle serialization failures by retrying the transaction. Rolled-back inserts leave no audit row.

### Integrity and Recovery

An owner or superuser can disable triggers and rewrite both history and head. Signed checkpoints stored outside the database machine are needed to detect that class of tampering; a valid local chain alone does not prove authenticity.

`pg_dump` includes the audit rows but not the head. The upstream restore procedure disables triggers while restoring, then has the owner reset the head to the final stored sequence/hash with the head's triggers disabled, and verifies the chain. Rehearse the complete restore procedure before relying on a backup.
