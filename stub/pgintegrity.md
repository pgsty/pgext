## Usage

Sources:

- [Version 0.0.1 extension SQL](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/pgintegrity--0.0.1.sql)
- [Trigger implementation](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/pgintegrity.c)
- [Upstream implementation notes](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/doc/README.md)

`pgintegrity` is an incomplete 2018 trigger prototype that attempts to maintain per-row “watermark” values in a separate local database through `dblink`. Its published SQL defines only the trigger function `pg_integrity`; it does not create the required remote database, tables, roles, grants, configuration, or triggers. It is not a deployable integrity or audit solution.

### Safe Inspection

If the historical binary is built in a disposable legacy environment, limit initial work to catalog inspection:

```sql
CREATE EXTENSION pgintegrity;

SELECT p.oid::regprocedure, p.prorettype::regtype, l.lanname
FROM pg_proc AS p
JOIN pg_language AS l ON l.oid = p.prolang
WHERE p.proname = 'pg_integrity';
```

Do not attach the trigger to application tables. The C function assumes it is called as an `AFTER` row trigger but does not first validate the trigger call context, and its external setup is neither installed nor documented sufficiently to reproduce safely.

### Source-Confirmed Hidden Assumptions

The implementation expects all of the following outside the extension contract:

- the `dblink` functions are available, although the control file declares no dependency;
- a custom setting named `pgintegrity.password` contains a clear-text password;
- a local database named `pgintegrity`, a hard-coded role named `umsuser`, and tables named `t_privilege` and `t_watermark` already exist;
- application tables expose row OIDs, a capability removed from modern PostgreSQL;
- the server can open extra database connections during each triggered change.

Connection strings and SQL are built through string concatenation, including role, database, table, password, and row data. This creates correctness, secret-handling, and injection risks. The code also uses a non-cryptographic hash over concatenated text representations without unambiguous field boundaries, so the “watermark” is not a trustworthy tamper-evidence primitive.

### Compatibility and Replacement

Version 0.0.1 is relocatable and declares no preload requirement. The repository includes compiled object/shared-library artifacts but no regression test configuration, setup DDL, upgrade path, compatibility matrix, or operational recovery procedure. Its reliance on row OIDs and old server structs prevents use as written on current PostgreSQL.

For real integrity requirements, first define the threat model. Use supported audit logging, restricted append-only audit tables, cryptographic hashes or signatures with canonical serialization, protected key management, and externally verifiable storage as appropriate. Never place database passwords in user-settable GUCs or compose privileged `dblink` SQL from unquoted identifiers and data.
