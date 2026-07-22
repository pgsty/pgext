## Usage

Sources:

- [Official README](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/README.md)
- [Version 0.1 SQL implementation](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/tablelog--0.1.sql)
- [Extension control file](https://github.com/snaga/tablelog/blob/0981b7efe69a477c06cce53f0b2151d6b1dd38ad/tablelog.control)

`tablelog` 0.1 is a PL/v8 trigger-based change log for `INSERT`, `UPDATE`, and `DELETE`. It writes changed column names and old/new values as text arrays to the shared `__table_logs__` table, together with transaction, user, table, event, and key information.

### Enable and Inspect Logging

A target table must have a primary key or unique constraint. Install PL/v8 first, enable a table by schema and name, and inspect the central log.

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION tablelog;

SELECT tablelog_enable_logging('public', 'account');

SELECT ts, txid, dbuser, schemaname, tablename,
       event, col_names, old_vals, new_vals, key_names, key_vals
FROM __table_logs__
WHERE schemaname = 'public' AND tablename = 'account'
ORDER BY ts;

SELECT tablelog_disable_logging('public', 'account');
```

The installed trigger is `tablelog_logging_trigger`. New log rows start with `status = 0`; other status meanings are application-defined. UPDATE rows contain changed columns rather than a guaranteed complete row image.

### Audit and Replay Boundaries

Trigger writes participate in the source transaction, so rollback removes both the business change and its log. Owners or privileged users can disable triggers, edit the log, or bypass it, and `TRUNCATE` is not among the documented events. This is useful change history, not an independent tamper-evident audit trail.

Values are converted to text arrays, which lose native type identity and can be affected by formatting, collation, time-zone, encoding, and extension-type output changes. Do not replay them blindly as SQL; preserve schema/version context and validate quoting, NULL, binary, generated, identity, and large-value behavior. PL/v8 executes JavaScript inside each writing backend, and a single central table can become a write, storage, vacuum, and retention bottleneck. Restrict log access, exclude secrets where possible, benchmark workload overhead, and assign explicit archival and deletion policies.
