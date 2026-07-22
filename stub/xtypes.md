## Usage

Sources:

- [Official README](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/README.md)
- [Official extension SQL](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/xtypes--0.1.sql)
- [Official bytes implementation](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/bytes.c)

`xtypes` version `0.1` currently provides a `bytes` type for storing byte quantities as unsigned 64-bit values, formatting them like `pg_size_pretty`, and performing limited arithmetic. It is useful for human-readable capacity fields when its unsigned range and small operator surface are acceptable.

### Core Workflow

Create the extension, store capacity values with explicit units, and use the supplied arithmetic:

```sql
CREATE EXTENSION xtypes;

CREATE TABLE storage_limits (
  service text PRIMARY KEY,
  quota bytes NOT NULL
);

INSERT INTO storage_limits VALUES ('archive', '1.5 TiB');

SELECT service, quota, quota + '512 GiB'::bytes
FROM storage_limits;
```

The output automatically selects a binary unit, for example `1.5 TB` or `2 TB`.

### Type and Operators

- `bytes` stores a base-byte count in an unsigned 64-bit representation.
- Input accepts `B`, `bytes`, `kB`/`KB`/`KiB`, and corresponding `MB`/`MiB` through `EB`/`EiB` suffixes. These scale by powers of 1024 even where the short label lacks `i`.
- Operators include `+`, `-`, `<`, `>`, multiplication by `int8` in either order, and division by `int8`.
- `round(bytes)` rounds the displayed quantity at its automatically selected unit and returns another `bytes` value.

### Operational Caveats

The maximum representable value is just below 16 EiB. Addition, subtraction, and multiplication check unsigned overflow or underflow, but the reviewed input and signed-`int8` arithmetic paths deserve explicit tests for negative values and boundary cases. Fractional input is converted to whole bytes, and formatted output uses short binary-scaled labels such as `KB`, which can be confused with decimal SI units. The SQL defines no equality operators, casts, or B-tree operator class, so do not assume ordinary unique-index or ordering behavior beyond the explicitly supplied operators. Pin the binary to the target PostgreSQL major version and validate dump/restore and client-driver text handling.
