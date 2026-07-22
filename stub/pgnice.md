## Usage

Sources:

- [Official README](https://github.com/tensorchord/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/README.md)
- [SQL-visible function definitions](https://github.com/tensorchord/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/src/lib.rs)
- [Resource-limit implementation](https://github.com/tensorchord/pgnice/blob/3abc7e1d9449097c17eae5346021f0c2ea576fce/src/rlimit.rs)

`pgnice` 0.0.0 lets a session inspect and change operating-system scheduling and resource limits for its current PostgreSQL backend process. Use it to deprioritize selected workloads or constrain a dedicated pool of backends. The settings affect a process, not a transaction or database role.

### Core Workflow

Create the extension, inspect the current backend, then apply only limits supported by the host OS and PostgreSQL service account:

```sql
CREATE EXTENSION pgnice;

SELECT pgnice.get_backend_nice();
SELECT pgnice.get_backend_ionice();
SELECT pgnice.get_backend_rlimit('nofile');

SELECT pgnice.set_backend_nice(15);
SELECT pgnice.set_backend_ionice('B', 7);
SELECT pgnice.set_backend_rlimit('nofile', 4096);
```

`set_backend_rlimit` assigns the same value to both the soft and hard limit. Supported resource names vary by operating system and include common values such as `cpu`, `data`, `fsize`, `nofile`, `nproc`, `rss`, `stack`, and `as`.

### Function Index

- `pgnice.get_backend_nice()` and `pgnice.set_backend_nice(integer)` read or set CPU scheduling niceness.
- `pgnice.get_backend_ionice()` and `pgnice.set_backend_ionice(char, integer)` read or set I/O class and level; classes are `R` (real-time), `B` (best-effort), and `I` (idle).
- `pgnice.get_backend_rlimit(text)` returns the current soft and hard resource limit.
- `pgnice.set_backend_rlimit(text, bigint)` sets both resource limits for the backend.

### Operational and Security Caveats

The kernel enforces privilege checks. An ordinary PostgreSQL service account can usually lower its priority or hard limits but cannot later raise them, so a setting may be irreversible for the lifetime of that backend. With session or transaction pooling, the modified process can be reused by another client; separate pools and initialize each pool's backends explicitly. I/O priority support is Linux-specific, while resource names differ across platforms.

Changing limits such as address space, open files, CPU time, or process count can terminate or destabilize a backend. Revoke public execution and grant only the functions and values required by a controlled workload policy.
