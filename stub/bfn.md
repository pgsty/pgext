## Usage

Sources:

- [bfn README at the reviewed commit](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/readme.md)
- [bfn Rust API at the reviewed commit](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/src/lib.rs)
- [bfn.control at the reviewed commit](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/bfn.control)

`bfn` is a Rust/`pgrx` utility collection. Version 2.0.1 exposes helpers for UUIDv7 generation and timestamp extraction, date ranges and month boundaries, string normalization and validation, JSONB address construction, and hashing or random-value operations.

The upstream installation places the extension in a schema named `bfn`. Representative functions include `bfn_version`, `new_uuid`, `uuid_to_ts`, `date_range`, `first_day_of_month`, and `san_trim`.

### Install and Inspect

```sql
CREATE SCHEMA bfn;
CREATE EXTENSION bfn SCHEMA bfn;

SELECT bfn.bfn_version();
SELECT bfn.new_uuid();
SELECT bfn.uuid_to_ts(bfn.new_uuid());
SELECT bfn.date_range(DATE '2026-01-01', DATE '2026-01-03');
```

### Caveats

- The reviewed control file sets `superuser = true`; installation therefore requires an appropriately privileged role.
- The reviewed Cargo manifest targets PostgreSQL 16, 17, and 18 and reports version 2.0.1, matching the catalog version.
- This is a broad utility collection rather than one narrowly scoped feature. Check the pinned Rust API for exact null handling and argument types before embedding a helper in application SQL.
