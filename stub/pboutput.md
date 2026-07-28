## Usage

Sources:

- [Official extension control file (pboutput.control)](https://github.com/semtexzv/pboutput/blob/2a1062eee71acec95866a15086edb5060caf1d66/pboutput.control)
- [Official implementation source](https://github.com/semtexzv/pboutput/blob/2a1062eee71acec95866a15086edb5060caf1d66/src/lib.rs)
- [Official Rust package manifest](https://github.com/semtexzv/pboutput/blob/2a1062eee71acec95866a15086edb5060caf1d66/Cargo.toml)

`pboutput` — Binary logical-decoding output plugin that encodes PostgreSQL changes as Protocol Buffers. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pboutput;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
