## Usage

Sources:

- [Alibaba Cloud oss_fdw documentation](https://www.alibabacloud.com/help/en/polardb/polardb-for-oracle/oss-fdw)

`oss_fdw` is an Alibaba Cloud PolarDB provider extension that maps Object Storage Service directories or file prefixes to PostgreSQL foreign tables. It is designed for append-only archives and cold data on supported PolarDB for PostgreSQL-compatible service revisions, not as a portable community FDW for self-managed PostgreSQL.

### Core Workflow

Enable the provider extension, define an OSS server, and map a directory to a foreign table:

```sql
CREATE EXTENSION oss_fdw;

CREATE SERVER archive_oss
FOREIGN DATA WRAPPER oss_fdw
OPTIONS (
    host 'oss-cn-hangzhou-internal.aliyuncs.com',
    bucket 'example-archive',
    id '<access-key-id>',
    key '<access-key-secret>'
);

CREATE FOREIGN TABLE public.events_archive (
    event_id bigint,
    occurred_at timestamptz,
    payload text
)
SERVER archive_oss
OPTIONS (dir 'events/', format 'csv', compressiontype 'gzip');

INSERT INTO public.events_archive
SELECT event_id, occurred_at, payload
FROM public.events
WHERE occurred_at < current_date - 90;

SELECT count(*) FROM public.events_archive;
SELECT * FROM oss_fdw_list_file('events_archive', 'public');
```

Use exactly one of `dir` or `prefix` for the object mapping. `format` defaults to CSV. `compressiontype` can select gzip or, on supported PostgreSQL 14 service revisions, Zstandard; `compressionlevel` controls the provider-supported range.

### Write and Removal Semantics

`SELECT`, `INSERT`, and `TRUNCATE` are supported; `UPDATE` and `DELETE` are not. Each `INSERT` execution creates a new OSS file instead of modifying an existing one. `TRUNCATE public.events_archive` removes all OSS files mapped to that foreign table, so treat it as destructive object-store deletion rather than ordinary local-table cleanup.

Dropping a foreign table or the extension is distinct from truncating mapped data. Confirm the provider's behavior and preserve independent OSS retention or backups before removing definitions.

### Provider and Security Boundaries

The server options contain OSS access credentials. Use a dedicated RAM identity with the minimum bucket permissions, restrict `USAGE` on the foreign server, and avoid exposing catalog metadata or DDL to untrusted roles. Prefer an internal endpoint in the same region as the PolarDB cluster when supported; network placement affects both reachability and performance.

Availability, exact version, supported PostgreSQL revisions, regions, compression algorithms, and privileges are controlled by Alibaba Cloud. Check the extension list and service documentation for the target cluster before running the example. Access to OSS has substantially different latency and transaction behavior from local PostgreSQL storage; test retries, partial failures, file accumulation, and downstream readers before using it for archival workflows.
