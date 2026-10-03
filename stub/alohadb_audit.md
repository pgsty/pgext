## Usage

Sources:

- [contrib/alohadb_audit/alohadb_audit.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit.control)
- [contrib/alohadb_audit/alohadb_audit.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit.c)
- [contrib/alohadb_audit/alohadb_audit--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit--1.0.sql)
- [contrib/alohadb_audit/alohadb_audit--1.0--1.1.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit--1.0--1.1.sql)

`alohadb_audit` 1.1 records DML and DDL activity in the pinned AlohaDB source. It is a kernel-associated extension; these notes do not claim compatibility with stock PostgreSQL.

### Enablement

A superuser must preload the library and restart the server. Merge it with any existing preload list, then create the extension in the fixed `public` schema.

```ini
shared_preload_libraries = 'alohadb_audit'
alohadb.audit_enabled = on
alohadb.audit_log_format = 'json'
```

```sql
CREATE EXTENSION alohadb_audit;
SELECT * FROM audit_log_status();
```

### Configuration and Operations

`alohadb.audit_log_directory` selects a server-writable log directory. `alohadb.audit_databases` and `alohadb.audit_operations` filter events; `alohadb.audit_log_query_text` controls SQL text capture. `audit_log_status()` reports the active configuration. Optional encryption uses OpenSSL AES-256-GCM and a 64-hex-character `alohadb.audit_encryption_key`. `audit_decrypt_log(line, key, redact)` decrypts a supplied line and optionally redacts password patterns.

Protect log files, key configuration and query text; passing a key as a SQL literal can itself enter logs. Redaction is pattern-based, not a guarantee that all secrets are removed. Arrange log retention externally and verify hooks and encryption in the exact AlohaDB build before relying on the audit trail.
