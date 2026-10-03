## Usage

Sources:

- [AlohaDB README](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [Control 1.0](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar.control)
- [SQL 1.0](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar--1.0.sql)
- [Access method implementation](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar.c)
- [Storage definitions](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar.h)
- [Meson build](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/meson.build)

`alohadb_columnar` 1.0 is an experimental table access method in the AlohaDB PostgreSQL fork. It stores append-only analytical data in row stripes with per-column zstd compression. This revision does not establish normal MVCC snapshot isolation, and is unsuitable for workloads that depend on it. Stock PostgreSQL compatibility is not claimed.

### Core Workflow

The fork's Meson build includes the module only when zstd is available. There is no module Makefile in this revision. Install the built module in the matching AlohaDB server, then use a superuser to create the extension:

```sql
CREATE EXTENSION alohadb_columnar;
CREATE TABLE events_col (
  id bigint,
  ts timestamptz,
  payload jsonb
) USING columnar;
SELECT * FROM alohadb_columnar_info('events_col');
```

The implementation targets INSERT and SELECT workloads. It registers a transaction callback to flush buffered regular inserts before commit. No preload or restart requirement is declared. The extension is relocatable, but the access method name is the database-wide `columnar`, which conflicts with other implementations using that name.

### Objects and Configuration

- `columnar` is the table access method; `columnar_tableam_handler` is its internal handler.
- `alohadb_columnar_info(regclass)` returns `total_stripes`, `total_rows`, `total_size`, and `compression`. The size is the sum of compressed stripe sizes, rather than a general-purpose relation disk-size measurement. Call it for tables using this access method.
- `alohadb.columnar_stripe_row_count` controls rows per stripe. It defaults to 150000, accepts 1000 through 10000000, and is a superuser-setting parameter in the inspected implementation.

### Operational Boundaries

UPDATE, DELETE, row locking, speculative insertion, and index scans are unsupported. The snapshot-check callback returns true unconditionally; do not infer ordinary PostgreSQL MVCC visibility from the append-only design. The README's selective-column-scan claim is not established here as a verified execution property.

Treat persistent data and upgrades as experimental: this source revision provides no documented storage-format migration path or production compatibility guarantee. Review the matching fork version and retain recoverable copies before changing the module.
