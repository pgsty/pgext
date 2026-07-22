## Usage

Sources:

- [Official README](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/readme.md)
- [Random FDW implementation](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/src/fdw/mod.rs)
- [GUC implementation](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/src/utils/guc.rs)

`pgrand` generates synthetic rows through a PostgreSQL foreign data wrapper. It is intended for quickly filling test tables or producing arrays for vector tests; the generated values are test data, not statistically uniform or cryptographically secure randomness.

### Core Workflow

The extension installs `random_fdw_handler`, but the wrapper and server must be created explicitly.

```sql
CREATE EXTENSION pgrand;
CREATE FOREIGN DATA WRAPPER random
HANDLER random_fdw_handler;
CREATE SERVER random_server
FOREIGN DATA WRAPPER random;

SET random.seed = 42;
SET random.array_length = 8;
SET random.min_text_length = 8;
SET random.max_text_length = 20;

CREATE FOREIGN TABLE generated_orders (
    order_id integer,
    amount numeric(10,2),
    tags text[],
    payload jsonb,
    source_ip inet
)
SERVER random_server
OPTIONS (total '1000');

SELECT * FROM generated_orders LIMIT 10;

INSERT INTO test_orders
SELECT * FROM generated_orders;
```

The foreign-table option `total` controls the number of rows emitted by a scan and defaults to `1000`. A nonzero `random.seed` initializes a repeatable ChaCha8 stream for each scan; the default `0` uses entropy.

### GUCs

- `random.min_int` and `random.max_int` bound integer generation.
- `random.min_text_length` and `random.max_text_length` bound text length.
- `random.array_length` sets array length from `1` to `16384`.
- `random.float_scale` controls generated floating-point scale.
- `random.seed` selects deterministic or entropy-seeded output.

The implementation generates many common numeric, text, array, JSON, network, and geometric PostgreSQL types. Unsupported types yield NULL rather than a value.

### Caveats

Upstream explicitly states that generated numbers are not evenly distributed and user-defined types are unsupported. The table access method is still marked in progress; the documented usable surface is the FDW. Scans synthesize all requested rows and do not push predicates to an external source, so use `total` conservatively and materialize only the data required. The reviewed `0.1.0` Cargo feature matrix targets PostgreSQL `15` and `16`; the extension is non-relocatable and requires superuser installation. No preload or restart is required.
