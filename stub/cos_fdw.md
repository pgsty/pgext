## Usage

Sources:

- [Tencent Cloud cos_fdw guide](https://cloud.tencent.com/document/product/409/73657)
- [TencentDB for PostgreSQL extension matrix](https://cloud.tencent.com/document/product/409/75121)

`cos_fdw` is a TencentDB for PostgreSQL provider extension that exposes headerless CSV files in Tencent Cloud Object Storage (COS) as read-only foreign tables. Use it for cold-data queries and migration into local tables; it is not a portable community FDW or a general object-storage writer.

### Core Workflow

Export data as CSV without a header, upload it to a COS bucket in an appropriate region, and then create the dependency, wrapper, server, and foreign table. The `host` value is the bucket access domain without an `http://` or `https://` prefix.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION cos_fdw;

CREATE SERVER cos_server
FOREIGN DATA WRAPPER cos_fdw
OPTIONS (
    host 'example.cos.ap-nanjing.myqcloud.com',
    bucket 'example-bucket',
    id 'COS_SECRET_ID',
    key 'COS_SECRET_KEY'
);

CREATE FOREIGN TABLE archive_sensor_log (
    sensor_log_id integer,
    location text OPTIONS (force_not_null 'true'),
    reading bigint,
    reading_date timestamp
)
SERVER cos_server
OPTIONS (
    filepath '/archive/sensor_log.csv',
    format 'csv',
    null 'NULL'
);

SELECT *
FROM archive_sensor_log
WHERE reading_date >= timestamp '2026-01-01';
```

To map several objects to one table, place comma-separated paths with no extra spaces in `filepath`. Use `EXPLAIN` to inspect the COS URL, per-file sizes, and total estimated size. Copy selected rows into a local table with `INSERT INTO local_table SELECT ... FROM archive_sensor_log` when local indexing or transactional writes are needed.

### Options and Boundaries

- Server options include `host`, `bucket`, `id`, and `key`.
- Table options include `filepath`, `format 'csv'`, CSV parsing options such as `delimiter` and `null`, and per-column `force_not_null`.
- The extension reads foreign data; inserts routed to a COS foreign partition are not supported.
- `pgcrypto` is required because TencentDB encrypts the stored COS credential fields. Treat `id` and `key` as secrets even though catalog storage is encrypted.

### Provider Operations

Availability, supported PostgreSQL versions, privileges, upgrades, and credential handling are controlled by TencentDB. Check the extension matrix for the exact instance kernel before running `CREATE EXTENSION`.

COS reads add network latency, request cost, and an external consistency boundary. Restrict who can create or alter foreign servers, scope COS credentials to the required bucket and paths, keep the database and bucket in suitable regions, and test malformed rows and missing objects before using a foreign table in production queries.
