## Usage

Sources:

- [examples/pgaudit_zig/README.md](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/README.md)
- [examples/pgaudit_zig/build.zig](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/build.zig)
- [examples/pgaudit_zig/src/main.zig](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/src/main.zig)
- [examples/pgaudit_zig/extension/pgaudit_zig--0.1.sql](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/extension/pgaudit_zig--0.1.sql)
- [examples/pgaudit_zig/extension/pgaudit_zig.control](https://github.com/xataio/pgzx/blob/edc51d36a19838e0afd76ab69fb5c2e0279b533c/examples/pgaudit_zig/extension/pgaudit_zig.control)

`pgaudit_zig` 0.1 demonstrates executor audit hooks using `pgzx`. Upstream explicitly describes it as an incomplete teaching example and recommends `pgaudit` for production auditing.

### Core Workflow

```sql
LOAD 'pg_audit_zig';
CREATE EXTENSION pgaudit_zig;
CREATE TABLE audit_demo (id integer);
SELECT * FROM audit_demo;
```

### Operational Boundaries

The SQL extension is named `pgaudit_zig`, while the build names its library `pg_audit_zig`. Load that library explicitly in the session to install hooks; `CREATE EXTENSION` alone runs an otherwise empty SQL script and does not activate the logger. These operations require appropriate superuser privileges. The hooks log command kind, referenced relations and statement text. `pgaudit_zig.log_statement` is a superuser-settable boolean, defaulting to true, for including statement text; logged text can contain sensitive values. The reviewed code sends records through PostgreSQL logging. No complete audit coverage, tamper protection or compatibility range is promised by the example.
