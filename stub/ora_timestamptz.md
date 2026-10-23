## Usage

Sources:

- [README.md](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/README.md)
- [ora_timestamptz--1.0.sql](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/ora_timestamptz--1.0.sql)
- [ora_tstz_ddl.c](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/ora_tstz_ddl.c)
- [ora_timestamptz.control](https://github.com/zwiebelchen/ora_timestamptz/blob/069ec156567c5792e7ab2c127c5c22987a9e2e86/ora_timestamptz.control)

`ora_timestamptz` stores a UTC instant together with its original region or fixed offset. The source targets PostgreSQL 17; other majors have not been verified upstream. A superuser installs the non-relocatable extension, which also creates functions in `pg_catalog`. Ordinary type use needs no preload or restart.

### Core Workflow

```sql
CREATE EXTENSION ora_timestamptz;
CREATE TABLE events (id bigint PRIMARY KEY, occurred_at ora_timestamptz(9));
INSERT INTO events VALUES
  (1, '2026-10-16 14:30:00 Europe/Berlin'),
  (2, '2026-10-16 08:30:00 -04:00');
SELECT id, ora_timestamptz_iso(occurred_at) FROM events;
SELECT count(DISTINCT occurred_at) FROM events;
```

### Operational Boundaries

Precision ranges from 0 to 9 fractional digits, defaulting to 6. Equality, ordering and uniqueness compare UTC instants; different zones can therefore compare equal. `ora_timestamptz_identical` also compares the zone, and `ora_timestamptz_iso` emits lossless text. Use `from_tz`, `sys_extract_utc`, `to_timestamp_tz` and `to_char` for conversion and formatting. B-tree, hash and BRIN operator classes are available.

DST gaps raise errors. Ambiguous times normally select standard time; `ora_timestamptz.error_on_overlap_time` can reject them instead. Arithmetic uses UTC, and EXTRACT returns UTC date/time fields. The default `ora_timestamptz.output_format` is iso; nls output can lose information and is unsuitable for dumps.

Optional DDL rewriting uses `ora_timestamptz.replace_timestamptz` and can be limited by `ora_timestamptz.replace_schemas`. It requires `session_preload_libraries` for new sessions or `shared_preload_libraries` plus a server restart. Leave rewriting off when restoring native timestamp columns, because dump syntax can otherwise change their type. Rewriting does not change existing columns, query literals or casts. Install the library on replicas and restore targets before accessing stored values.
