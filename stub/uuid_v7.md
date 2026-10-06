## Usage

Sources:

- [examples/uuid_v7/README.md](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/README.md)
- [examples/uuid_v7/uuid_v7--1.0.sql](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/uuid_v7--1.0.sql)
- [examples/uuid_v7/Makefile](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/Makefile)
- [examples/uuid_v7/uuid_v7.control](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/uuid_v7.control)

`uuid_v7` is an AWS Trusted Language Extensions example implemented in PL/Rust. It generates time-based UUIDs, creates one for a supplied timestamp and extracts its millisecond timestamp.

### Core Workflow

```sql
CREATE EXTENSION plrust;
CREATE EXTENSION uuid_v7;
SELECT generate_uuid_v7();
SELECT timestamptz_to_uuid_v7('2026-10-06 00:00:00+00'::timestamptz);
SELECT uuid_v7_to_timestamptz(generate_uuid_v7());
```

### Operational Boundaries

Install `pg_tle` and PL/Rust 1.2.0 or later, then register the example through its documented TLE installation process. Once it appears among available extensions, create it in the database. The control’s SQL dependency is `plrust`; this is not a separate native pgrx shared library.

`generate_uuid_v7` uses the current time; `timestamptz_to_uuid_v7` embeds a supplied timestamp; `uuid_v7_to_timestamptz` decodes the first timestamp bits. The decoder does not establish that an arbitrary UUID is a valid version-7 value. Time is millisecond-granular and random suffixes do not provide strict total ordering within a millisecond. The example was developed against a UUID v7 draft; assess its semantics before using it as a standards-conformance boundary.
