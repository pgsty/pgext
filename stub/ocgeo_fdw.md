## Usage

Sources:

- [Pinned upstream README](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/README.md)
- [Version 1.0 installation SQL](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/sql/ocgeo_fdw--1.0.sql)
- [Pinned FDW implementation](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/ocgeo_fdw.c)
- [Pinned API client implementation](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/ocgeo_api.c)

ocgeo_fdw version 1.0 maps OpenCage's JSON geocoding API to a read-only foreign table. An equality condition on the q input column performs forward or reverse geocoding; a minimum confidence condition can be sent as an API parameter. Returned fields can include the raw JSON, PostgreSQL point and box values, formatted address, confidence, and address components.

### Server, mapping, and query

```sql
CREATE EXTENSION ocgeo_fdw;

CREATE SERVER ocdata_server
FOREIGN DATA WRAPPER ocgeo_fdw
OPTIONS (uri 'https://api.opencagedata.com/geocode/v1/json');

CREATE USER MAPPING FOR current_user
SERVER ocdata_server
OPTIONS (api_key '<api-key>', max_reqs_sec '1');

CREATE FOREIGN TABLE geocode_api (
    json_response jsonb,
    location point,
    confidence integer,
    formatted text,
    q text
) SERVER ocdata_server;

SELECT confidence, location, formatted
FROM geocode_api
WHERE q = 'Eiffel Tower, France' AND confidence >= 5;
```

Without a usable q equality restriction, a simple scan makes no API request. A nested-loop join can make one request per outer row.

### Network, quota, and observability limits

Every API request is synchronous inside the PostgreSQL backend and can hold a transaction, connection, and locks while DNS, TLS, rate-limit sleep, and remote execution complete. The built-in limiter controls requests per second only within one backend. It does not enforce daily quota globally across sessions. Joins and rescans can multiply calls, so materialize or cache stable results.

The API key is stored in a user mapping. Debug output can include the full request URL; the README explicitly shows enabling client debug to view it, so do not enable that logging where the key could enter client or server logs. Restrict server usage, apply outbound allowlists, use a low-privilege key, set statement timeouts, and monitor the provider's authoritative quota counters.

The ocgeo_stats function reports process-local in-memory counters, can include entries for roles used by that backend, and is incorrectly declared IMMUTABLE even though the values change after requests. Do not use it as an audit or billing source. The reviewed source documents only PostgreSQL 9.6 through 14 and has not changed since 2021; libcurl, planner APIs, JSON parsing, TLS policy, provider response fields, and current server majors all require fresh integration tests.
