## Usage

Sources:

- [Official README](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/README.md)
- [Control file](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/pg_gftr.control)
- [Extension SQL](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/sql/pg_gftr--0.1.0.sql)
- [Upstream validation record](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/VALIDATION.md)

`pg_gftr` 0.1.0 is an experimental PostgreSQL 18 extension for CUDA sort-based grouped aggregation. GPU execution needs an NVIDIA GPU, compatible driver and a CUDA-enabled build using CUDA Toolkit 12.2 or later. A separate CPU backend is a correctness reference. No PostgreSQL core patch is required.

### Core Workflow

Install and load as a superuser, then explicitly enable the planner path:

```sql
CREATE EXTENSION pg_gftr;
LOAD 'pg_gftr';
SELECT pg_gftr_version();
SET pg_gftr.enabled = on;
SET pg_gftr.backend = 'cuda';
SET pg_gftr.mode = 'gftr';

CREATE TEMP TABLE measurements (k integer, a integer, b integer);
INSERT INTO measurements
SELECT g % 100, g % 1000, g % 50 FROM generate_series(1, 200000) g;
ANALYZE measurements;
EXPLAIN (ANALYZE, BUFFERS)
SELECT k, count(*), count(a), sum(a), min(b), max(b)
FROM measurements GROUP BY k;
```

Eligible plans can use `Custom Scan (GftrAggregate)`. PostgreSQL performs the child scan, filtering and any CPU joins; the extension aggregates the resulting rows. It supports one plain integer grouping column and the shown built-in integer aggregates, including null keys and payloads.

### Controls

| Setting | Default | Purpose |
| --- | --- | --- |
| `pg_gftr.enabled` | off | Offer custom aggregation paths |
| `pg_gftr.force` | off | Give eligible paths an artificially low cost for experiments |
| `pg_gftr.backend` | cuda | Choose CUDA or the explicit CPU reference |
| `pg_gftr.mode` | gftr | Select GFTR or GFUR CUDA aggregation |
| `pg_gftr.device` | 0 | Device ordinal within the visible GPU set |
| `pg_gftr.min_rows` | 100000 | Estimated input threshold unless forced |
| `pg_gftr.max_input_mb` | 256 | Packed input limit in MiB |
| `pg_gftr.max_device_mb` | 4096 | Per-operator device allocation limit in MiB |
| `pg_gftr.setup_cost` | 1000 | Provisional setup cost |
| `pg_gftr.tuple_cost` | 0.02 | Provisional per-row cost |

The cost model is not calibrated for every GPU. Reprepare cached statements after changing planner controls. `CUDA_VISIBLE_DEVICES` restricts the available device set. CUDA failures are reported as query errors; there is no automatic CPU fallback.

### Session and Resource Boundaries

The installation script loads the hook in its creating session. Each new session must load the library, call its version function, or use administrator-configured preloading. Session loading needs no server restart. Simply installing SQL objects does not activate other sessions, and the planner path remains disabled until enabled.

The entire input is buffered. Host results, sort workspaces and PostgreSQL child-plan memory are additional to the input limit. Device scratch space counts toward the device limit, also capped at 90% of currently free GPU memory; concurrent sessions have no coordinated reservation. There is no disk spill, partial/final aggregation or multi-GPU execution. Exceeding capacity fails the statement without partial results. An in-flight GPU operation is not immediately interruptible.

Multiple grouping keys, expression keys, non-integer payloads, global aggregation, distinct or filtered aggregates, grouping sets and window queries are outside this version; unsupported shapes keep native paths. Upstream records GPU correctness checks on RTX A6000 and CPU checks on macOS; A100 execution remains untested in that record. Treat this as a research implementation and measure the whole SQL query, including data transfer.
