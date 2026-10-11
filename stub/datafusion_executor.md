## Usage

Sources:

- [README.md](https://github.com/MisterRaindrop/cloudberry-datafusion/blob/cc65001b59e2dcc7ba1d76ad825b3c92729b8511/README.md)
- [datafusion_executor--1.0.sql](https://github.com/MisterRaindrop/cloudberry-datafusion/blob/cc65001b59e2dcc7ba1d76ad825b3c92729b8511/datafusion_executor--1.0.sql)
- [datafusion_executor.control](https://github.com/MisterRaindrop/cloudberry-datafusion/blob/cc65001b59e2dcc7ba1d76ad825b3c92729b8511/datafusion_executor.control)

`datafusion_executor` routes eligible Cloudberry plan fragments through a Rust DataFusion engine. It is an extension for the Cloudberry kernel, with a native C interface; stock PostgreSQL portability is not established.

### Core Workflow

```ini
shared_preload_libraries = 'datafusion_executor'
```

```sql
CREATE EXTENSION datafusion_executor;
SELECT datafusion_version();
SET datafusion.mode = 'explain';
```

### Operational Boundaries

Load `datafusion_executor` in `shared_preload_libraries` and restart coordinator and segment processes so executor hooks exist in every backend. Upstream also documents database-level session preload for new sessions. Install the SQL objects as an administrator. `datafusion.mode` defaults to `off`; `explain` diagnoses eligible fragments without using the alternative executor, while `on` enables it. Unsupported fragments keep the ordinary executor. `datafusion.worker_threads` controls the per-backend Rust worker count; the automatic CPU-count setting can oversubscribe a multi-backend installation. Memory is accounted against Cloudberry limits, but spill files are not covered by its normal workfile limits. Floating-point aggregates can differ in their final digits. Diagnostic panic and runtime test functions are intentionally revoked from PUBLIC.
