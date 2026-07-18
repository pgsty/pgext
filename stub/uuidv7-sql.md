## Usage

Sources:

- [Upstream README](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/README.md)
- [Version 1.0 generation functions](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/sql/uuidv7-sql--1.0.sql)
- [Version 1.1 encryption functions](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/sql/uuidv7-sql--1.0--1.1.sql)
- [Version 1.1 control file](https://github.com/dverite/postgres-uuidv7-sql/blob/396a44433e6e0eb63b1d9d1517e9098256d97351/uuidv7-sql.control)
- [PostgreSQL 18 UUID functions](https://www.postgresql.org/docs/18/functions-uuid.html)

uuidv7-sql 1.1 is a pure-SQL implementation of UUIDv7 generation, timestamp extraction, range boundaries, sub-millisecond variants, and reversible UUIDv4 or UUIDv8 transformation.

### Generate and inspect

Install in an explicit schema so calls remain unambiguous on PostgreSQL 18 and later:

```sql
CREATE EXTENSION "uuidv7-sql" WITH SCHEMA public;
SELECT public.uuidv7() AS id;
SELECT public.uuidv7_extract_timestamp(public.uuidv7());
SELECT public.uuidv7_boundary('2026-01-01 00:00:00+00');
```

PostgreSQL 18 includes a pg_catalog uuidv7 function that wins normal name resolution. Prefer the built-in function for ordinary generation, or schema-qualify this extension's version as shown.

### Caveats

- UUIDv7 exposes its creation timestamp. Do not place raw values in contexts where approximate event time is sensitive.
- Values generated within the same millisecond contain random ordering bits and are not guaranteed to be strictly increasing. The sub-millisecond variant also does not provide a global monotonic sequence.
- uuidv7_extract_timestamp does not verify the input version or variant; it interprets the first 48 bits of any UUID as milliseconds.
- uuidv7_boundary is deterministic and non-random. Use it only as a range bound, never as a row identifier.
- Version 1.1's XTEA transformation is a reversible, unauthenticated 64-bit block construction, not general authenticated encryption. It does not detect tampering, define key storage or rotation, or validate that input is UUIDv7; SQL-supplied keys can also appear in logs and activity views.
- Generation depends on gen_random_uuid, which is built into PostgreSQL 13 and later; older releases need pgcrypto even though the control file does not declare that dependency.
- The repository contains no automated tests. Validate RFC conformance, timestamp boundaries, time zones, supported date range, encryption round trips, and upgrade from 1.0 before relying on it.
