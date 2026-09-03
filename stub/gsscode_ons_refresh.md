## Usage

Sources:

- [PGXN 1.0.0 README](https://pgxn.org/dist/gsscode/1.0.0/README.html)
- [gsscode_ons_refresh control file](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/gsscode_ons_refresh.control)
- [gsscode_ons_refresh 1.0.0 SQL definitions](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/gsscode_ons_refresh--1.0.0.sql)

`gsscode_ons_refresh` is the optional companion that refreshes the `gsscode_types` registry used by `gsscode`. It is useful only when an installation wants an in-database refresh; the core packed type and operators do not require it.

### Core Workflow

```sql
CREATE EXTENSION gsscode_ons_refresh CASCADE;

SELECT update_gsscode_types();
```

The control file requires `gsscode` and `http`; `CASCADE` can create those dependencies when their files are available. `update_gsscode_types()` downloads the project's generated JSON registry, upserts its rows, and returns the number processed.

### Trust and Operations

The exact 1.0.0 SQL implementation is PL/pgSQL plus `http`, despite stale PGXN metadata that describes an earlier `plpython3u` design. Enabling `http` normally requires a superuser. The database server needs outbound HTTPS access to `raw.githubusercontent.com`, so network policy, DNS, TLS trust, timeouts, and proxy behavior become part of the refresh path.

The JSON file is produced by the upstream repository's GitHub Actions workflow from an ONS release. A refresh therefore trusts ONS data, the project's automation, and repository write access. Review that supply chain before using it for authoritative data. The function performs an upsert and does not remove local rows absent from the downloaded payload; run it in a controlled maintenance session and verify the returned count and representative rows.

The upstream README recommends the external `update_gsscode_types.py` path for most installations because it avoids database-server egress and the `http` dependency. The 1.0.0 release publishes no PostgreSQL-major support matrix, so validate both dependencies on the target server.
