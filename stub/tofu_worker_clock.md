## Usage

Sources:

- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/docs/STORAGE_OPERATIONS.md)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/tofu_worker_clock.control)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/tofu_worker_clock--1.0.sql)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/worker_clock.c)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/generated_clock_contract.h)

`tofu_worker_clock` 1.0 supplies ToFu SQL worker operations with a server-owned lease clock and epoch. It requires PostgreSQL 17 or newer and is a separate native module inside the ToFu application repository. The application release number is not the extension's control version.

### Core Workflow

The deployment operator installs matching extension files and creates the extension as a superuser in each participating database. Application startup does not install it. The fixed schema is `public`.

```sql
CREATE EXTENSION tofu_worker_clock;
SELECT public.tofu_worker_lease_clock_v1();
```

The function returns JSONB containing `format`, `epoch`, `nowMs`, `utcMs` and `stateBytes`. `nowMs` is the shared high-water lease time; `utcMs` is sampled wall-clock time. Keep absolute UTC scheduling separate from elapsed lease expiry. `epoch` identifies the server lifetime; ToFu additionally scopes that identity by database OID.

### Lifetime and Failure Boundaries

The first call lazily creates bounded shared dynamic-memory state. Other sessions share it; closing clients or replacing a connection pool does not reset it. No preload, clock table, background worker or restart is needed to enable the function.

Restarting PostgreSQL creates a new epoch and revokes prior ToFu lease authority. This is not a persistent sequence, cross-server clock or guarantee that wall time cannot move backwards. The upstream SQL worker layer rejects missing or incompatible output rather than falling back to client time. Restrict installation privileges and preserve the exact result contract when integrating it with worker claims.
