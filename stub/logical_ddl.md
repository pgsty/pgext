## Usage
- GitHub Repo: [`samedyildirim/logical_ddl`](https://github.com/samedyildirim/logical_ddl)
- README: [samedyildirim/logical_ddl/blob/master/README.md](https://github.com/samedyildirim/logical_ddl/blob/master/README.md)

`logical_ddl` captures supported DDL changes on tables and replays them through built-in logical replication. The project is aimed at reducing manual DDL operations and avoiding schema drift between publisher and subscriber.

The README states PostgreSQL 11 and above are supported, and the extension works under superuser privileges.

### What It Captures

Supported DDL operations are:

- `ALTER TABLE ... RENAME TO ...`
- `ALTER TABLE ... RENAME COLUMN ... TO ...`
- `ALTER TABLE ... ADD COLUMN ...`
- `ALTER TABLE ... ALTER COLUMN ... TYPE ...`
- `ALTER TABLE ... DROP COLUMN ...`

The README also notes that built-in types, arrays, composite types, domains, and enums are supported, but definitions of composite types, domains, and enums themselves are not replicated.

### Publisher Side

```sql
CREATE EXTENSION logical_ddl;
INSERT INTO logical_ddl.settings VALUES (true, 'source1');
INSERT INTO logical_ddl.publish_tablelist (relid) VALUES ('table1'::regclass);
ALTER PUBLICATION log_pub_1 ADD TABLE logical_ddl.shadow_table;
```

The extension captures DDL with an event trigger, stores it in `logical_ddl.shadow_table`, and publishes that table through logical replication.

### Subscriber Side

```sql
CREATE EXTENSION logical_ddl;
INSERT INTO logical_ddl.settings VALUES (false, 'source1');
INSERT INTO logical_ddl.subscribe_tablelist (source, relid) VALUES ('source1', 'table1'::regclass);
ALTER SUBSCRIPTION log_sub_1 REFRESH PUBLICATION;
```

On the subscriber, the extension listens for incoming DDL records and replays the generated SQL against the target table.

### Data Model

- `logical_ddl.settings` controls whether the node acts as publisher, subscriber, or both.
- `logical_ddl.publish_tablelist` controls which tables and command kinds are captured.
- `logical_ddl.subscribe_tablelist` controls which tables receive replayed DDL.
- `logical_ddl.shadow_table` stores captured commands.
- `logical_ddl.applied_commands` tracks generated SQL and whether execution failed.

### Caveats

- DDL support is limited to the operations listed above.
- Default expressions, constraints, indexes, and `USING` expressions are not implemented.
- The README notes the project is still under development and may change incompatibly.

### Scope

The upstream README was enough to cover the extension model, captured DDL types, publisher/subscriber setup, and the known caveats. No extra docs page or homepage was needed for the stub.
