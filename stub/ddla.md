## Usage

Sources:

- [README at the reviewed commit](https://github.com/asotolongo/ddla/blob/385eeaafad77ae401a64a1f60181869ca2d1b973/README)
- [Extension control file](https://github.com/asotolongo/ddla/blob/385eeaafad77ae401a64a1f60181869ca2d1b973/ddla.control)
- [Version 0.1 installation SQL](https://github.com/asotolongo/ddla/blob/385eeaafad77ae401a64a1f60181869ca2d1b973/ddla--0.1.sql)

`ddla` 0.1 is a pure-SQL DDL audit extension. Installation creates the `ddla` schema, the `ddla.ddl_logs` table, and two database-wide event triggers. The triggers capture completed DDL commands and dropped objects together with the command tag, object identity and type, timestamp, current user, client address, and full current query text.

### Inspect the audit log

```sql
CREATE EXTENSION ddla;

CREATE TABLE ddl_audit_example (id bigint PRIMARY KEY);

SELECT * FROM ddla.get_ddl_cmd_by_cmd('CREATE TABLE');
SELECT * FROM ddla.get_ddl_cmd_stats();
```

The query helpers also filter by object type, timestamp range, or user. `ddla.reset_logs()` truncates the audit table, while `ddla.reset_id_seq_logs()` also restarts its sequence.

### Operational cautions

The control file requires superuser installation, and the event triggers run for DDL throughout the database. Plan retention for `ddla.ddl_logs`, restrict access because `query` may contain sensitive text, and measure overhead on DDL-heavy systems. The upstream README explicitly says version 0.1 has known bugs; validate behavior on the exact PostgreSQL release before relying on it as an audit control.
