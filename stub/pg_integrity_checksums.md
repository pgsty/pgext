## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.2/Changes)
- [Official source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.2/LICENSE)
- [Official source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.2/pg_integrity_checksums.control)

- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/README.md)
- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/pg_integrity_checksums.control)
- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/Changes)
- [Current version source](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.1/pg_integrity_checksums--2.9.0--2.9.1.sql)

- [2.9.0 Changes](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/Changes)
- [2.8.0 to 2.9.0 SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/pg_integrity_checksums--2.8.0--2.9.0.sql)
- [2.9.0 control](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/pg_integrity_checksums.control)
- [2.9.0 function reference](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.9.0/docs/functions.md)
- [Current upstream](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/Changes)
- [Current upstream](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.6.0--2.7.0.sql)

- [PGXN metadata](https://api.pgxn.org/dist/pg_integrity_checksums/2.7.0/META.json)
- [Official README](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/README.md)
- [Control file](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums.control)
- [Function reference](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/functions.md)
- [Operations guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/operations.md)
- [2.5 to 2.6 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.5.0--2.6.0.sql)
- [2.4 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.3.0--2.4.0.sql)
- [2.5 upgrade SQL](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/pg_integrity_checksums--2.4.0--2.5.0.sql)
- [Migration guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/MIGRATION.md)
- [Upgrade and baseline guide](https://api.pgxn.org/src/pg_integrity_checksums/pg_integrity_checksums-2.7.0/docs/installation.md)

`pg_integrity_checksums` computes physical and logical checksums for PostgreSQL pages, cells, tuples, tables, indexes and databases. This C extension supports PostgreSQL 12–19 upstream. Pigsty packages target PostgreSQL 15–19. Installation and database-wide checks require a superuser.

### Install and migrate

```sql
CREATE EXTENSION pg_integrity_checksums;
```

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

### Upgrade and baseline changes in 2.7.0

Install matching files, reconnect and apply the SQL upgrade:

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.7.0';
SHOW pg_integrity_checksums.strict;
```

Database-level logical checksums now include qualified relation names, so every database logical baseline changes even if row contents do not. Some physical index baselines also change after index metadata masking was corrected. Recompute and review these baselines; physical checksums are still sensitive to real storage changes, including row locks and rolled-back writes. A checksum difference alone does not establish data corruption.

Tuple and cell checksums now check MVCC visibility and return NULL for invisible rows. Index expressions execute as the table owner in a restricted context, and serializable scans take the required conflict locks. Page-status checks retry disk mismatches; a relation dropped during a database scan is skipped with NOTICE even in strict mode. Other unreadable relations still follow `pg_integrity_checksums.strict`, which defaults to on. These semantics differ from the earlier 2.6.0 masking and baseline rules.

### 2.9.0 Algorithms, Sequences and Baselines

```sql
ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.9.0';
```

Version 2.9.0 adds `crc32c` and `xxh3` to the generic algorithm functions and extends parallel workers to `pg_table_bucket_checksums()`. New sequence logical checksums compare a sequence's definition with `last_value` and `is_called`; new relation-page status functions verify all pages and forks in one call.

When `include_system = false`, database logical and physical checksums now exclude `information_schema` as well as `pg_catalog`. Because `information_schema` content varies by PostgreSQL version, every non-system database checksum from an older release changes after this correction: recompute and replace those stored baselines. Table, index and sequence checksum values are not changed by this fix, and CLI baselines already excluded `information_schema`.


### 2.9.1 Scan Changes

On PostgreSQL 17+, table, index, bucket and database scans use read streams; PostgreSQL 18 can use asynchronous I/O. Parallel workers process larger ranges. Version 2.9.1 does not change checksum values. Install the matching library and apply `ALTER EXTENSION pg_integrity_checksums UPDATE TO '2.9.1'`; the older 2.9.0 information_schema exclusion still requires replacing affected non-system database baselines.

### Version 2.9.2 and Checksum Baselines

In databases using neither UTF8 nor SQL_ASCII, logical checksums involving XML can change when the old baseline was computed with client encoding equal to server encoding. Version 2.9.2 removes that session-dependent XML serialization. UTF8 and SQL_ASCII databases are unaffected by this fix.

Physical checksums now include BRIN range-map contents and relevant legacy index-page data previously treated as unused space. Recapture physical baselines for BRIN indexes and indexes upgraded from PostgreSQL 10 or earlier, including database-level physical baselines containing them. This physical fix does not change logical checksums. Install matching files, update the extension and replace sessions loading the old library. PostgreSQL 19 is supported. Version 2.9.2 is distributed under the PostgreSQL License; earlier releases retain their original license.
