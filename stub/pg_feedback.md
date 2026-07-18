## Usage

Sources:

- [Pinned upstream README](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/README.md)
- [Version 1.0 installation SQL](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/pg_feedback--1.0.sql)
- [Pinned PostgreSQL integration](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/pg_feedback.c)
- [Pinned operating-system helpers](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/helpers.c)
- [Pinned distribution metadata](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/META.json)

pg_feedback creates a local JSON report about a PostgreSQL installation. The report contains a cluster-system identifier, a value labeled postmaster uptime, data-directory filesystem type, server version, selected settings, available extensions, and total database size. The extension does not transmit the report; the caller decides whether and where to send it.

### Generate and review locally

```sql
CREATE EXTENSION pg_feedback;

SELECT jsonb_pretty(feedback()::jsonb);
```

Review the result as sensitive operational inventory before exporting it. The SQL object version is 1.0, matching this catalog, while the pinned PGXN-style distribution metadata is 0.2.2. These are different version namespaces.

### Data quality and exposure

The implementation contains material correctness defects. It formats the 64-bit cluster identifier into a heap buffer using the size of a pointer, which truncates the identifier on common platforms. The function named postmaster_uptime actually returns the parent process ID. On Linux, the standalone sysinfo helper assigns system uptime to the memory field and memory size to the uptime field; that helper is currently omitted from the main report but remains callable.

The installation SQL does not revoke default PUBLIC execution. Reports expose settings, installed capabilities, aggregate database size, platform details, and a cluster-derived identifier to any role that can call the functions. Revoke execution and grant it only to a monitoring role. Do not upload output without an explicit privacy and destination review.

The C code reads PostgreSQL control-file internals and operating-system APIs, so exact-major and platform tests are required. Treat reported fields as hints rather than authoritative telemetry, label missing or truncated values, and prefer maintained PostgreSQL views and host monitoring for production inventory.
