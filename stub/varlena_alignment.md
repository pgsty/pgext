## Usage

Sources:

- [Official defect analysis](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/README.md)
- [Reproduction tests](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/src/lib.rs)
- [Pinned pgrx dependency and PostgreSQL features](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/Cargo.toml)

`varlena_alignment` version `0.0.0` is a developer reproduction for a historical pgrx alignment bug. It demonstrates that casting a one-byte-aligned PostgreSQL `varlena` pointer to four-byte-aligned `varattrib_4b` in `set_varsize_4b` can invoke undefined behavior.

### Intended Workflow

The repository defines only `#[pg_test]` cases. It exports no application SQL functions, types, operators, or configuration, so installing it with `CREATE EXTENSION` provides no user-facing workflow. Its purpose is to run the pgrx test harness against the pinned dependency:

```sh
cargo pgrx test
```

One test calls `set_varsize_1b` on a deliberately offset allocation and is expected to succeed. The other calls `set_varsize_4b` on the same one-byte-aligned `varlena` pointer to reproduce the stricter-alignment failure visible with Rust debug pointer checks.

### Developer Boundary

This project pins `pgrx` version `0.11.1` and declares PostgreSQL feature sets from `pg11` through `pg16`. It is evidence about that implementation context, not a general runtime diagnostic and not a fix. Do not install it in production or treat a passing release build as proof that an unsafe cast is valid; assess the relevant pgrx version and memory-alignment contract directly.
