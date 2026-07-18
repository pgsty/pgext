## Usage

Sources:

- [Project README and compatibility notes](https://github.com/pgspider/influxdb_fdw/blob/9cb845ecee3f6a66c2abd2bcdf6a7a5266e756f7/README.md)
- [Extension control file](https://github.com/pgspider/influxdb_fdw/blob/9cb845ecee3f6a66c2abd2bcdf6a7a5266e756f7/influxdb_fdw.control)
- [Version 1.3 SQL API](https://github.com/pgspider/influxdb_fdw/blob/9cb845ecee3f6a66c2abd2bcdf6a7a5266e756f7/influxdb_fdw--1.3.sql)

`influxdb_fdw` 1.3 exposes InfluxDB measurements as PostgreSQL foreign tables. The reviewed upstream matrix covers PostgreSQL 13 through 17, InfluxDB 1.8, and InfluxDB 2.7 through its v1 compatibility API. Supported operations include reads, inserts, deletes, schema import, and several query pushdowns; `TRUNCATE` is not supported.

### Define a foreign server and table

The server below targets an InfluxDB 1.x database. Store remote credentials in a separately managed user mapping with the least required privilege rather than in application SQL or logs.

```sql
CREATE EXTENSION influxdb_fdw;

CREATE SERVER influxdb_svr
  FOREIGN DATA WRAPPER influxdb_fdw
  OPTIONS (dbname 'metrics', host 'http://127.0.0.1', port '8086');

CREATE FOREIGN TABLE sensor_reading (
  time timestamptz,
  device_id text,
  temperature double precision
)
SERVER influxdb_svr
OPTIONS (table 'sensor_reading', tags 'device_id');

SELECT time, device_id, temperature
FROM sensor_reading
WHERE time >= now() - interval '15 minutes';
```

InfluxDB 2.x requires the C++ client path, `version '2'`, and an `auth_token` user-mapping option. Protect PostgreSQL catalogs and backups that contain user mappings. In schemaless mode a foreign table has fixed `time`, `tags`, and `fields` columns, with the last two represented as `jsonb`.

### Semantic and operational boundaries

InfluxDB can omit rows containing null fields when the remote target list names specific fields, while a remote `SELECT *` can return them. Consequently, changing the PostgreSQL projection may change which rows are visible. Inspect `EXPLAIN VERBOSE` and compare results against native InfluxDB queries before depending on pushdown behavior.

Writes cross a foreign-system boundary and do not acquire PostgreSQL's local transactional guarantees. Treat retries, partial failures, time precision, retention policy, and remote authentication as application-level concerns. The `time` column uses microsecond precision; `time_text` is available when nanosecond text representation is required. Validate the exact PostgreSQL, client-library, and InfluxDB combination because the reviewed compatibility matrix does not establish support beyond PostgreSQL 17.
