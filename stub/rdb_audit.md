## Usage

Sources:

- [docker/raodb/ext/rdb_audit/README.md](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/README.md)
- [docker/raodb/ext/rdb_audit/rdb_audit--1.0.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit--1.0.sql)
- [docker/raodb/ext/rdb_audit/rdb_audit.conf](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit.conf)
- [docker/raodb/ext/rdb_audit/rdb_audit.c](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit.c)
- [docker/raodb/ext/rdb_audit/rdb_audit.control](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_audit/rdb_audit.control)

`rdb_audit` 1.0 is a source audit-logging extension in the RAO DB repository. It records configured statement classes and object information; a complete PostgreSQL compatibility matrix and redistribution license are not established by the reviewed files.

### Core Workflow

```ini
shared_preload_libraries = 'rdb_audit'
rdb_audit.log_scope = 'read,write,ddl,role'
```

```sql
CREATE EXTENSION rdb_audit;
SHOW rdb_audit.log_scope;
```

### Operational Boundaries

Append the library to existing preload entries and restart before creating the extension as a superuser. The SQL installs DDL/drop event triggers and an audit table; library loading alone does not install these objects.

`rdb_audit.log_scope` selects read, write, DDL, role and other classes. `rdb_audit.log_format` selects the output representation, while `rdb_audit.log_directory` controls the log destination. `rdb_audit.log_parameter` includes bind values and `rdb_audit.log_statement` includes SQL text. These may contain secrets or personal data: restrict log access and set retention and disk limits.

Object auditing uses an audit role configured with `rdb_audit.role` and table privileges. Audit output can be much larger than the underlying changes; qualify overhead and completeness in a test deployment before relying on it.
