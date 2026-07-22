## Usage

Sources:

- [Official upstream documentation](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/readme.md)
- [Official Rust implementation](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/src/lib.rs)
- [Official Rust package manifest](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/Cargo.toml)

`bfn` 2.0.1 is a pgrx collection of general-purpose PostgreSQL helper functions for UUIDv7, dates, text sanitization, JSON addresses, parsing, hashing, and numeric defaults. It is not a single data model: review and grant only the functions an application actually needs. The release declares build features for PostgreSQL 16 through 18 and requires superuser installation.

### Install in Its Schema

The extension is non-relocatable and upstream installs it in a dedicated `bfn` schema:

```sql
CREATE SCHEMA IF NOT EXISTS bfn;
CREATE EXTENSION bfn SCHEMA bfn;

SELECT bfn.bfn_version();
SELECT bfn.new_uuid();
SELECT bfn.uuid_to_ts(bfn.new_uuid());
SELECT bfn.date_range(DATE '2026-01-01', DATE '2026-01-03');
```

Schema-qualify calls rather than placing a broad utility schema first in an untrusted `search_path`.

### Important Function Groups

- `new_uuid()` creates a UUIDv7 value and `uuid_to_ts(uuid)` extracts its embedded timestamp, returning NULL for a non-v7 or invalid value.
- `all_dates_from(date, date)`, `date_range(date, date)`, `first_day_of_month(date)`, and `last_day_of_month(date)` provide date helpers.
- `to_address(...)` constructs a JSONB address with optional coordinates.
- `isi(...)`, `is_empty(...)`, `is_null(...)`, `not_null(...)`, `trim(...)`, `san_trim(...)`, `strip_tags(...)`, and `upper_first(text)` provide common predicates and text normalization.
- `random_base64()`, `md5_as_base64(text)`, `md5_as_uuid(text)`, and their verification helpers generate tokens or deterministic MD5-derived values.
- Parsing, name-joining, metric-scaling, and zero-default helpers cover application-specific convenience cases.

### Security and Compatibility

MD5-derived helpers are not suitable for password storage, signatures, or collision-resistant identity. Treat `random_base64()` output according to the guarantees of the pinned Rust `rand` dependency, not as a documented authentication protocol. Text sanitizers are convenience transformations, not HTML or SQL security boundaries. Pin the extension and Rust dependency graph, because output formats and parser behavior can affect stored values and indexes. Test UUID ordering, time extraction, Unicode case conversion, numeric precision, and NULL/default semantics on the exact PostgreSQL major before using a helper in constraints, generated columns, or persistent keys.
