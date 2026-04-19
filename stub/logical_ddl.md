## Usage

Source: [README](https://github.com/samedyildirim/logical_ddl/blob/master/README.md)

`logical_ddl` captures a limited set of `ALTER TABLE` changes, writes them into a replicated shadow table, and replays equivalent DDL on logical-replication subscribers.

### Supported DDL

- `ALTER TABLE ... RENAME TO ...`
- `ALTER TABLE ... RENAME COLUMN ... TO ...`
- `ALTER TABLE ... ADD COLUMN ...`
- `ALTER TABLE ... ALTER COLUMN ... TYPE ...`
- `ALTER TABLE ... DROP COLUMN ...`

Built-in types, arrays, composite types, domains, and enums are supported as column types, but the extension does not replicate the definitions of those custom types themselves.

### Publisher and subscriber setup

```sql
CREATE EXTENSION logical_ddl;

-- Publisher
INSERT INTO logical_ddl.settings VALUES (true, 'source1');
INSERT INTO logical_ddl.publish_tablelist (relid) VALUES ('table1'::regclass);
ALTER PUBLICATION log_pub_1 ADD TABLE logical_ddl.shadow_table;

-- Subscriber
INSERT INTO logical_ddl.settings VALUES (false, 'source1');
INSERT INTO logical_ddl.subscribe_tablelist (source, relid)
VALUES ('source1', 'table1'::regclass);
ALTER SUBSCRIPTION log_sub_1 REFRESH PUBLICATION;
```

### Main tables

- `logical_ddl.settings`: declares publisher/subscriber role and source name.
- `logical_ddl.publish_tablelist`: tables and command kinds to capture.
- `logical_ddl.subscribe_tablelist`: target tables and command kinds to replay.
- `logical_ddl.shadow_table`: replicated command log.
- `logical_ddl.applied_commands`: replay history and failure tracking.

### Caveats

- The extension works under superuser privileges.
- `USING` expressions for type changes, default expressions, constraints, and indexes are not implemented.
- Pigsty notes an upstream `RAISE WARNING` typo fix for `0.1.0`; that does not change the user-facing usage documented here.
