## Usage

Sources:

- [Upstream README](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/README.md)
- [Extension control file](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/pg_temporal.control)
- [Cargo package metadata](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/Cargo.toml)
- [Rust implementation](https://github.com/Wundero/pg_temporal/blob/48e7a74f38d59d59f1006180bf3a1528c0b5f632/src/lib.rs)

`pg_temporal` version `1.0.0` adds a `ZonedDateTime` type backed by the Rust Jiff library, intended to represent a local date/time together with an offset and named time zone. The type has input/output, equality, ordering, and hashing support; `zdt_now()` returns the current zoned time.

### Example

```sql
CREATE EXTENSION pg_temporal;
SELECT zdt_now(), pg_typeof(zdt_now());
```

The current API is very small: it supplies the type and current-time function but no documented conversions or arithmetic. Installation is superuser-only. Parsing uses the bundled Jiff/tz data and invalid text reaches an unwrap path; behavior can differ from PostgreSQL `timestamptz` and its time-zone database. Validate accepted serialization, comparison semantics, upgrades, and tz-data changes before using the type in durable schemas.
