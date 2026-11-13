## Usage

Sources:

- [PGXN metadata](https://api.pgxn.org/dist/pg_integrity_checksums/2.5.0/META.json)
- [Official README](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/README.md)
- [Control file](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/pg_integrity_checksums.control)
- [Function reference](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/docs/functions.md)
- [Operations guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/docs/operations.md)
- [Upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/pg_integrity_checksums--2.2.0--2.3.0.sql)
- [Migration guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/MIGRATION.md)
- [Upgrade and baseline guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/docs/installation.md)
- [2.4 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/pg_integrity_checksums--2.3.0--2.4.0.sql)
- [2.5 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.5.0/pg_integrity_checksums--2.4.0--2.5.0.sql)

`pg_integrity_checksums` computes physical and logical checksums for PostgreSQL pages, cells, tuples, tables, indexes and databases. This C extension supports PostgreSQL 12–19 upstream; the catalog's Pigsty packages target PostgreSQL 14–18. Installation and database-wide checks require a superuser.

### Install and migrate

```sql
CREATE EXTENSION pg_integrity_checksums;
```

### Upgrade to 2.5.0

Version 2.5.0 defaults `pg_integrity_checksums.strict` to `on`: database-wide checks fail if a relation cannot be checked, including invalid indexes, unsupported access methods or scan errors. Setting it to `off` explicitly preserves skip-with-warning behavior; warnings now count skipped relations. The new library default takes effect in new sessions. Install matching files, reconnect and update the SQL version:

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.5.0';
SHOW pg_integrity_checksums.strict;
```

The 2.4.0 changes included in this upgrade affect older baselines for dropped columns or composite attributes, user-defined types whose schema was outside `search_path`, and physical tuple headers with `include_header = true`. Recompute and review those baselines after upgrading. Other checksum values are unchanged; 2.5.0 itself changes no checksum algorithm. The upgrade also revokes public execution of generic database-wide overloads, which already enforce superuser access in C.

### Version 2.3.0 and page verification

Version 2.3.0 adds `pg_page_checksum_status(regclass, integer)`, which compares a page's calculated 16-bit checksum with its stored `pd_checksum`. The following checks remain available after the upgrade:

```sql
SHOW data_checksums;
SELECT pg_page_checksum_status('public.orders'::regclass, 0);
```

The relation must exist and contain the requested block. Results are `ok` for a match, `mismatch` for a difference, `new` for an uninitialized page, or `checksums_disabled` when stored page checksums are disabled. Stored-page verification requires `data_checksums = on`; an invalid relation or block raises an error.

`pg_page_checksum()` only calculates PostgreSQL's 16-bit page checksum; it does not compare the stored value. Its result of zero can mean an uninitialized page or a calculated zero, so zero alone does not establish validity. The status function reads relation storage directly and reports the page it reads, rather than a persistent guarantee about later writes.

The FNV-1a 32/64-bit overloads introduced in 2.2.0 still return raw `bytea` digests alongside the legacy 32-bit SQL API. They are separate from PostgreSQL's 16-bit page-checksum format. Changing the selected algorithm or overload requires a matching baseline; the 2.3.0 page-status addition does not itself require recomputing all existing baselines.

### Migration from the old extension name

Version 2.0.0 is the renamed continuation of the earlier `pg_checksums` extension. PostgreSQL has no `ALTER EXTENSION ... RENAME`; review dependent objects, drop the old extension, create the new one, and change the GUC prefix before using the new name:

```sql
DROP EXTENSION pg_checksums;
CREATE EXTENSION pg_integrity_checksums;
SET pg_integrity_checksums.workers = 4;
SET pg_integrity_checksums.strict = on;
```

The old and new names cannot coexist because they define the same SQL functions. The unrelated core `pg_checksums` utility is a separate program.

### Checksum levels

Use logical checksums for content comparison and physical checksums for storage integrity. Logical table checksums do not require a primary key, preserve duplicate rows, and exclude relation and tuple identifiers. Logical index checksums are derived from visible table rows and the index definition, so they are stable across `VACUUM`, `CLUSTER` and `REINDEX`.

```sql
SELECT pg_table_logical_checksum('public.orders'::regclass);
SELECT pg_index_logical_checksum('public.orders_pkey'::regclass);
SELECT pg_database_logical_checksum(false, true);
```

Physical checksums include storage details and can change after `VACUUM FULL`, `CLUSTER`, `REINDEX` or other physical rewrites. Use the page, tuple and table physical functions when those changes are relevant to the integrity baseline.

### Privileges and maintenance

Relation-level functions require table-level `SELECT` and reject row-level-security protected relations because direct storage scans cannot apply RLS; column-only grants are insufficient. Database-level functions are superuser-only. `pg_integrity_checksums.workers` controls parallel scans; `pg_integrity_checksums.strict` makes database scans fail when a relation cannot be checked instead of skipping it. The extension does not require shared preload. Preserve a known-good baseline for the same algorithm and compatible schema, and account for physical rewrites when comparing physical checksums. These checksums are not cryptographic authentication.
