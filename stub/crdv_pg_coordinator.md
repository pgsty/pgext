## Usage

Sources:

- [Official upstream README](https://github.com/ricarim/secure-crdv/blob/3c224559b6c16bd850978845cf172e9fbcae6f00/README.md)
- [Official extension control file (crdv_pg_coordinator.control)](https://github.com/ricarim/secure-crdv/blob/3c224559b6c16bd850978845cf172e9fbcae6f00/mpc/crates/crdv_pg_coordinator/crdv_pg_coordinator.control)
- [Official implementation source](https://github.com/ricarim/secure-crdv/blob/3c224559b6c16bd850978845cf172e9fbcae6f00/mpc/crates/crdv_pg_coordinator/src/lib.rs)

`crdv_pg_coordinator` — Secure Conflict-free Replicated Data Views with MPC-backed private values. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION crdv_pg_coordinator;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mpc_benchmark_set` is an extension function.
- `mpc_counter_dec` is an extension function.
- `mpc_counter_inc` is an extension function.
- `mpc_declassify` is an extension function.
- `mpc_input_secret_i64` is an extension function.
- `mpc_open` is an extension function.
- `mpc_reconfigure` is an extension function.
- `mpc_reset_party_state` is an extension function.
- `mpc_reshare` is an extension function.
- `mpc_set` is an extension function.
- `mpc_set_add` is an extension function.
- `mpc_set_contains` is an extension function.
- `mpc_set_coordinator_quorum` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
