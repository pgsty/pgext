## Usage

Sources:

- [Pinned upstream README](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/README.md)
- [Pinned PostgreSQL configuration reference](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/doc/src/config-pg.md)
- [Pinned function guide](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/doc/src/use-plrust.md)
- [Pinned extension control file](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/plrust/plrust.control)
- [Pinned crate metadata](https://github.com/pgcentralfoundation/plrust/blob/5a73fa70987d47ac003008a3369a970b846b617e/plrust/Cargo.toml)

plrust is a procedural-language handler that compiles each LANGUAGE plrust function into native machine code using Rust and pgrx. It exposes PostgreSQL data types, SPI queries, set-returning functions, and triggers through Rust APIs.

### Required configuration and example

The reviewed release requires preload and a writable compilation work directory:

```conf
shared_preload_libraries = 'plrust'
plrust.work_dir = '/var/lib/postgresql/plrust-work'
```

```sql
CREATE EXTENSION plrust;

CREATE FUNCTION add_two_numbers(a numeric, b numeric)
RETURNS numeric
STRICT
LANGUAGE plrust
AS $$
    Ok(Some(a + b))
$$;

SELECT add_two_numbers(2, 2);
```

CREATE FUNCTION invokes the Rust toolchain and may take significant time. The control file's extension object version is 1.1, matching this catalog, while the pinned crate metadata says 1.2.8. Do not confuse the SQL object version with the crate/package version.

### Trust and build-chain boundary

A trusted PL/Rust build is intended only for supported x86_64 or aarch64 Linux configurations using its restricted postgrestd environment and mandatory lints. Other builds are untrusted and do not provide the same containment. Trusted and untrusted compiled functions are different targets and cannot be mixed transparently.

The compiler, linker, work directory, generated shared objects, pgrx bindings, and Rust dependencies become part of the database attack surface. If plrust.allowed_dependencies is unset, the reviewed documentation says all crates are allowed. A user allowed to create functions can consume CPU, disk, memory, and dependency-download resources and may introduce supply-chain code. Use an explicit pinned dependency allowlist, immutable toolchain, restricted writable directory, no unnecessary network access, and tight language CREATE privileges. Do not weaken the required lints; upstream states that doing so removes the expectation of trusted operation.
