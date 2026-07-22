## Usage

Sources:

- [ProcessUtility hook implementation](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.c)
- [Extension SQL](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate--1.0.sql)
- [Official preload configuration](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.conf)
- [Extension control file](https://github.com/pgexperts/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.control)

`pg_reject_truncate` installs a utility-command hook that rejects every `TRUNCATE` statement with `TRUNCATE IS NOT ALLOWED`. It is a coarse, server-wide guard with no per-table, per-role, or administrative bypass; it is not a complete data-loss prevention policy.

### Setup and Behavior

Preload the library in every server process and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_reject_truncate'
```

The version `1.0` extension SQL contains no callable SQL objects. `CREATE EXTENSION pg_reject_truncate` can record the extension in the catalog, but does not itself load the hook because the SQL script never references the shared library. Preloading is therefore the effective activation step. For an isolated diagnostic session, `LOAD 'pg_reject_truncate'` activates the hook only in that backend.

Once loaded, any direct `TRUNCATE`, including one issued by a superuser, is rejected. The implementation provides no GUC or function for temporarily disabling or scoping the rule.

### Boundaries and Hook Compatibility

The hook checks only utility statements whose node type is `T_TruncateStmt`. It does not stop `DELETE`, `DROP TABLE`, partition operations, table rewrites, filesystem loss, or destructive actions by other mechanisms. Use privileges, backups, auditing, and tested recovery procedures as the primary controls.

The implementation saves the previous `ProcessUtility_hook` when it loads, but calls `standard_ProcessUtility` directly for every permitted command instead of invoking the saved hook. As a result, a utility hook loaded before `pg_reject_truncate` can be bypassed for non-`TRUNCATE` commands. Load order with auditing, security, DDL replication, or other utility-hook extensions is therefore unsafe without a source-level fix and integration test.

The reviewed source was last updated in 2020 and uses version-specific PostgreSQL hook signatures. Build and test it for the exact server major, verify hook ordering after restart, and run both positive and negative command tests before relying on the guard.
