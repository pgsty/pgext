## Usage

Sources:

- [Implementation (polar_audit.c)](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/aabf18dcf39abee55de91f84630cfa3143ea7431/external/polar_audit/polar_audit.c)
- [Extension control file](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/aabf18dcf39abee55de91f84630cfa3143ea7431/external/polar_audit/polar_audit.control)
- [Installation SQL](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/aabf18dcf39abee55de91f84630cfa3143ea7431/external/polar_audit/polar_audit--1.0.sql)

`polar_audit` logs statement classes and database objects using PolarDB-specific audit hooks. This record is verified against the upstream PostgreSQL-15-based PolarDB branch, not a competition fork or stock PostgreSQL build.

### Enablement

Add the library to the existing preload list and restart PolarDB. Choose audit classes explicitly; the default logs no classes. The versioned extension SQL contains no SQL objects, so registration alone does not activate auditing.

```conf
shared_preload_libraries = 'polar_audit'
polar_audit.log = 'read,write,ddl'
```

### Settings and Boundaries

`polar_audit.log` selects comma-separated classes; prefix a class with `-` to subtract it. `polar_audit.log_catalog` controls catalog-only activity, `polar_audit.log_relation` emits per-relation entries, and `polar_audit.log_statement` / `polar_audit.log_parameter` control SQL text and parameter logging. `polar_audit.role` selects the audit role. These settings require privileged administration. Logs can expose application data and grow rapidly; apply retention and access control. The source calls PolarDB audit functions and rejects loading outside server preload, so no stock-PostgreSQL support is claimed.
