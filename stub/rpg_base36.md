## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/README.md)
- [SQL function definition](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/sql/rpg_base36.sql)
- [Rust implementation](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/src/lib.rs)
- [Cargo package metadata](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/Cargo.toml)
- [Extension control file](https://github.com/durch/rpg_base36/blob/7a04c31887f2965e5f0a6fe904258765349eb9ec/rpg_base36.control)

`rpg_base36` is a historical toy Rust extension with one function, `base36_encode(integer)`, that converts a 32-bit PostgreSQL integer to uppercase base 36. Zero becomes `0`, and ordinary negative values keep a leading minus sign. It provides no decoder or alternate alphabet.

### Core Workflow

```sql
CREATE EXTENSION rpg_base36;

SELECT base36_encode(0);       -- 0
SELECT base36_encode(35);      -- Z
SELECT base36_encode(123456);  -- 2N9C
SELECT base36_encode(-35);     -- -Z
```

The canonical extension name is `rpg_base36`, even though the SQL load-guard message contains the stale name `rust_base36`. The control file and SQL extension version are 0.0.1, while the Rust crate metadata says 0.1.0; that crate version does not establish a newer database extension release.

### Compatibility and Input Boundaries

The implementation calls Rust integer absolute value before conversion. The minimum 32-bit integer, `-2147483648`, cannot be represented as a positive 32-bit integer and can overflow or panic in the reviewed old toolchain; reject that value.

This project predates `pgrx` and uses old `rpgffi` bindings, C glue, and unsafe PostgreSQL function-call access. Build and test it for the exact PostgreSQL major, Rust compiler, and target architecture. Treat it as a learning example rather than a maintained production dependency.
