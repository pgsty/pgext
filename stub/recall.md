## Usage

Sources:

- [Official README](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/README.md)
- [Version 0.9.6 extension SQL](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/recall--0.9.6.sql)
- [Official worked example](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/examples/blog.md)

recall keeps full-row temporal history for selected PostgreSQL tables. A trigger copies each version to a per-table log with a timestamp range, allowing point-in-time queries within a configured retention interval. It is a user-data safety net, not a backup or tamper-resistant audit log.

### Core Workflow

The extension requires `btree_gist`, and the target table must have a primary key. Enable history, modify data normally, and ask recall to create a session-local past view:

```sql
CREATE EXTENSION recall CASCADE;

CREATE TABLE account_setting (
  account_id bigint PRIMARY KEY,
  value text NOT NULL
);

SELECT recall.enable('account_setting', '6 months');

INSERT INTO account_setting VALUES (1, 'old');
UPDATE account_setting SET value = 'new' WHERE account_id = 1;

SELECT recall.at('account_setting', now() - interval '1 minute');
SELECT * FROM account_setting_past WHERE account_id = 1;
```

The log can also be queried directly with `_log_time @> timestamp`.

### Important Objects

- `recall.enable` creates the template and log tables, copies existing rows, and installs the history trigger.
- `recall.at` creates a temporary view named with the `_past` suffix for one timestamp.
- `recall.cleanup` and `recall.cleanup_all` delete versions older than the configured interval.
- `recall.disable` removes history machinery and permanently drops that table's retained log.
- `_log_time` is a `tstzrange`; a GiST exclusion constraint prevents overlapping versions for the same primary key.

### Operational Notes

Each change stores a full row, increasing storage, WAL, and write cost. Cleanup is manual. Schema changes must be applied through the generated template table, while log tables do not inherit ordinary constraints or foreign keys. The log is directly mutable by privileged users and is not evidence against tampering. Test primary-key changes and modern PostgreSQL compatibility carefully; always keep independent backups.
