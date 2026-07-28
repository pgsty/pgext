## Usage

Sources:

- [Official upstream README](https://github.com/cairn-ehr/cairn-ehr/blob/bf5d2cf13c493edcba84685f2f8714851f64ea6e/README.md)
- [Official extension control file (cairn_pgx.control)](https://github.com/cairn-ehr/cairn-ehr/blob/bf5d2cf13c493edcba84685f2f8714851f64ea6e/extensions/cairn_pgx/cairn_pgx.control)
- [Official implementation source](https://github.com/cairn-ehr/cairn-ehr/blob/bf5d2cf13c493edcba84685f2f8714851f64ea6e/extensions/cairn_pgx/src/lib.rs)

`cairn_pgx` — Cairn in-database verify gate — COSE Sign1/Ed25519 event verification (Spike 0002). Use it when implementing the corresponding security, audit, or access-control workflow. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION cairn_pgx;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `cairn_actor_id` is an extension function.
- `cairn_attestation_ok` is an extension function.
- `cairn_blob_verify` is an extension function.
- `cairn_blob_verify_error` is an extension function.
- `cairn_body` is an extension function.
- `cairn_pgx_version()` is an extension function.
- `cairn_unseal_body` is an extension function.
- `cairn_verify` is an extension function.
- `cairn_verify_error` is an extension function.
- `cairn_wrap_dek` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
