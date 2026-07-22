## Usage

Sources:

- [Archived snapshot of the official upstream repository](https://github.pkg.st/dmtkfs/pg-tamper-log)
- [Last-known official repository URL](https://github.com/dmtkfs/pg-tamper-log)

`pg_tamperlog_rust` is the optional pgrx helper shipped with `pg_tamperlog` 1.1. It supplies the faster SHA-256 calculation used by the SQL audit-chain extension; it does not create or protect an audit log by itself.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_tamperlog_rust;
CREATE EXTENSION pg_tamperlog;
```

The archived upstream instructions build the helper with `cargo pgrx` for a specific PostgreSQL installation. Native pgrx modules are tied to the target PostgreSQL major and Rust/C ABI: build reproducibly against the exact server, then compare its byte-for-byte hash results with an independent SHA-256 implementation for empty, `NULL`, Unicode, large, and concurrent inputs.

Hash speed does not strengthen the chain's threat model. A privileged actor can replace the native library or SQL wrapper, and a crash or upgrade can change runtime behavior. Restrict function and library installation, hash the deployed artifact, monitor reloads, and retain signed off-host checkpoints.

The official GitHub repository now returns HTTP 404; the source above is an archived rendering of its README, not a complete auditable source archive. Do not trust a similarly named crate or binary until the exact content-addressed source has been recovered and reviewed.
